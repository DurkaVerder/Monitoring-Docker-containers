package pinger

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"pinger-service/config"
	"sync"
	"time"

	"github.com/DurkaVerder/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/go-ping/ping"
)

type Pinger interface {
	Run()
}

type PingerService struct {
	dockerClient     *client.Client
	pingResultChan   chan models.PingResult
	dockerContainers chan types.Container
	wg               sync.WaitGroup
	config           *config.Config
}

func NewPingerService(cli *client.Client, pingResultChan chan models.PingResult, dockerContainers chan types.Container, cfg *config.Config) *PingerService {
	return &PingerService{
		dockerClient:     cli,
		pingResultChan:   pingResultChan,
		dockerContainers: dockerContainers,
		wg:               sync.WaitGroup{},
		config:           cfg,
	}
}

func (p *PingerService) getContainers() ([]types.Container, error) {
	containers, err := p.dockerClient.ContainerList(context.Background(), container.ListOptions{})
	log.Printf("Count containers: %d", len(containers))
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (p *PingerService) pingAllContainer() error {
	containers, err := p.getContainers()
	if err != nil {
		log.Printf("Error getting containers: %v", err)
		return err
	}

	if len(containers) == 0 {
		log.Println("No containers to ping")
		return nil
	}

	for _, container := range containers {
		p.dockerContainers <- container
	}

	return nil
}

func (p *PingerService) pingContainer(container types.Container) error {
	var pingResult models.PingResult

	details, err := p.dockerClient.ContainerInspect(context.Background(), container.ID)
	if err != nil {
		log.Printf("Error inspecting container: %v", err)
		return err
	}

	for networkName, network := range details.NetworkSettings.Networks {
		ip := network.IPAddress
		if network == nil || ip == "" {
			log.Printf("Container %s has no IP address in network %s", container.ID, networkName)
			continue
		}

		pinger, err := ping.NewPinger(ip)
		if err != nil {
			log.Printf("Error creating pinger: %v", err)
			return err
		}

		pinger.Count = 5
		pinger.Timeout = time.Second * 3
		pinger.Run()

		pingResult.IPAddress = ip
		pingResult.PingTime = int(pinger.Statistics().AvgRtt.Milliseconds())
		if pinger.Statistics().PacketLoss == 0 {
			pingResult.DateSuccessfulPing.Time = time.Now()
			pingResult.DateSuccessfulPing.Valid = true
		} else {
			pingResult.DateSuccessfulPing.Valid = false
		}

		p.pingResultChan <- pingResult
	}
	return errors.New(container.ID + " no network found")
}

func (p *PingerService) sendPingResults(pingResult models.PingResult) error {
	for i := 0; i < p.config.Response.RetryCount; i++ {
		err := p.trySendPingResult(pingResult)
		if err == nil {
			log.Printf("Ping result sent: %v", pingResult)
			return nil
		}

		time.Sleep(time.Second * time.Duration(p.config.Response.RetryInterval))
	}
	return errors.New("error sending ping result")
}

func (p *PingerService) trySendPingResult(pingResult models.PingResult) error {
	jsonPingResult, err := json.Marshal(pingResult)
	if err != nil {
		log.Printf("Error marshalling ping result: %v", err)
		return err
	}

	response, err := http.Post(p.config.Response.Address, "application/json", bytes.NewBuffer(jsonPingResult))
	if err != nil {
		log.Printf("Error sending ping result: %v", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error sending ping result: %v", err)
		return errors.New("Error sending ping result with status code: " + response.Status)
	}

	return nil
}

func (p *PingerService) workerPingContainer(ctx context.Context) {
	defer p.wg.Done()
	for {
		select {
		case container, ok := <-p.dockerContainers:
			if !ok {
				log.Printf("Worker stopping: dockerContainers channel closed")
				return
			}
			err := p.pingContainer(container)
			if err != nil {
				log.Printf("Error pinging container: %v", err)
			}
		case <-ctx.Done():
			log.Printf("Worker stopping: context done")
			return
		}
	}
}

func (p *PingerService) workerSendPingResult(ctx context.Context) {
	defer p.wg.Done()
	for {
		select {
		case pingResult, ok := <-p.pingResultChan:
			if !ok {
				log.Printf("Worker stopping: pingResultChan channel closed")
				return
			}
			err := p.sendPingResults(pingResult)
			if err != nil {
				log.Printf("Error sending ping result: %v", err)
			}
		case <-ctx.Done():
			log.Printf("Worker stopping: context done")
			return
		}
	}
}

func (p *PingerService) startWorkers(ctx context.Context) {
	for i := 0; i < p.config.Worker.Count; i++ {
		p.wg.Add(1)
		go p.workerPingContainer(ctx)
		p.wg.Add(1)
		go p.workerSendPingResult(ctx)
	}
}

func (p *PingerService) Stop() {
	close(p.dockerContainers)
	close(p.pingResultChan)
	p.wg.Wait()
	log.Println("PingerService stopped")
}

func (p *PingerService) Run(ctx context.Context) {
	p.startWorkers(ctx)

	for {
		err := p.pingAllContainer()
		if err != nil {
			log.Printf("Error pinging containers: %v", err)
		}

		time.Sleep(time.Second * 10)
	}
}

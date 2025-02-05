package server

import (
	"backend-service/config"
	"backend-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handlers
	config   *config.Config
	r        *gin.Engine
}

func NewServer(handlers handlers.Handlers, config *config.Config) *Server {
	return &Server{
		handlers: handlers,
		config:   config,
	}
}

func (s *Server) Run() {
	s.r = gin.Default()

	s.RegisterRoutes()

	s.r.Run(s.config.Server.Port)
}

func (s *Server) RegisterRoutes() {
	s.r.GET(s.config.Routes.GetAllPing, s.handlers.HandlerGetAllPing)
	s.r.POST(s.config.Routes.AddPing, s.handlers.HandlerAddPing)
}

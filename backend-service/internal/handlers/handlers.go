package handlers

import (
	"backend-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	HandlerGetAllPing(ctx *gin.Context)
	HandlerAddPing(ctx *gin.Context)
}

type HandlersManager struct {
	service service.Service
}

func NewHandlersManager(service service.Service) *HandlersManager {
	return &HandlersManager{
		service: service,
	}
}

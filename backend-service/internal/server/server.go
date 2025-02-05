package server

import (
	"backend-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handlers
	r        *gin.Engine
}

func NewServer(handlers handlers.Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) Run(port string) {
	s.r = gin.Default()

	s.RegisterRoutes()

	s.r.Run(port)
}

func (s *Server) RegisterRoutes() {
	s.r.GET("/pings", s.handlers.HandlerGetAllPing)
	s.r.POST("/ping/add", s.handlers.HandlerAddPing)
}

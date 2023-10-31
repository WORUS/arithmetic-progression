package handler

import (
	"github.com/WORUS/arithmetic-progression/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", h.SetTask)
	r.GET("/tasks", h.GetTasks)

	return r
}

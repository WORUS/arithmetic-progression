package handler

import (
	"github.com/WORUS/arithmetic-progression/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func Newhandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/task", h.SetTask)
	r.GET("/task", h.SetTask)

	return r
}

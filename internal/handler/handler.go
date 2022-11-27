package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"virtual-strike-backend-go/internal/service"
	"virtual-strike-backend-go/pkg/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/api/login", h.Login)

	protected := router.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.POST("/register", h.Register)
	protected.POST("/time", h.Time)
	protected.POST("/upload", h.Upload)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}

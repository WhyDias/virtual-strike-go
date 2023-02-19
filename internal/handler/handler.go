package handler

import (
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
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

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	router.Use(mw)
	router.POST("/api/login", h.Login)
	protected := router.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.POST("/register", h.Register)
	protected.POST("/points", h.Point)
	protected.POST("/time", h.Time)
	protected.POST("/upload", h.Upload)
	protected.POST("/workDayInfo", h.WorkDayInfo)
	protected.POST("/logging", h.Logging)
	protected.POST("/customerInfo", h.Customer)
	protected.POST("/tariffInfo", h.Tariff)

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}

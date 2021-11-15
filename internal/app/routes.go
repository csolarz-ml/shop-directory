package app

import (
	"github.com/gin-gonic/gin"
)

func (h *handlers) mapRoutes(router *gin.Engine) {
	router.GET("/ping", h.HealthCheckHandler.Ping)
}

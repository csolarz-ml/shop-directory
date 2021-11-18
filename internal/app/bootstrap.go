package app

import (
	"github.com/csolarz-ml/shop-directory/internal/handler"
	"github.com/gin-gonic/gin"
)

type handlers struct {
	*handler.HealthCheckHandler
}

func Start() *gin.Engine {
	router := gin.Default()

	h := &handlers{
		HealthCheckHandler: handler.NewHealthCheckHandler(),
	}

	h.mapRoutes(router)

	return router
}

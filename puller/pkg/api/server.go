package api

import (
	"github.com/gin-gonic/gin"

	handler "github.com/thnkrn/comet/puller/pkg/api/handler"
	middleware "github.com/thnkrn/comet/puller/pkg/api/middleware"
	config "github.com/thnkrn/comet/puller/pkg/config"
	log "github.com/thnkrn/comet/puller/pkg/driver/log"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type Middlewares struct {
	ErrorHandler *middleware.ErrorHandler
}

type Handlers struct {
	TaskHandler *handler.TaskHandler
}

func NewServerHTTP(middlewares *Middlewares, handlers Handlers, log log.Logger, cfg config.Config) *ServerHTTP {
	engine := gin.New()
	log.Info("Server is started")

	engine.Use(gin.Recovery())

	// Use logger from Gin
	if cfg.CometPuller.Log.Tracing {
		engine.Use(gin.Logger())
	}

	// Use error handler
	engine.Use(middlewares.ErrorHandler.Handler())

	engine.GET("healthcheck", func(c *gin.Context) {
		c.String(200, "OK")
	})

	engine.POST("/manual/:db", handlers.TaskHandler.ManualIngest)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8081")
}

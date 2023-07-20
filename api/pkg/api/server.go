package api

import (
	"fmt"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	handler "github.com/thnkrn/comet/api/pkg/api/handler"
	middleware "github.com/thnkrn/comet/api/pkg/api/middleware"
	config "github.com/thnkrn/comet/api/pkg/config"
	log "github.com/thnkrn/comet/api/pkg/driver/log"
)

type Middlewares struct {
	ErrorHandler   *middleware.ErrorHandler
	Authentication *middleware.Authentication
}

type Handlers struct {
	UserHandler  *handler.UserHandler
	AdminHandler *handler.AdminHandler
	DevHandler   *handler.DevHandler
}

type ServerHTTP struct {
	app *fiber.App
}

func NewServerHTTP(middlewares *Middlewares, handlers Handlers, log log.Logger, cfg config.Config) *ServerHTTP {
	app := fiber.New(
		fiber.Config{
			// NOTE: enable SO_REUSEPORT,
			// https://pkg.go.dev/github.com/valyala/fasthttp/reuseport, https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/, https://github.com/gofiber/fiber/issues/180
			Prefork: cfg.Comet.Prefork,
			// NOTE: Override default JSON encoding, ref: https://docs.gofiber.io/guide/faster-fiber#custom-json-encoder-decoder
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
			// NOTE: Override default error handler
			ErrorHandler: middlewares.ErrorHandler.FiberErrorHandler(),
		})

	log.Info(fmt.Sprintf("Server is started with PID: %v and PPID: %v", os.Getpid(), os.Getppid()))

	// NOTE: Enable log tracing from Fiber, https://docs.gofiber.io/api/middleware/logger
	if cfg.Comet.Log.Tracing {
		app.Use(logger.New())
	}

	if cfg.Comet.Recover {
		app.Use(recover.New())
	}

	// NOTE: Prometheus
	app.Get("metrics", adaptor.HTTPHandler(promhttp.Handler()))

	// NOTE: Healthcheck
	app.Get("healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// NOTE: Group user API
	userAPI := app.Group("/databases")
	userAPI.Get(":db<minLen(1)>/keys/:key<minLen(1)>", handlers.UserHandler.Get)
	userAPI.Put(":db<minLen(1)>/keys/:key<minLen(1)>", handlers.UserHandler.Create)
	userAPI.Delete(":db<minLen(1)>/keys/:key<minLen(1)>", handlers.UserHandler.Delete)
	userAPI.Get(":db<minLen(1)>/count", handlers.UserHandler.Count)
	userAPI.Get(":db<minLen(1)>/keys", handlers.UserHandler.MultiGet)

	// NOTE: Group admin API and init authentication
	adminAPI := app.Group("/admin/databases", middlewares.Authentication.Authentication(middleware.ADMIN_ROLE))
	adminAPI.Post(":db<minLen(1)>/catch-up-with-primary", handlers.AdminHandler.CatchUpWithPrimary)
	adminAPI.Get(":db<minLen(1)>/properties/:property<minLen(1)>", handlers.AdminHandler.GetDBProperty)
	adminAPI.Put(":db<minLen(1)>/checkpoint/:directory<minLen(1)>", handlers.AdminHandler.CreateCheckPoint)
	adminAPI.Put(":db<minLen(1)>/ingests/:directory<minLen(1)>", handlers.AdminHandler.Ingest)
	adminAPI.Get(":db<minLen(1)>/ingests/last", handlers.AdminHandler.GetLastIngest)

	// NOTE: Group dev API and init authentication
	devAPI := app.Group("/dev", middlewares.Authentication.Authentication(middleware.DEV_ROLE))
	devAPI.Put(":fileName<minLen(1)>/sst/:key<minLen(1)>", handlers.DevHandler.AddValueToSSTFile)
	devAPI.Put("sst/:fileName<minLen(1)>/ingest/:source<minLen(1)>/:ingestFolder<minLen(1)>", handlers.DevHandler.PullFile)
	devAPI.Get("db", handlers.DevHandler.ListDB)

	return &ServerHTTP{app}
}

func (sh *ServerHTTP) Start() {
	sh.app.Listen(":8080")
}

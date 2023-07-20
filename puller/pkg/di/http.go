package di

import (
	"github.com/google/wire"

	api "github.com/thnkrn/comet/puller/pkg/api"
	middleware "github.com/thnkrn/comet/puller/pkg/api/middleware"
)

var HttpSet = wire.NewSet(
	api.NewServerHTTP,
	middleware.NewErrorHandler,
	wire.Struct(new(api.Middlewares), "*"),
	wire.Struct(new(api.Handlers), "*"),
)

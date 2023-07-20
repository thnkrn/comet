package di

import (
	"github.com/google/wire"
	client "github.com/thnkrn/comet/puller/pkg/driver/client"
	adapterClient "github.com/thnkrn/comet/puller/pkg/driver/client/adapter"
)

var HTTPClientSet = wire.NewSet(
	wire.Bind(new(client.HTTP), new(*adapterClient.HTTPClient)),
	adapterClient.NewHTTPClient,
)

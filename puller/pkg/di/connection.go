package di

import (
	"github.com/google/wire"

	config "github.com/thnkrn/comet/puller/pkg/config"
)

var ConnectionSet = wire.NewSet(
	config.ConnectGCS,
)

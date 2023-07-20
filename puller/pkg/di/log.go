package di

import (
	"github.com/google/wire"

	log "github.com/thnkrn/comet/puller/pkg/driver/log"
	logAdapter "github.com/thnkrn/comet/puller/pkg/driver/log/adapter"
	logConfig "github.com/thnkrn/comet/puller/pkg/driver/log/config"
)

var LogSet = wire.NewSet(logConfig.ProvidZapLogger, wire.Bind(new(log.Logger), new(*logAdapter.ZapImplement)), logAdapter.ProvideLogger)

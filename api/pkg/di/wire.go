//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	api "github.com/thnkrn/comet/api/pkg/api"
	config "github.com/thnkrn/comet/api/pkg/config"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(ConnectionSet, UserSet, AdminSet, DevSet, LogSet, HTTPSet)

	return &api.ServerHTTP{}, nil
}

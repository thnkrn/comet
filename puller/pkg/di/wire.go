//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	api "github.com/thnkrn/comet/puller/pkg/api"
	config "github.com/thnkrn/comet/puller/pkg/config"
	scheduler "github.com/thnkrn/comet/puller/pkg/scheduler"
)

type Application struct {
	Server    *api.ServerHTTP
	Scheduler *scheduler.Scheduler
}

func InitializeApp(config config.Config) (*Application, error) {
	wire.Build(LogSet, TaskSet, HttpSet, HTTPClientSet, RocksDBAdminSet, FileSet, ConnectionSet, JobSet, JobStageSet, wire.Struct(new(Application), "*"))

	return &Application{}, nil
}

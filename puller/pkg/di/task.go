package di

import (
	"github.com/google/wire"
	handler "github.com/thnkrn/comet/puller/pkg/api/handler"
	repositoryAdapter "github.com/thnkrn/comet/puller/pkg/repository/adapter"
	usecaseAdapter "github.com/thnkrn/comet/puller/pkg/usecase/adapter"
)

var TaskSet = wire.NewSet(
	usecaseAdapter.NewTaskUsecase,
	handler.NewTaskHandler,
	repositoryAdapter.NewTaskRepository,
)

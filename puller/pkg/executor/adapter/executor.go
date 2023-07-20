package adapter

import (
	"context"

	"github.com/go-co-op/gocron"

	config "github.com/thnkrn/comet/puller/pkg/config"
	domain "github.com/thnkrn/comet/puller/pkg/domain"
	usecase "github.com/thnkrn/comet/puller/pkg/usecase"
)

type ExecutorImplement struct {
	taskUsecase usecase.TaskUsecase
	cfg         config.Config
}

func NewExecutor(u usecase.TaskUsecase, cfg config.Config) *ExecutorImplement {
	return &ExecutorImplement{taskUsecase: u, cfg: cfg}
}

func ProvideExecutor(u usecase.TaskUsecase, cfg config.Config) *ExecutorImplement {
	return NewExecutor(u, cfg)
}

func (e *ExecutorImplement) Execute(j domain.JobData, job gocron.Job) {
	ctx := context.Background()
	e.taskUsecase.Perform(ctx, j.DBNameKey, j.IgnoreLastIngest)
}

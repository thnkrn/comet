package di

import (
	"github.com/google/wire"

	executor "github.com/thnkrn/comet/puller/pkg/executor"
	executorAdapter "github.com/thnkrn/comet/puller/pkg/executor/adapter"
	scheduler "github.com/thnkrn/comet/puller/pkg/scheduler"
)

var JobSet = wire.NewSet(
	wire.Bind(new(executor.Executor), new(*executorAdapter.ExecutorImplement)),
	executorAdapter.ProvideExecutor,
	scheduler.NewScheduler,
)

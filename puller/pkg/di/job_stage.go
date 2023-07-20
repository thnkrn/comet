package di

import (
	"github.com/google/wire"
	jobStage "github.com/thnkrn/comet/puller/pkg/driver/job_stage"
)

var JobStageSet = wire.NewSet(
	jobStage.CreateJobPool,
)

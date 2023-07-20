package executor

import (
	"github.com/go-co-op/gocron"

	domain "github.com/thnkrn/comet/puller/pkg/domain"
)

type Executor interface {
	Execute(j domain.JobData, job gocron.Job)
}

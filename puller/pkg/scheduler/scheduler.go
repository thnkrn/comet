package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"

	config "github.com/thnkrn/comet/puller/pkg/config"
	domain "github.com/thnkrn/comet/puller/pkg/domain"
	log "github.com/thnkrn/comet/puller/pkg/driver/log"
	executor "github.com/thnkrn/comet/puller/pkg/executor"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
}

func NewScheduler(log log.Logger, cfg config.Config, executor executor.Executor) (*Scheduler, error) {
	scheduler := gocron.NewScheduler(time.UTC)
	log.Info("Scheduler is started")

	for _, bucket := range cfg.App.SSTBuckets {
		jobData := domain.NewJobData(bucket.Name, bucket.Source, cfg.App.APIToken, bucket.IngestOnlyLatestDir, false)

		if bucket.Schedule != "" {
			_, err := scheduler.Cron(bucket.Schedule).DoWithJobDetails(executor.Execute, jobData)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Scheduler{scheduler: scheduler}, nil
}

func (sj *Scheduler) Start() {
	sj.scheduler.StartAsync()
}

package job_stage

import (
	config "github.com/thnkrn/comet/puller/pkg/config"
	domain "github.com/thnkrn/comet/puller/pkg/domain"
)

func CreateJobPool(cfg config.Config) (domain.JobPools, error) {
	jobPool := make(map[string]domain.Job)
	for _, s := range cfg.App.SSTBuckets {
		newJob := domain.NewJob(s.Name, s.Source, cfg.App.APIToken, s.IngestOnlyLatestDir)
		jobPool[s.Name] = newJob
	}
	return domain.JobPools(jobPool), nil
}

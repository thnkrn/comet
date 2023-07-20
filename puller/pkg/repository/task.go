package repository

import (
	"context"

	domain "github.com/thnkrn/comet/puller/pkg/domain"
)

type TaskRepository interface {
	Ingest(ctx context.Context, authorization, db, directory string) error
	GetLastIngest(ctx context.Context, authorization, db string) (string, error)
	GetLists(ctx context.Context, prefix string) ([]string, error)
	DownloadSST(ctx context.Context, fromGCS string) (bool, error)
	PerformJobPool(jobPool domain.JobPools, db string) (domain.Job, error)
}

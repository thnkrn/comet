package usecase

import (
	"context"
)

type TaskUsecase interface {
	GetLatestSeries(objects []string, ingestOnlyLatestDirectory bool) ([]string, error)
	GetTargetStorageObjects(ctx context.Context, dbName, source, authorization string, ingestOnlyLatestDirectory, ignoreLastIngest bool) ([]string, error)
	Ingest(ctx context.Context, authorization, dbName, directory string) (bool, error)
	Download(ctx context.Context, gcsObject string) (bool, error)
	RemoveTemp(directory string) error
	Perform(ctx context.Context, db string, ignoreLastIngest bool) error
}

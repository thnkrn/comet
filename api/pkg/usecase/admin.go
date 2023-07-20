package usecase

import (
	"context"
)

type AdminUsecase interface {
	CatchUpWithPrimary(ctx context.Context, dbName string) error
	GetDBProperty(ctx context.Context, dbName, property string) (string, error)
	CreateCheckPoint(ctx context.Context, dbName, dir string) error
	Ingest(ctx context.Context, dbName, dir string) error
	GetLastIngest(ctx context.Context, dbName string) ([]byte, error)
}

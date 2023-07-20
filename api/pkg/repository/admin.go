package repository

import (
	"context"
)

type AdminRepository interface {
	CatchUpWithPrimary(ctx context.Context, dbName string) error
	GetDBProperty(ctx context.Context, dbName, property string) (string, error)
	CreateCheckPoint(ctx context.Context, dbName, dir string) error
	Ingest(ctx context.Context, dbName, dir string, filesPath []string) error
	GetLastIngest(ctx context.Context, dbName string) ([]byte, error)
}

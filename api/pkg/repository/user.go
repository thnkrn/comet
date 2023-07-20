package repository

import (
	"context"
)

type UserRepository interface {
	GetByKey(ctx context.Context, dbName, key string) ([]byte, error)
	CreateByKey(ctx context.Context, dbName, key string, value []byte) ([]byte, error)
	DeleteByKey(ctx context.Context, dbName, key string) error
	Count(ctx context.Context, dbName string) (string, error)
	GetLastIngest(ctx context.Context, dbName string) ([]byte, error)
	MultiGet(ctx context.Context, dbName string, keys [][]byte) ([]string, error)
}

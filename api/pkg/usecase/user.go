package usecase

import (
	"context"
)

type UserUsecase interface {
	Get(ctx context.Context, dbName, key string) ([]byte, error)
	Create(ctx context.Context, dbName, key string, value []byte) ([]byte, error)
	Delete(ctx context.Context, dbName, key string) error
	Count(ctx context.Context, dbName string) (string, error)
	MultiGet(ctx context.Context, dbName string, keys []string) ([]string, error)
}

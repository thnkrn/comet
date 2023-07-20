package repository

import (
	"context"

	domain "github.com/thnkrn/comet/api/pkg/domain"
)

type DevRepository interface {
	AddValueToSSTFile(ctx context.Context, fileName, key string, value []byte) ([]byte, error)
	PullFile(ctx context.Context, fileName, source, ingestFolder string) error
	ListDB(ctx context.Context) []domain.DB
}

package adapter

import (
	"context"

	config "github.com/thnkrn/comet/api/pkg/config"
	domain "github.com/thnkrn/comet/api/pkg/domain"
	repository "github.com/thnkrn/comet/api/pkg/repository"
	usecase "github.com/thnkrn/comet/api/pkg/usecase"
)

type devUsecase struct {
	devRepo repository.DevRepository
	cfg     config.Config
}

func NewDevUseCase(r repository.DevRepository, cfg config.Config) usecase.DevUsecase {
	return &devUsecase{devRepo: r, cfg: cfg}
}

func (u *devUsecase) AddValueToSSTFile(ctx context.Context, fileName, key string, value []byte) ([]byte, error) {
	result, err := u.devRepo.AddValueToSSTFile(ctx, fileName, key, value)
	return result, err
}

func (u *devUsecase) PullFile(ctx context.Context, fileName, source, ingestFolder string) error {
	err := u.devRepo.PullFile(ctx, fileName, source, ingestFolder)
	return err
}

func (u *devUsecase) ListDB(ctx context.Context) []domain.DB {
	result := u.devRepo.ListDB(ctx)
	return result
}

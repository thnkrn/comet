package adapter

import (
	"context"
	"errors"
	"strconv"

	repository "github.com/thnkrn/comet/api/pkg/repository"
	usecase "github.com/thnkrn/comet/api/pkg/usecase"
	usecaseError "github.com/thnkrn/comet/api/pkg/usecase/error"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{userRepo: r}
}

func (u *userUsecase) Get(ctx context.Context, dbName, key string) ([]byte, error) {
	result, err := u.userRepo.GetByKey(ctx, dbName, key)
	if err == nil && (result == nil || len(result) <= 0) {
		return result, usecaseError.NewErrorNotFound(errors.New("value not found"))
	}
	return result, err
}

func (u *userUsecase) Create(ctx context.Context, dbName, key string, value []byte) ([]byte, error) {
	result, err := u.userRepo.CreateByKey(ctx, dbName, key, value)
	return result, err
}

func (u *userUsecase) Delete(ctx context.Context, dbName, key string) error {
	err := u.userRepo.DeleteByKey(ctx, dbName, key)
	return err
}

func (u *userUsecase) Count(ctx context.Context, dbName string) (string, error) {
	var result string

	estimatedKeys, err := u.userRepo.Count(ctx, dbName)
	if err != nil {
		return result, err
	}

	iEstimatedKeys, err := strconv.Atoi(estimatedKeys)
	if err != nil {
		return result, err
	}

	lastIngestResult, err := u.userRepo.GetLastIngest(ctx, dbName)
	if err != nil {
		return result, err
	}

	if lastIngestResult == nil && err == nil {
		return result, nil
	}

	result = strconv.Itoa(iEstimatedKeys - 1)

	return result, nil
}

func (u *userUsecase) MultiGet(ctx context.Context, dbName string, keys []string) ([]string, error) {
	byteKeys := make([][]byte, len(keys))
	for i, v := range keys {
		byteKeys[i] = []byte(v)
	}

	result, err := u.userRepo.MultiGet(ctx, dbName, byteKeys)
	if err != nil {
		return nil, err
	}

	return result, err
}

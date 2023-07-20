package adapter_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	config "github.com/thnkrn/comet/api/pkg/config"
	"github.com/thnkrn/comet/api/pkg/domain"
	mrepository "github.com/thnkrn/comet/api/pkg/mocks/repository"
	usecase "github.com/thnkrn/comet/api/pkg/usecase"
	usecaseAdapter "github.com/thnkrn/comet/api/pkg/usecase/adapter"
)

type devDependencies struct {
	mockDevRepo *mrepository.DevRepository
}

func createDevUsecase() (usecase.DevUsecase, *devDependencies) {
	mockDevRepo := new(mrepository.DevRepository)
	mockConfig := config.Config{}
	usecase := usecaseAdapter.NewDevUseCase(mockDevRepo, mockConfig)

	return usecase, &devDependencies{mockDevRepo}
}

func TestDevAddValueToSSTFile(t *testing.T) {
	t.Run("It should return result without an error", func(t *testing.T) {
		mockResponse := []byte("value")
		usecase, deps := createDevUsecase()
		deps.mockDevRepo.On("AddValueToSSTFile", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.AddValueToSSTFile(context.TODO(), "fileName", "key", []byte("value"))

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockDevRepo.AssertExpectations(t)
	})
}

func TestDevPullFile(t *testing.T) {
	t.Run("It should return err as nil if pull file successfully", func(t *testing.T) {
		usecase, deps := createDevUsecase()
		deps.mockDevRepo.On("PullFile", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		err := usecase.PullFile(context.TODO(), "fileName", "source", "ingestFolder")

		assert.NoError(t, err)

		deps.mockDevRepo.AssertExpectations(t)
	})
}

func TestListDB(t *testing.T) {
	t.Run("It should return result without an error", func(t *testing.T) {
		mockResponse := []domain.DB{}
		usecase, deps := createDevUsecase()
		deps.mockDevRepo.On("ListDB", context.TODO()).Return(mockResponse).Once()
		res := usecase.ListDB(context.TODO())

		assert.Equal(t, res, mockResponse)

		deps.mockDevRepo.AssertExpectations(t)
	})
}

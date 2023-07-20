package adapter_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mrepository "github.com/thnkrn/comet/api/pkg/mocks/repository"
	usecase "github.com/thnkrn/comet/api/pkg/usecase"
	usecaseAdapter "github.com/thnkrn/comet/api/pkg/usecase/adapter"
)

type userDependencies struct {
	mockUserRepo *mrepository.UserRepository
}

func createUserUsecase() (usecase.UserUsecase, *userDependencies) {
	mockUserRepo := new(mrepository.UserRepository)
	usecase := usecaseAdapter.NewUserUseCase(mockUserRepo)

	return usecase, &userDependencies{mockUserRepo}
}

func TestUserGet(t *testing.T) {
	dbName := "db"
	key := "xyz"

	t.Run("It should return result", func(t *testing.T) {
		mockResponse := []byte("result")

		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("GetByKey", context.TODO(), mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.Get(context.TODO(), dbName, key)

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return error if result is nil", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("GetByKey", context.TODO(), mock.Anything, mock.Anything).Return(nil, nil).Once()
		_, err := usecase.Get(context.TODO(), dbName, key)

		assert.Equal(t, "value not found", err.Error())

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserCreate(t *testing.T) {
	dbName := "db"
	key := "xyz"

	t.Run("It should create without an error", func(t *testing.T) {
		mockResponse := []byte("result")

		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("CreateByKey", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.Create(context.TODO(), dbName, key, mockResponse)

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserDelete(t *testing.T) {
	dbName := "db"
	key := "xyz"

	t.Run("It should delete without an error", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("DeleteByKey", context.TODO(), mock.Anything, mock.Anything).Return(nil).Once()
		err := usecase.Delete(context.TODO(), dbName, key)

		assert.NoError(t, err)

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserCount(t *testing.T) {
	dbName := "db1"

	t.Run("It should count correctly if the database is ingested", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("Count", context.TODO(), mock.Anything).Return("1000", nil).Once()
		deps.mockUserRepo.On("GetLastIngest", context.TODO(), mock.Anything).Return([]byte("ingestFolder"), nil).Once()

		res, err := usecase.Count(context.TODO(), dbName)

		assert.NoError(t, err)
		assert.Equal(t, res, "999")

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return empty string if the database is not ingested", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("Count", context.TODO(), mock.Anything).Return("1000", nil).Once()
		deps.mockUserRepo.On("GetLastIngest", context.TODO(), mock.Anything).Return(nil, nil).Once()

		res, err := usecase.Count(context.TODO(), dbName)

		assert.NoError(t, err)
		assert.Equal(t, res, "")

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return error if userRepo on Count occurs an error", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("Count", context.TODO(), mock.Anything).Return("", errors.New("Unexpected Error")).Once()

		res, err := usecase.Count(context.TODO(), dbName)

		assert.Equal(t, "Unexpected Error", err.Error())
		assert.Equal(t, res, "")

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return error if userRepo on GetLastIngest occurs an error", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("Count", context.TODO(), mock.Anything).Return("1000", nil).Once()
		deps.mockUserRepo.On("GetLastIngest", context.TODO(), mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()

		res, err := usecase.Count(context.TODO(), dbName)

		assert.Equal(t, "Unexpected Error", err.Error())
		assert.Equal(t, res, "")

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return error if strconv occurs an error", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("Count", context.TODO(), mock.Anything).Return("", nil).Once()

		res, err := usecase.Count(context.TODO(), dbName)

		assert.Equal(t, `strconv.Atoi: parsing "": invalid syntax`, err.Error())
		assert.Equal(t, res, "")

		deps.mockUserRepo.AssertExpectations(t)
	})
}

func TestUserMultiGet(t *testing.T) {
	dbName := "db"
	key := []string{"key1", "key2"}

	t.Run("It should multi get without an error", func(t *testing.T) {
		mockResponse := []string{"res1", "res2"}

		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("MultiGet", context.TODO(), mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.MultiGet(context.TODO(), dbName, key)

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockUserRepo.AssertExpectations(t)
	})

	t.Run("It should return if an error occurs", func(t *testing.T) {
		usecase, deps := createUserUsecase()
		deps.mockUserRepo.On("MultiGet", context.TODO(), mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()
		_, err := usecase.MultiGet(context.TODO(), dbName, key)

		assert.Equal(t, "Unexpected Error", err.Error())

		deps.mockUserRepo.AssertExpectations(t)
	})
}

package adapter_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	config "github.com/thnkrn/comet/api/pkg/config"
	mlog "github.com/thnkrn/comet/api/pkg/mocks/driver/log"
	mrepository "github.com/thnkrn/comet/api/pkg/mocks/repository"
	usecase "github.com/thnkrn/comet/api/pkg/usecase"
	usecaseAdapter "github.com/thnkrn/comet/api/pkg/usecase/adapter"
)

type adminDependencies struct {
	mockAdminRepo *mrepository.AdminRepository
	mockLog       *mlog.Logger
}

func createAdminUsecase() (usecase.AdminUsecase, *adminDependencies) {
	curDir, _ := os.Getwd()
	myDir := strings.ReplaceAll(curDir, "usecase/adapter", "")

	mockAdminRepo := new(mrepository.AdminRepository)
	mockConfig := config.Config{
		App: config.AppConfig{
			IngestPath: myDir + "mocks",
			Databases: []config.Databases{
				{
					Name:   "dbName",
					Source: "source",
				},
			},
		},
	}
	mockLog := new(mlog.Logger)
	usecase := usecaseAdapter.NewAdminUseCase(mockAdminRepo, mockConfig, mockLog)

	return usecase, &adminDependencies{mockAdminRepo, mockLog}
}

func TestAdminCatchUpWithPrimary(t *testing.T) {
	t.Run("It should return err as nil if catch up with primary successfully", func(t *testing.T) {
		usecase, deps := createAdminUsecase()
		deps.mockAdminRepo.On("CatchUpWithPrimary", context.TODO(), mock.Anything).Return(nil).Once()
		err := usecase.CatchUpWithPrimary(context.TODO(), "dbName")

		assert.NoError(t, err)

		deps.mockAdminRepo.AssertExpectations(t)
	})
}

func TestAdminGetDBProperty(t *testing.T) {
	t.Run("It should return result without an error", func(t *testing.T) {
		mockResponse := "property"
		usecase, deps := createAdminUsecase()
		deps.mockAdminRepo.On("GetDBProperty", context.TODO(), mock.Anything, mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.GetDBProperty(context.TODO(), "dbName", "property")

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockAdminRepo.AssertExpectations(t)
	})

	t.Run("It should return error if result not found", func(t *testing.T) {
		usecase, deps := createAdminUsecase()
		deps.mockAdminRepo.On("GetDBProperty", context.TODO(), mock.Anything, mock.Anything).Return("", nil).Once()
		res, err := usecase.GetDBProperty(context.TODO(), "dbName", "property")

		assert.Equal(t, "property not found", err.Error())
		assert.Equal(t, res, "")

		deps.mockAdminRepo.AssertExpectations(t)
	})
}

func TestAdminCreateCheckPoint(t *testing.T) {
	t.Run("It should return err as nil if create checkpoint successfully", func(t *testing.T) {
		usecase, deps := createAdminUsecase()
		deps.mockAdminRepo.On("CreateCheckPoint", context.TODO(), mock.Anything, mock.Anything).Return(nil).Once()
		err := usecase.CreateCheckPoint(context.TODO(), "dbName", "dir")

		assert.NoError(t, err)

		deps.mockAdminRepo.AssertExpectations(t)
	})
}

func TestAdminIngest(t *testing.T) {
	t.Run("It should return err as nil if ingest successfully", func(t *testing.T) {
		usecase, deps := createAdminUsecase()
		deps.mockLog.On("Info", mock.Anything).Once()
		deps.mockLog.On("Error", mock.Anything).Once()
		deps.mockAdminRepo.On("Ingest", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		err := usecase.Ingest(context.TODO(), "dbName", "ingest")

		assert.NoError(t, err)

		deps.mockAdminRepo.AssertExpectations(t)
	})

	t.Run("It should return err if SST file not found", func(t *testing.T) {
		usecase, deps := createAdminUsecase()
		deps.mockLog.On("Info", mock.Anything).Once()
		deps.mockLog.On("Error", mock.Anything).Once()
		err := usecase.Ingest(context.TODO(), "dbName", "invalidIngest")

		assert.Equal(t, "sst files not found", err.Error())

		deps.mockAdminRepo.AssertExpectations(t)
	})
}

func TestAdminGetLastIngest(t *testing.T) {
	t.Run("It should return err as nil if create checkpoint successfully", func(t *testing.T) {
		mockResponse := []byte("result")
		usecase, deps := createAdminUsecase()
		deps.mockAdminRepo.On("GetLastIngest", context.TODO(), mock.Anything).Return(mockResponse, nil).Once()
		res, err := usecase.GetLastIngest(context.TODO(), "dbName")

		assert.NoError(t, err)
		assert.Equal(t, res, mockResponse)

		deps.mockAdminRepo.AssertExpectations(t)
	})

	t.Run("It should return error if result is nil", func(t *testing.T) {
		usecase, deps := createAdminUsecase()
		deps.mockAdminRepo.On("GetLastIngest", context.TODO(), mock.Anything).Return(nil, nil).Once()
		_, err := usecase.GetLastIngest(context.TODO(), "dbName")

		assert.Equal(t, "value not found", err.Error())

		deps.mockAdminRepo.AssertExpectations(t)
	})
}

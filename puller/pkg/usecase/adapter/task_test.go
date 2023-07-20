package adapter_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	domain "github.com/thnkrn/comet/puller/pkg/domain"
	mfile "github.com/thnkrn/comet/puller/pkg/mocks/driver/file"
	mlog "github.com/thnkrn/comet/puller/pkg/mocks/driver/log"
	mrepository "github.com/thnkrn/comet/puller/pkg/mocks/repository"
	usecase "github.com/thnkrn/comet/puller/pkg/usecase"
	usecaseAdapter "github.com/thnkrn/comet/puller/pkg/usecase/adapter"
)

type taskDependencies struct {
	mockTaskRepo *mrepository.TaskRepository
	mockLog      *mlog.Logger
	mockFile     *mfile.File
}

func createTaskUsecase() (usecase.TaskUsecase, *taskDependencies) {

	mockTaskRepo := new(mrepository.TaskRepository)
	mockFile := new(mfile.File)
	mockLog := new(mlog.Logger)
	usecase := usecaseAdapter.NewTaskUsecase(mockTaskRepo, mockFile, domain.JobPools{}, mockLog)

	return usecase, &taskDependencies{mockTaskRepo, mockLog, mockFile}
}

func TestTaskGetLatestSeries(t *testing.T) {
	t.Run("It should return err as nil if get latest series successfully", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		objects := []string{"0_0", "0_1", "1_0", "1_1"}
		expectedResult := []string{"1_0", "1_1"}
		res, err := usecase.GetLatestSeries(objects, false)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err as nil if get latest series successfully", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		objects := []string{"0_0", "0_1", "1_0", "1_1"}
		expectedResult := []string{"1_1"}
		res, err := usecase.GetLatestSeries(objects, true)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err if get latest series failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		objects := []string{"ss"}
		expectedResult := []string(nil)
		deps.mockLog.On("Error", mock.Anything)
		res, err := usecase.GetLatestSeries(objects, false)

		assert.Equal(t, "strconv.Atoi: parsing \"ss\": invalid syntax", err.Error())
		assert.Equal(t, expectedResult, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskGetTargetStorageObjects(t *testing.T) {
	t.Run("It should return err as nil when latest series base value is more than current value", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{"0_3", "1_1"}
		mockLastIngest := "0_2"
		expectedResult := []string{"1_1"}
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, false)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err as nil when latest series base value is equal to current value", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{"0_3", "0_4"}
		mockLastIngest := "0_2"
		expectedResult := []string{"0_3", "0_4"}
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, false)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err as nil when we ignore lastingest", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{"0_3", "0_4"}
		mockLastIngest := "0_2"
		expectedResult := []string{"0_3", "0_4"}
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, true)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when we ignore last ingest but cannot get latest series", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{"ss"}
		mockLastIngest := "0_2"
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockLog.On("Error", mock.Anything)

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, true)

		assert.Equal(t, "strconv.Atoi: parsing \"ss\": invalid syntax", err.Error())
		assert.Equal(t, []string(nil), res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when cannot find object", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{}
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockLog.On("Error", mock.Anything)

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, false)

		assert.Equal(t, "cannot find object for db db1 from source", err.Error())
		assert.Equal(t, []string(nil), res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when cannot get last ingest", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{"0_3", "1_1"}
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return("", errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, false)

		assert.Equal(t, "Unexpected Error", err.Error())
		assert.Equal(t, []string(nil), res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when cannot get base on get lists", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{"ss", "aa"}
		mockLastIngest := "ss"
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockLog.On("Error", mock.Anything)

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, false)

		assert.Equal(t, "strconv.Atoi: parsing \"ss\": invalid syntax", err.Error())
		assert.Equal(t, []string(nil), res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when cannot get base on get last ingest", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockList := []string{"ss", "aa"}
		mockLastIngest := "1_0"
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockLog.On("Error", mock.Anything)

		res, err := usecase.GetTargetStorageObjects(context.TODO(), "db1", "source", "auth", false, false)

		assert.Equal(t, "strconv.Atoi: parsing \"ss\": invalid syntax", err.Error())
		assert.Equal(t, []string(nil), res)

		deps.mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskIngest(t *testing.T) {
	t.Run("It should return err as nil if ingest successfully", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		deps.mockTaskRepo.On("Ingest", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		res, err := usecase.Ingest(context.TODO(), "auth", "db", "dir")

		assert.NoError(t, err)
		assert.Equal(t, true, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err if ingest failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		deps.mockTaskRepo.On("Ingest", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)
		res, err := usecase.Ingest(context.TODO(), "auth", "db", "dir")

		assert.Equal(t, "Unexpected Error", err.Error())
		assert.Equal(t, false, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskDownload(t *testing.T) {
	t.Run("It should return err as nil if download successfully", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		deps.mockTaskRepo.On("DownloadSST", context.TODO(), mock.Anything).Return(true, nil).Once()

		res, err := usecase.Download(context.TODO(), "gcsObject")

		assert.NoError(t, err)
		assert.Equal(t, true, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err if download failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		deps.mockTaskRepo.On("DownloadSST", context.TODO(), mock.Anything).Return(false, errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)
		res, err := usecase.Download(context.TODO(), "gcsObject")

		assert.Equal(t, "Unexpected Error", err.Error())
		assert.Equal(t, false, res)

		deps.mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskRemoveTemp(t *testing.T) {
	t.Run("It should return err as nil if remove temp successfully", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		deps.mockFile.On("Remove", mock.Anything).Return(nil).Once()

		err := usecase.RemoveTemp("directory")

		assert.NoError(t, err)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err if remove temp failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		deps.mockFile.On("Remove", mock.Anything).Return(errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)
		err := usecase.RemoveTemp("directory")

		assert.Equal(t, "Unexpected Error", err.Error())

		deps.mockTaskRepo.AssertExpectations(t)
	})
}

func TestTaskPerform(t *testing.T) {
	t.Run("It should return err as nil when perform method run successfully", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockJob := domain.NewJob("db1", "sst_example", "JWT_TOKEN", false)
		mockList := []string{"0_3", "1_1"}
		mockLastIngest := "0_2"

		deps.mockTaskRepo.On("PerformJobPool", mock.Anything, mock.Anything).Return(mockJob, nil).Once()
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockTaskRepo.On("DownloadSST", context.TODO(), mock.Anything).Return(true, nil).Once()
		deps.mockTaskRepo.On("Ingest", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		deps.mockFile.On("Remove", mock.Anything).Return(nil).Once()

		err := usecase.Perform(context.TODO(), "db1", false)

		assert.NoError(t, err)

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when perform job pool failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockJob := domain.NewJob("db1", "sst_example", "JWT_TOKEN", false)
		deps.mockTaskRepo.On("PerformJobPool", mock.Anything, mock.Anything).Return(mockJob, errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)

		err := usecase.Perform(context.TODO(), "db1", false)

		assert.Equal(t, "couldn't perform job pool", err.Error())

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when get target storage object failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockJob := domain.NewJob("db1", "sst_example", "JWT_TOKEN", false)
		deps.mockTaskRepo.On("PerformJobPool", mock.Anything, mock.Anything).Return(mockJob, nil).Once()
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return([]string{}, nil).Once()
		deps.mockLog.On("Error", mock.Anything)

		err := usecase.Perform(context.TODO(), "db1", false)

		assert.Equal(t, "couldn't get target storage objects", err.Error())

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when download failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockJob := domain.NewJob("db1", "sst_example", "JWT_TOKEN", false)
		mockList := []string{"0_3", "1_1"}
		mockLastIngest := "0_2"
		deps.mockTaskRepo.On("PerformJobPool", mock.Anything, mock.Anything).Return(mockJob, nil).Once()
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockTaskRepo.On("DownloadSST", context.TODO(), mock.Anything).Return(false, errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)

		err := usecase.Perform(context.TODO(), "db1", false)

		assert.Equal(t, "download from cloud storage failed", err.Error())

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when ingest failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockJob := domain.NewJob("db1", "sst_example", "JWT_TOKEN", false)
		mockList := []string{"0_3", "1_1"}
		mockLastIngest := "0_2"
		deps.mockTaskRepo.On("PerformJobPool", mock.Anything, mock.Anything).Return(mockJob, nil).Once()
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockTaskRepo.On("DownloadSST", context.TODO(), mock.Anything).Return(true, nil).Once()
		deps.mockTaskRepo.On("Ingest", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)

		err := usecase.Perform(context.TODO(), "db1", false)

		assert.Equal(t, "ingest data failed", err.Error())

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when remove temp folder failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockJob := domain.NewJob("db1", "sst_example", "JWT_TOKEN", false)
		mockList := []string{"0_3", "1_1"}
		mockLastIngest := "0_2"
		deps.mockTaskRepo.On("PerformJobPool", mock.Anything, mock.Anything).Return(mockJob, nil).Once()
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockTaskRepo.On("DownloadSST", context.TODO(), mock.Anything).Return(true, nil).Once()
		deps.mockTaskRepo.On("Ingest", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		deps.mockFile.On("Remove", mock.Anything).Return(errors.New("Unexpected Error")).Once()
		deps.mockLog.On("Error", mock.Anything)

		err := usecase.Perform(context.TODO(), "db1", false)

		assert.Equal(t, "remove temp folder failed", err.Error())

		deps.mockTaskRepo.AssertExpectations(t)
	})

	t.Run("It should return err when job failed", func(t *testing.T) {
		usecase, deps := createTaskUsecase()
		mockJob := domain.NewJob("db1", "sst_example", "JWT_TOKEN", false)
		mockList := []string{"0_3", "1_1"}
		mockLastIngest := "0_2"
		deps.mockTaskRepo.On("PerformJobPool", mock.Anything, mock.Anything).Return(mockJob, nil).Once()
		deps.mockTaskRepo.On("GetLists", context.TODO(), mock.Anything).Return(mockList, nil).Once()
		deps.mockTaskRepo.On("GetLastIngest", context.TODO(), mock.Anything, mock.Anything).Return(mockLastIngest, nil).Once()
		deps.mockTaskRepo.On("DownloadSST", context.TODO(), mock.Anything).Return(false, nil).Once()
		deps.mockLog.On("Error", mock.Anything)

		err := usecase.Perform(context.TODO(), "db1", false)

		assert.Equal(t, "job failed", err.Error())

		deps.mockTaskRepo.AssertExpectations(t)
	})
}

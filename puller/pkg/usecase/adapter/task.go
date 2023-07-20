package adapter

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	domain "github.com/thnkrn/comet/puller/pkg/domain"
	file "github.com/thnkrn/comet/puller/pkg/driver/file"
	log "github.com/thnkrn/comet/puller/pkg/driver/log"
	repository "github.com/thnkrn/comet/puller/pkg/repository"
	usecase "github.com/thnkrn/comet/puller/pkg/usecase"
	usecaseError "github.com/thnkrn/comet/puller/pkg/usecase/error"
	util "github.com/thnkrn/comet/puller/pkg/utils"
)

const APITokenPrefix = "Bearer"

type TaskUsecase struct {
	taskRepo repository.TaskRepository
	file     file.File
	jobPool  domain.JobPools
	log      log.Logger
}

func NewTaskUsecase(taskRepo repository.TaskRepository, file file.File, jobPool domain.JobPools, log log.Logger) usecase.TaskUsecase {
	return &TaskUsecase{taskRepo, file, jobPool, log}
}

func (tc *TaskUsecase) GetLatestSeries(objects []string, ingestOnlyLatestDirectory bool) ([]string, error) {
	latestSeries := []string{}
	sort.Strings(objects)
	if ingestOnlyLatestDirectory {
		return []string{objects[len(objects)-1]}, nil
	}
	markedPrefix, err := util.GetBaseWithSplitter(objects[len(objects)-1])
	if err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot get base with splitter: %v", err))
		return nil, err
	}
	for _, s := range objects {
		if strings.Contains(s, markedPrefix) {
			latestSeries = append(latestSeries, s)
		}
	}
	return latestSeries, nil
}

func (tc *TaskUsecase) GetTargetStorageObjects(ctx context.Context, dbName, source, authorization string, ingestOnlyLatestDirectory, ignoreLastIngest bool) ([]string, error) {
	allStorageObjects, err := tc.taskRepo.GetLists(ctx, source)

	if len(allStorageObjects) == 0 || err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot find object due to: %v", err))
		return nil, fmt.Errorf("cannot find object for db %v from %v", dbName, source)
	}

	lastIngestDirectory, err := tc.taskRepo.GetLastIngest(ctx, authorization, dbName)
	if err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot get last ingest: %v", err))
		return nil, err
	}
	if !ignoreLastIngest && len(lastIngestDirectory) > 0 {
		markedPrefix, err := util.GetBase(lastIngestDirectory)
		if err != nil {
			tc.log.Error(fmt.Sprintf("usecase: cannot get base: %v", err))
			return nil, err
		}
		candidate := []string{}
		for _, s := range allStorageObjects {
			baseValue, err := util.GetBase(s)
			if err != nil {
				tc.log.Error(fmt.Sprintf("usecase: cannot get base: %v", err))
				return nil, err
			}
			if baseValue > markedPrefix {
				candidate = append(candidate, strconv.Itoa(baseValue))
			}
		}
		if len(candidate) > 0 {
			result, err := tc.GetLatestSeries(allStorageObjects, ingestOnlyLatestDirectory)
			if err != nil {
				tc.log.Error(fmt.Sprintf("usecase: cannot get latest series due to: %v", err))
				return nil, err
			}
			return result, err
		} else {
			result := []string{}
			current, err := util.GetIncremental(lastIngestDirectory)
			if err != nil {
				tc.log.Error(fmt.Sprintf("usecase: cannot get incremental: %v", err))
				return nil, err
			}
			for _, s := range allStorageObjects {
				base, err := util.GetBaseWithSplitter(lastIngestDirectory)
				if err != nil {
					tc.log.Error(fmt.Sprintf("usecase: cannot get base with splitter: %v", err))
					return nil, err
				}
				incremental, err := util.GetIncremental(s)
				if err != nil {
					tc.log.Error(fmt.Sprintf("usecase: cannot get incremental: %v", err))
					return nil, err
				}
				if strings.Contains(s, base) && incremental > current {
					result = append(result, s)
				}
			}
			return result, nil
		}
	} else {
		result, err := tc.GetLatestSeries(allStorageObjects, ingestOnlyLatestDirectory)
		if err != nil {
			tc.log.Error(fmt.Sprintf("usecase: cannot get latest series: %v", err))
			return nil, err
		}
		return result, nil
	}
}

func (tc *TaskUsecase) Ingest(ctx context.Context, authorization, dbName, directory string) (bool, error) {
	err := tc.taskRepo.Ingest(ctx, authorization, dbName, util.GetLastPath(directory))
	if err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot ingest due to: %v", err))
		return false, err
	}
	return true, nil
}

func (tc *TaskUsecase) Download(ctx context.Context, gcsObject string) (bool, error) {
	result, err := tc.taskRepo.DownloadSST(ctx, gcsObject)
	if err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot download due to: %v", err))
		return false, err
	}
	return result, nil
}

func (tc *TaskUsecase) RemoveTemp(directory string) error {
	err := tc.file.Remove(directory)
	if err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot remove temp directory due to: %v", err))
		return err
	}
	return nil
}

func (tc *TaskUsecase) Perform(ctx context.Context, db string, ignoreLastIngest bool) error {
	tc.log.Info(fmt.Sprintf("usecase: start perform job pool: %v", db))
	job, err := tc.taskRepo.PerformJobPool(tc.jobPool, db)
	if err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot perform job pool due to: %v", err))
		return usecaseError.NewErrorBusinessException(errors.New("couldn't perform job pool"))
	}
	tc.log.Info(fmt.Sprintf("usecase: get target storage objects to %v from %v", job.DBName, job.Source))
	storageObjects, err := tc.GetTargetStorageObjects(ctx, job.DBName, job.Source, job.ApiToken, job.IngestOnlyLastDir, ignoreLastIngest)
	if err != nil {
		tc.log.Error(fmt.Sprintf("usecase: cannot get target storage objects due to: %v", err))
		return usecaseError.NewErrorBusinessException(errors.New("couldn't get target storage objects"))
	}
	for _, s := range storageObjects {
		shouldIngest := job.ShouldIngest()
		if shouldIngest || ignoreLastIngest {
			job.SetIngesting()
			tc.log.Info(fmt.Sprintf("usecase: downloading objects from %v", util.AddTrailingSlash(job.Source)+s))
			downloadResult, err := tc.Download(ctx, util.AddTrailingSlash(job.Source)+s)
			if err != nil {
				tc.log.Error(fmt.Sprintf("usecase: cannot download due to: %v", err))
				return usecaseError.NewErrorBusinessException(errors.New("download from cloud storage failed"))
			}
			if downloadResult {
				tc.log.Info(fmt.Sprintf("usecase: finished download objects from %v", util.AddTrailingSlash(job.Source)+s))
				result, err := tc.Ingest(ctx, job.ApiToken, job.DBName, s)
				tc.log.Info(fmt.Sprintf("usecase: start ingest %v to %v", s, job.DBName))
				if err != nil {
					tc.log.Error(fmt.Sprintf("usecase: cannot ingest data due to: %v", err))
					return usecaseError.NewErrorBusinessException(errors.New("ingest data failed"))
				}
				if !result {
					tc.log.Error(fmt.Sprintf("usecase: cannot ingest data due to: %v", err))
					return usecaseError.NewErrorBusinessException(errors.New("ingest data failed"))
				}
				tc.log.Info(fmt.Sprintf("usecase: ingest %v to %v successfully", s, job.DBName))
				tc.log.Info("usecase: start removing downloaded objects in local")
				err = tc.RemoveTemp(util.AddTrailingSlash(job.Source) + s)
				if err != nil {
					tc.log.Error(fmt.Sprintf("usecase: cannot remove temp directory due to: %v", err))
					return usecaseError.NewErrorBusinessException(errors.New("remove temp folder failed"))
				}
				tc.log.Info("usecase: removing downloaded objects in local succesfully")
				job.SetSuccess()
			} else {
				job.SetFailed()
				tc.log.Error(fmt.Sprint("usecase: job failed", err))
				return usecaseError.NewErrorBusinessException(errors.New("job failed"))
			}
		}
	}
	return nil
}

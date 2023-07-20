package adapter

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"cloud.google.com/go/storage"
	config "github.com/thnkrn/comet/puller/pkg/config"
	domain "github.com/thnkrn/comet/puller/pkg/domain"
	file "github.com/thnkrn/comet/puller/pkg/driver/file"
	log "github.com/thnkrn/comet/puller/pkg/driver/log"
	rocksdbAdmin "github.com/thnkrn/comet/puller/pkg/driver/rocksdb_admin"
	repository "github.com/thnkrn/comet/puller/pkg/repository"
	util "github.com/thnkrn/comet/puller/pkg/utils"
	"google.golang.org/api/iterator"
)

type TaskRepository struct {
	rocksdbAdmin rocksdbAdmin.RocksdbAdmin
	file         file.File
	cfg          config.Config
	gcsClient    *storage.Client
	log          log.Logger
}

func NewTaskRepository(rocksdbAdmin rocksdbAdmin.RocksdbAdmin, file file.File, cfg config.Config, gcsClient *storage.Client, log log.Logger) repository.TaskRepository {
	return &TaskRepository{rocksdbAdmin, file, cfg, gcsClient, log}
}

func (tr *TaskRepository) Ingest(ctx context.Context, authorization, db, directory string) error {
	err := tr.rocksdbAdmin.Ingest(authorization, db, directory)
	if err != nil {
		tr.log.Error(fmt.Sprintf("repository: cannot ingest data due to: %v", err))
		return err
	}
	return nil
}

func (tr *TaskRepository) GetLastIngest(ctx context.Context, authorization, db string) (string, error) {
	result, err := tr.rocksdbAdmin.GetLastIngest(authorization, db)
	result = strings.Replace(result, "\"", "", -1)
	if err != nil {
		tr.log.Error(fmt.Sprintf("repository: cannot get last ingest due to: %v", err))
		return "", err
	}
	return result, nil
}

func (tr *TaskRepository) GetLists(ctx context.Context, prefix string) ([]string, error) {
	result := []string{}
	it := tr.gcsClient.Bucket(tr.cfg.GoogleCloud.BucketName).Objects(ctx, &storage.Query{
		Prefix:    util.AddTrailingSlash(prefix),
		Delimiter: "/",
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			tr.log.Error(fmt.Sprintf("repository: get object attributes failed due to: %v", err))
			return nil, err
		}
		blobName := attrs.Prefix
		if len(blobName) > 0 {
			blobName = blobName[:len(blobName)-1]
			seriesName := util.GetLastPath(blobName)
			r, err := regexp.Compile(`\d+_+\d`)
			if err != nil {
				tr.log.Error(fmt.Sprintf("repository: compile regex failed due to: %v", err))
				return nil, err
			}
			if r.MatchString(seriesName) {
				result = append(result, seriesName)
			}
		}
	}
	defer tr.gcsClient.Close()
	return result, nil
}

func (tr *TaskRepository) DownloadSST(ctx context.Context, gcsObject string) (bool, error) {
	var storageQuery = storage.Query{
		Prefix: util.AddTrailingSlash(gcsObject),
	}
	it := tr.gcsClient.Bucket(tr.cfg.GoogleCloud.BucketName).Objects(ctx, &storageQuery)
	isReady := false
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		isReady = strings.Contains(attrs.Name, "_SUCCESS")
		if isReady {
			break
		}
		if err != nil {
			tr.log.Error(fmt.Sprintf("repository: get object attributes failed due to: %v", err))
			return false, err
		}
	}
	if !isReady {
		tr.log.Error(fmt.Sprintf("repository: skip download from %v as success flag not found", gcsObject))
		return false, fmt.Errorf("skip download from %v as success flag not found ", gcsObject)
	}

	localPath, err := tr.file.CreateDirectoryInStaging(gcsObject)
	if err != nil {
		tr.log.Error(fmt.Sprintf("repository: create directory in staging failed due to: %v", err))
		return false, err
	}

	it = tr.gcsClient.Bucket(tr.cfg.GoogleCloud.BucketName).Objects(ctx, &storageQuery)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if strings.HasSuffix(attrs.Name, ".sst") {
			rc, err := tr.gcsClient.Bucket(tr.cfg.GoogleCloud.BucketName).Object(attrs.Name).NewReader(ctx)
			if err != nil {
				tr.log.Error(fmt.Sprintf("repository: get object from cloud storage failed due to: %v", err))
				return false, err
			}
			f, err := os.Create(util.AddTrailingSlash(localPath) + util.GetLastPath(attrs.Name))
			if err != nil {
				tr.log.Error(fmt.Sprintf("repository: create file failed due to: %v", err))
				return false, err
			}
			if _, err := io.Copy(f, rc); err != nil {
				tr.log.Error(fmt.Sprintf("repository: copy file failed due to: %v", err))
				return false, fmt.Errorf("io.Copy: %v", err)
			}
			if err = f.Close(); err != nil {
				tr.log.Error(fmt.Sprintf("repository: close file failed due to: %v", err))
				return false, fmt.Errorf("f.Close: %v", err)
			}
			defer rc.Close()
		}
		if err != nil {
			tr.log.Error(fmt.Sprintf("repository: get object attributes failed due to: %v", err))
			return false, err
		}
	}
	tr.file.CreateSuccesFile(util.AddTrailingSlash(localPath))
	return true, nil
}

func (tr *TaskRepository) PerformJobPool(jobPool domain.JobPools, db string) (domain.Job, error) {
	job, err := jobPool.Perform(db)
	return job, err
}

package adapter

import (
	"context"
	"fmt"
	"os"

	"github.com/linxGnu/grocksdb"

	config "github.com/thnkrn/comet/api/pkg/config"
	domain "github.com/thnkrn/comet/api/pkg/domain"
	log "github.com/thnkrn/comet/api/pkg/driver/log"
	repository "github.com/thnkrn/comet/api/pkg/repository"
)

type devRepository struct {
	rocksdbPool *domain.RocksDBPool
	log         log.Logger
	cfg         config.Config
}

func NewDevRepository(rocksdbPool *domain.RocksDBPool, log log.Logger, cfg config.Config) repository.DevRepository {
	return &devRepository{rocksdbPool: rocksdbPool, log: log, cfg: cfg}
}

func (r *devRepository) AddValueToSSTFile(ctx context.Context, fileName, key string, value []byte) ([]byte, error) {
	envo, o := grocksdb.NewDefaultEnvOptions(), grocksdb.NewDefaultOptions()
	sstWriter := grocksdb.NewSSTFileWriter(envo, o)

	sstDir := fmt.Sprintf("%s/%s.sst", r.cfg.Comet.Debug.SSTWriterPath, fileName)

	err := sstWriter.Open(sstDir)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error open SST writer: %v", err))
		return value, err
	}

	err = sstWriter.Put([]byte(key), value)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error put SST writer: %v", err))
		return value, err
	}

	err = sstWriter.Finish()
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error close SST writer: %v", err))
		return value, err
	}

	return value, err
}

func (r *devRepository) PullFile(ctx context.Context, fileName, source, ingestFolder string) error {
	fileDir := fmt.Sprintf("%s/%s.sst", r.cfg.Comet.Debug.SSTWriterPath, fileName)
	ingestDir := fmt.Sprintf("%s/%s/%s/%s.sst", r.cfg.App.IngestPath, source, ingestFolder, fileName)

	err := os.Rename(fileDir, ingestDir)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error pull SST file from: %s to %s by %v", fileDir, ingestDir, err))
		return err
	}

	return err
}

func (r *devRepository) ListDB(ctx context.Context) []domain.DB {
	return r.rocksdbPool.ListDB()
}

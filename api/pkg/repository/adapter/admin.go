package adapter

import (
	"context"
	"fmt"

	"github.com/linxGnu/grocksdb"

	config "github.com/thnkrn/comet/api/pkg/config"
	domain "github.com/thnkrn/comet/api/pkg/domain"
	log "github.com/thnkrn/comet/api/pkg/driver/log"
	repository "github.com/thnkrn/comet/api/pkg/repository"
)

type adminRepository struct {
	rocksdbPool *domain.RocksDBPool
	log         log.Logger
	cfg         config.Config
}

const METADATA_LAST_INGEST_KEY = "\\xc3\\x28LAST_INGEST"

func NewAdminRepository(rocksdbPool *domain.RocksDBPool, log log.Logger, cfg config.Config) repository.AdminRepository {
	return &adminRepository{log: log, rocksdbPool: rocksdbPool, cfg: cfg}
}

func (r *adminRepository) CatchUpWithPrimary(ctx context.Context, dbName string) error {
	rdb, mode, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return err
	}

	if mode != domain.SECONDARY {
		r.log.Error(fmt.Sprintf("repository: invalid db open mode for: %s", dbName))
		return ErrInvalidDBMode
	}

	err = rdb.TryCatchUpWithPrimary()

	return err
}

func (r *adminRepository) GetDBProperty(ctx context.Context, dbName, property string) (string, error) {
	var result string

	rdb, _, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return result, err
	}

	result = rdb.GetProperty(property)

	return result, nil
}

func (r *adminRepository) CreateCheckPoint(ctx context.Context, dbName, dir string) error {
	rdb, _, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return err
	}

	checkPoint, err := rdb.NewCheckpoint()
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error new checkpoint: %v", err))
		return err
	}

	checkPointDir := fmt.Sprintf("%s/%s", r.cfg.App.BackupPath, dir)

	err = checkPoint.CreateCheckpoint(checkPointDir, 0)
	r.log.Info(fmt.Sprintf("create backup for: %s at backup directory: %s", dbName, checkPointDir))
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error create checkpoint: %v", err))
	}
	defer checkPoint.Destroy()

	return err
}

func (r *adminRepository) Ingest(ctx context.Context, dbName, dir string, filesPath []string) error {
	rdb, mode, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return err
	}

	if mode != domain.READ_WRITE {
		r.log.Error(fmt.Sprintf("repository: invalid db open mode for: %s", dbName))
		return ErrInvalidDBMode
	}

	r.log.Info(fmt.Sprintf("start ingest file at: %s", dbName))
	for _, file := range filesPath {
		r.log.Info(fmt.Sprintf("file to ingest: %s into %s", file, dbName))
	}

	io := grocksdb.NewDefaultIngestExternalFileOptions()
	err = rdb.IngestExternalFile(filesPath, io)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error ingest file: %v", err))
		return err
	}
	defer io.Destroy()

	wo := grocksdb.NewDefaultWriteOptions()
	err = rdb.Put(wo, []byte(METADATA_LAST_INGEST_KEY), []byte(dir))
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error insert last ingest key: %v", err))
		return err
	}
	defer wo.Destroy()

	r.log.Info("finish ingest file")

	return err
}

func (r *adminRepository) GetLastIngest(ctx context.Context, dbName string) ([]byte, error) {
	rdb, _, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return nil, err
	}

	ro := grocksdb.NewDefaultReadOptions()
	result, err := rdb.Get(ro, []byte(METADATA_LAST_INGEST_KEY))
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error get last ingest key: %v", err))
		return nil, err
	}

	value := make([]byte, len(result.Data()))
	copy(value, result.Data())

	defer ro.Destroy()
	defer result.Free()

	return value, err
}

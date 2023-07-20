package adapter

import (
	"context"
	"errors"
	"fmt"

	"github.com/linxGnu/grocksdb"

	domain "github.com/thnkrn/comet/api/pkg/domain"
	log "github.com/thnkrn/comet/api/pkg/driver/log"
	repository "github.com/thnkrn/comet/api/pkg/repository"
)

const NUMBER_OF_KEYS = "rocksdb.estimate-num-keys"

var ErrInvalidDBMode = errors.New("invalid db open mode")

type userRepository struct {
	rocksdbPool *domain.RocksDBPool
	log         log.Logger
}

func NewUserRepository(rocksdbPool *domain.RocksDBPool, log log.Logger) repository.UserRepository {
	return &userRepository{rocksdbPool: rocksdbPool, log: log}
}

func (r *userRepository) GetByKey(ctx context.Context, dbName, key string) ([]byte, error) {
	rdb, _, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return nil, err
	}

	ro := grocksdb.NewDefaultReadOptions()
	result, err := rdb.Get(ro, []byte(key))
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error get value: %v", err))
		return nil, err
	}

	value := make([]byte, len(result.Data()))
	copy(value, result.Data())

	defer ro.Destroy()
	defer result.Free()

	return value, err
}

func (r *userRepository) CreateByKey(ctx context.Context, dbName, key string, value []byte) ([]byte, error) {
	rdb, mode, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return nil, err
	}

	if mode != domain.READ_WRITE {
		r.log.Error(fmt.Sprintf("repository: invalid db open mode for: %s", dbName))
		return value, ErrInvalidDBMode
	}

	wo := grocksdb.NewDefaultWriteOptions()
	err = rdb.Put(wo, []byte(key), value)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error create value: %v", err))
	}
	defer wo.Destroy()

	return value, err
}

func (r *userRepository) DeleteByKey(ctx context.Context, dbName, key string) error {
	rdb, mode, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return err
	}

	if mode != domain.READ_WRITE {
		r.log.Error(fmt.Sprintf("repository: invalid db open mode for: %s", dbName))
		return ErrInvalidDBMode
	}

	wo := grocksdb.NewDefaultWriteOptions()
	err = rdb.Delete(wo, []byte(key))
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error delete value: %v", err))
	}
	defer wo.Destroy()

	return err
}

func (r *userRepository) Count(ctx context.Context, dbName string) (string, error) {
	var result string

	rdb, _, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return result, err
	}

	result = rdb.GetProperty(NUMBER_OF_KEYS)

	return result, nil
}

func (r *userRepository) GetLastIngest(ctx context.Context, dbName string) ([]byte, error) {
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

func (r *userRepository) MultiGet(ctx context.Context, dbName string, keys [][]byte) ([]string, error) {
	rdb, _, err := r.rocksdbPool.GetConnection(dbName)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error db connection: %v", err))
		return nil, err
	}

	ro := grocksdb.NewDefaultReadOptions()
	result, err := rdb.MultiGet(ro, keys...)
	if err != nil {
		r.log.Error(fmt.Sprintf("repository: error multi get value: %v", err))
		return nil, err
	}

	slices := make(grocksdb.Slices, len(result))
	copy(slices, result)

	value := make([]string, len(slices))
	for i, v := range slices {
		value[i] = string(v.Data())
	}

	defer ro.Destroy()
	defer result.Destroy()

	return value, err
}

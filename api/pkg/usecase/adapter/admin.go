package adapter

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"

	config "github.com/thnkrn/comet/api/pkg/config"
	log "github.com/thnkrn/comet/api/pkg/driver/log"
	repository "github.com/thnkrn/comet/api/pkg/repository"
	usecase "github.com/thnkrn/comet/api/pkg/usecase"
	usecaseError "github.com/thnkrn/comet/api/pkg/usecase/error"
)

type adminUsecase struct {
	adminRepo repository.AdminRepository
	cfg       config.Config
	log       log.Logger
}

const INGEST_FILE_EXTENSTION = ".sst"

func find(dir, ext string) []string {
	var filesPath []string

	filepath.WalkDir(dir, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(d.Name()) == ext {
			filesPath = append(filesPath, s)
		}

		return nil
	})

	return filesPath
}

func NewAdminUseCase(r repository.AdminRepository, cfg config.Config, log log.Logger) usecase.AdminUsecase {
	return &adminUsecase{adminRepo: r, cfg: cfg, log: log}
}

func (u *adminUsecase) CatchUpWithPrimary(ctx context.Context, dbName string) error {
	err := u.adminRepo.CatchUpWithPrimary(ctx, dbName)
	return err
}

func (u *adminUsecase) GetDBProperty(ctx context.Context, dbName, property string) (string, error) {
	result, err := u.adminRepo.GetDBProperty(ctx, dbName, property)
	if result == "" && err == nil {
		return result, usecaseError.NewErrorNotFound(errors.New("property not found"))
	}
	return result, err
}

func (u *adminUsecase) CreateCheckPoint(ctx context.Context, dbName, dir string) error {
	err := u.adminRepo.CreateCheckPoint(ctx, dbName, dir)
	return err
}

func (u *adminUsecase) Ingest(ctx context.Context, dbName, dir string) error {
	dbConfig := u.cfg.FindDatabaseConfig(dbName)
	if dbConfig == nil || dbConfig.Source == "" {
		u.log.Error(fmt.Sprintf("usecase: can not find database config or source config not found from database name: %s", dbName))
		return usecaseError.NewErrorNotFound(errors.New("database config not found or source config not found"))
	}

	ingestsDir := fmt.Sprintf("%s/%s/%s", u.cfg.App.IngestPath, dbConfig.Source, dir)
	u.log.Info(fmt.Sprintf("find ingest file at directory: %s", ingestsDir))

	filesPath := find(ingestsDir, INGEST_FILE_EXTENSTION)

	if filesPath == nil || len(filesPath) <= 0 {
		u.log.Error(fmt.Sprintf("usecase: ingest file not found at directory: %s", ingestsDir))
		return usecaseError.NewErrorNotFound(errors.New("sst files not found"))
	}

	err := u.adminRepo.Ingest(ctx, dbName, dir, filesPath)
	return err
}

func (u *adminUsecase) GetLastIngest(ctx context.Context, dbName string) ([]byte, error) {
	result, err := u.adminRepo.GetLastIngest(ctx, dbName)
	if err == nil && (result == nil || len(result) <= 0) {
		return result, usecaseError.NewErrorNotFound(errors.New("value not found"))
	}
	return result, err
}

package adapter

import (
	"os"

	config "github.com/thnkrn/comet/puller/pkg/config"
	util "github.com/thnkrn/comet/puller/pkg/utils"
)

const successFile = "_SUCCESS"

type File struct {
	cfg config.Config
}

func NewFile(cfg config.Config) *File {
	return &File{cfg}
}

func (f *File) CreateDirectoryInStaging(newDirectory string) (string, error) {
	err := os.MkdirAll(util.AddTrailingSlash(f.cfg.App.IngestStagingPath)+newDirectory, os.ModePerm)
	if err != nil {
		return "", err
	}
	return util.AddTrailingSlash(f.cfg.App.IngestStagingPath) + newDirectory, nil
}

func (f *File) CreateSuccesFile(directory string) error {
	successFile, err := os.Create(directory + successFile)
	if err != nil {
		return err
	}
	successFile.Close()
	return nil
}

func (f *File) Remove(directory string) error {
	err := os.RemoveAll(util.AddTrailingSlash(f.cfg.App.IngestStagingPath) + directory)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) IsExistsWithSuccessFlag(path string) (bool, error) {
	if _, err := os.Stat(util.AddTrailingSlash(f.cfg.App.IngestStagingPath) + util.AddTrailingSlash(path) + successFile); err != nil {
		return false, err
	}
	return true, nil
}

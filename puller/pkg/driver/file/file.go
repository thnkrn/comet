package file

type File interface {
	CreateDirectoryInStaging(newDirectory string) (string, error)
	CreateSuccesFile(directory string) error
	Remove(directory string) error
	IsExistsWithSuccessFlag(path string) (bool, error)
}

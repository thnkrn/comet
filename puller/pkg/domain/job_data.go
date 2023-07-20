package domain

type JobData struct {
	DBNameKey         string
	StorageSourceKey  string
	APITokenKey       string
	IngestOnlyLastDir bool
	IgnoreLastIngest  bool
}

func NewJobData(dbName, storageSource, apiToken string, ingestFlag, ignoreLastIngest bool) JobData {
	return JobData{
		DBNameKey:         dbName,
		StorageSourceKey:  storageSource,
		APITokenKey:       apiToken,
		IngestOnlyLastDir: ingestFlag,
		IgnoreLastIngest:  ignoreLastIngest,
	}
}

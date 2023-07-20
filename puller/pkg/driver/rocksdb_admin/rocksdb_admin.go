package rocksdb_admin

type RocksdbAdmin interface {
	Ingest(authorization, db, directory string) error
	GetLastIngest(authorization, db string) (string, error)
}

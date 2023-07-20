package domain

import (
	"fmt"
	"sync"

	"github.com/linxGnu/grocksdb"
)

type RocksDBPool struct {
	mu    sync.Mutex
	dbs   map[string]*grocksdb.DB
	paths map[string]string
	opts  map[string]*grocksdb.Options
	modes map[string]string
}

type DB struct {
	Name string
	Mode string
}

const (
	READ_ONLY  = "read"
	READ_WRITE = "read-write"
	SECONDARY  = "secondary"
)

func NewRocksDBPool(paths, modes, primaryRef map[string]string, opts map[string]*grocksdb.Options) (*RocksDBPool, error) {
	dbs := make(map[string]*grocksdb.DB, len(paths))

	for name, path := range paths {
		dbOpts := opts[name]
		dbMode := modes[name]

		switch dbMode {
		case READ_ONLY:
			db, err := grocksdb.OpenDbForReadOnly(dbOpts, path, false)
			if err != nil {
				return nil, err
			}
			dbs[name] = db

		case READ_WRITE:
			db, err := grocksdb.OpenDb(dbOpts, path)
			if err != nil {
				return nil, err
			}
			dbs[name] = db

		case SECONDARY:
			dbPrimaryRef, ok := primaryRef[name]
			if !ok {
				return nil, fmt.Errorf("primary ref for '%s' does not exist", path)
			}
			db, err := grocksdb.OpenDbAsSecondary(dbOpts, dbPrimaryRef, path)
			if err != nil {
				return nil, err
			}
			dbs[name] = db
		}
	}

	return &RocksDBPool{
		dbs:   dbs,
		paths: paths,
		opts:  opts,
		modes: modes,
	}, nil
}

// Get connection and mode of database
func (p *RocksDBPool) GetConnection(name string) (*grocksdb.DB, string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.dbs == nil {
		return nil, "", fmt.Errorf("pool is closed")
	}

	db, ok := p.dbs[name]
	if !ok {
		return nil, "", fmt.Errorf("database '%s' does not exist", name)
	}

	return db, p.modes[name], nil
}

// Close database connections pool
func (p *RocksDBPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.dbs != nil {
		for _, db := range p.dbs {
			db.Close()
		}
		p.dbs = nil
	}

	return nil
}

// List all connected databases
func (p *RocksDBPool) ListDB() []DB {
	dbList := make([]DB, len(p.modes))
	idx := 0

	for name, mode := range p.modes {
		dbList[idx] = DB{
			Name: name,
			Mode: mode,
		}
		idx++
	}

	return dbList
}

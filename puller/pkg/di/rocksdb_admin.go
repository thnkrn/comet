package di

import (
	"github.com/google/wire"
	rocksdbAdmin "github.com/thnkrn/comet/puller/pkg/driver/rocksdb_admin"
	adapterRocksdbAdmin "github.com/thnkrn/comet/puller/pkg/driver/rocksdb_admin/adapter"
)

var RocksDBAdminSet = wire.NewSet(
	wire.Bind(new(rocksdbAdmin.RocksdbAdmin), new(*adapterRocksdbAdmin.RocksdbAdmin)),
	adapterRocksdbAdmin.NewRocksdbAdmin,
)

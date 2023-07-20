package di

import (
	"github.com/google/wire"

	db "github.com/thnkrn/comet/api/pkg/driver/db"
)

var ConnectionSet = wire.NewSet(db.ConnectRocksdb)

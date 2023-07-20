package di

import (
	"github.com/google/wire"
	file "github.com/thnkrn/comet/puller/pkg/driver/file"
	adapterfile "github.com/thnkrn/comet/puller/pkg/driver/file/adapter"
)

var FileSet = wire.NewSet(
	wire.Bind(new(file.File), new(*adapterfile.File)),
	adapterfile.NewFile,
)

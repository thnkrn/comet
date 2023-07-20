package types

import "github.com/linxGnu/grocksdb"

type OptsModifier func(opts *grocksdb.Options)

type BBOptsModifier func(bbopts *grocksdb.BlockBasedTableOptions)

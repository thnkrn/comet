package db

import (
	"github.com/linxGnu/grocksdb"

	types "github.com/thnkrn/comet/api/pkg/types"
)

func WithSetBlockSize(v *int) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bbopts.SetBlockSize(*v * int(types.SizeUnit.KB))
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetBlockSize(4 * int(types.SizeUnit.KB))
	}
}

func WithSetBlockCache(v *int) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			lruc := grocksdb.NewLRUCache(uint64(*v * types.SizeUnit.MB))
			defer lruc.Destroy()
			bbopts.SetBlockCache(lruc)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {}
}

func WithSetFilterPolicy(v *float64) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bf := grocksdb.NewRibbonFilterPolicy(*v)
			defer bf.Destroy()
			bbopts.SetFilterPolicy(bf)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bf := grocksdb.NewRibbonFilterPolicy(10)
		defer bf.Destroy()
		bbopts.SetFilterPolicy(bf)
	}
}

func WithSetCacheIndexAndFilterBlocks(v *bool) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bbopts.SetCacheIndexAndFilterBlocks(*v)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetCacheIndexAndFilterBlocks(true)
	}
}

func WithSetCacheIndexAndFilterBlocksWithHighPriority(v *bool) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bbopts.SetCacheIndexAndFilterBlocksWithHighPriority(*v)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetCacheIndexAndFilterBlocksWithHighPriority(true)
	}
}

func WithSetNoBlockCache(v *bool) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bbopts.SetNoBlockCache(*v)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetNoBlockCache(false)
	}
}

func WithSetPinL0FilterAndIndexBlocksInCache(v *bool) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bbopts.SetPinL0FilterAndIndexBlocksInCache(*v)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetPinL0FilterAndIndexBlocksInCache(true)
	}
}

func WithSetPinTopLevelIndexAndFilter(v *bool) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bbopts.SetPinTopLevelIndexAndFilter(*v)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetPinTopLevelIndexAndFilter(true)
	}
}

func WithSetFormatVersion(v *int) types.BBOptsModifier {
	if v != nil {
		return func(bbopts *grocksdb.BlockBasedTableOptions) {
			bbopts.SetFormatVersion(*v)
		}
	}
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetFormatVersion(4)
	}
}

func WithSetUseDeltaEncoding() types.BBOptsModifier {
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetUseDeltaEncoding(true)
	}
}

func WithSetWholeKeyFiltering() types.BBOptsModifier {
	return func(bbopts *grocksdb.BlockBasedTableOptions) {
		bbopts.SetWholeKeyFiltering(true)
	}
}

func BlockBasedOptions(bbopts ...types.BBOptsModifier) *grocksdb.BlockBasedTableOptions {
	bbOptions := grocksdb.NewDefaultBlockBasedTableOptions()

	// NOTE: Execute block based options
	for _, mod := range bbopts {
		mod(bbOptions)
	}

	return bbOptions
}

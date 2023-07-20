package db

import (
	"runtime"

	"github.com/linxGnu/grocksdb"

	"github.com/thnkrn/comet/api/pkg/config"
	types "github.com/thnkrn/comet/api/pkg/types"
)

const (
	SNAPPY_COMPRESSION = "SNAPPY_COMPRESSION"
	ZLIB_COMPRESSION   = "ZLIB_COMPRESSION"
	BZ2_COMPRESSION    = "BZ2_COMPRESSION"
	LZ4_COMPRESSION    = "LZ4_COMPRESSION"
	LZ4H_COMPRESSION   = "LZ4H_COMPRESSION"
	XPRESS_COMPRESSION = "XPRESS_COMPRESSION"
	ZSTD_COMPRESSION   = "ZSTD_COMPRESSION"
	NO_COMPRESSION     = "NO_COMPRESSION"

	LEVEL = "LEVEL"
	FIFO  = "FIFO"
)

func getCompressionType(s string) grocksdb.CompressionType {
	switch s {
	case SNAPPY_COMPRESSION:
		return grocksdb.SnappyCompression
	case ZLIB_COMPRESSION:
		return grocksdb.ZLibCompression
	case BZ2_COMPRESSION:
		return grocksdb.Bz2Compression
	case LZ4_COMPRESSION:
		return grocksdb.LZ4Compression
	case LZ4H_COMPRESSION:
		return grocksdb.LZ4HCCompression
	case XPRESS_COMPRESSION:
		return grocksdb.XpressCompression
	case ZSTD_COMPRESSION:
		return grocksdb.ZSTDCompression
	case NO_COMPRESSION:
		return grocksdb.NoCompression
	default:
		return grocksdb.SnappyCompression
	}
}

func getCompactStyle(s string) grocksdb.CompactionStyle {
	switch s {
	case LEVEL:
		return grocksdb.LevelCompactionStyle
	case FIFO:
		return grocksdb.FIFOCompactionStyle
	default:
		return grocksdb.UniversalCompactionStyle
	}
}

func WithCompression(v string) types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetCompression(getCompressionType(v))
	}
}

func WithCompactionStyle(v string) types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetCompactionStyle(getCompactStyle(v))
	}
}

func WithCreateIfMissing(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetCreateIfMissing(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetCreateIfMissing(false)
	}
}

func WithWriteBufferSize(v *int) types.OptsModifier {
	if v != nil && *v >= 0 {
		return func(opts *grocksdb.Options) {
			opts.SetWriteBufferSize(uint64(*v * types.SizeUnit.MB))
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetWriteBufferSize(uint64(8 * types.SizeUnit.MB))
	}
}

func WithMaxWriteBufferSize(v *int) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetMaxWriteBufferNumber(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetMaxWriteBufferNumber(128)
	}
}

func WithIncreaseParallelism(v *int) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.IncreaseParallelism(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.IncreaseParallelism(runtime.NumCPU())
	}
}

func WithOptimizeUniversalStyleCompaction(v *int) types.OptsModifier {
	if v != nil && *v >= 0 {
		return func(opts *grocksdb.Options) {
			opts.OptimizeUniversalStyleCompaction(uint64(*v))
		}
	}
	return func(opts *grocksdb.Options) {}
}

func WithCompactionReadaheadSize(v *int) types.OptsModifier {
	if v != nil && *v >= 0 {
		return func(opts *grocksdb.Options) {
			opts.CompactionReadaheadSize(uint64(*v * types.SizeUnit.MB))
		}
	}
	return func(opts *grocksdb.Options) {
		opts.CompactionReadaheadSize(uint64(2 * types.SizeUnit.MB))
	}
}

func WithOptimizeForPointLookup(v *int) types.OptsModifier {
	if v != nil && *v >= 0 {
		return func(opts *grocksdb.Options) {
			opts.OptimizeForPointLookup(uint64(*v))
		}
	}
	return func(opts *grocksdb.Options) {}
}

func WithSetOptimizeFiltersForHits(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetOptimizeFiltersForHits(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetOptimizeFiltersForHits(true)
	}
}

func WithSetUseDirectReads(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetUseDirectReads(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetUseDirectReads(false)
	}
}

func WithSetUseDirectIOForFlushAndCompaction(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetUseDirectIOForFlushAndCompaction(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetUseDirectIOForFlushAndCompaction(false)
	}
}

func WithSetMaxBackgroundJobs(v *int) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetMaxBackgroundJobs(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetMaxBackgroundJobs(4)
	}
}

func WithSetDisableAutoCompactions(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetDisableAutoCompactions(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetDisableAutoCompactions(false)
	}
}

func WithSetStatsDumpPeriodSec(v *int) types.OptsModifier {
	if v != nil && *v >= 0 {
		return func(opts *grocksdb.Options) {
			opts.SetStatsDumpPeriodSec(uint(*v))
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetStatsDumpPeriodSec(3600)
	}
}

func WithSetNumLevels(v *int) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetNumLevels(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetNumLevels(7)
	}
}

func WithSetAllowMmapReads(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetAllowMmapReads(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetAllowMmapReads(false)
	}
}

func WithSetAllowMmapWrites(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetAllowMmapWrites(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetAllowMmapWrites(false)
	}
}

func WithSetMaxLogFileSize(v *int) types.OptsModifier {
	if v != nil && *v >= 0 {
		return func(opts *grocksdb.Options) {
			opts.SetMaxLogFileSize(uint64(*v))
		}
	}
	return func(opts *grocksdb.Options) {}
}

func WithSetSkipStatsUpdateOnDBOpen(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetSkipStatsUpdateOnDBOpen(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetSkipStatsUpdateOnDBOpen(true)
	}
}

func WithSetSkipCheckingSSTFileSizesOnDBOpen(v *bool) types.OptsModifier {
	if v != nil {
		return func(opts *grocksdb.Options) {
			opts.SetSkipCheckingSSTFileSizesOnDBOpen(*v)
		}
	}
	return func(opts *grocksdb.Options) {
		opts.SetSkipCheckingSSTFileSizesOnDBOpen(true)
	}
}

func WithSetMaxOpenFiles() types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetMaxOpenFiles(-1)
	}
}

func WithSetRecycleLogFileNum() types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetRecycleLogFileNum(0)
	}
}

func WithSetTargetFileSizeBase() types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetTargetFileSizeBase(uint64(64 * types.SizeUnit.MB))
	}
}

func WithSetMaxBytesForLevelBase() types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetMaxBytesForLevelBase(uint64(1 * types.SizeUnit.GB))
	}
}

func WithSetMaxTotalWalSize() types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetMaxTotalWalSize(uint64(1 * types.SizeUnit.GB))
	}
}

func WithSetCompressionOptionsMaxDictBufferBytes() types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetCompressionOptionsMaxDictBufferBytes(0)
	}
}

func WithSetMaxFileOpeningThreads() types.OptsModifier {
	return func(opts *grocksdb.Options) {
		opts.SetMaxFileOpeningThreads(16)
	}
}

func BuildOptions(bbs *config.BlockBasedSettings, opts ...types.OptsModifier) *grocksdb.Options {
	options := grocksdb.NewDefaultOptions()

	// NOTE: Execute build options
	for _, mod := range opts {
		mod(options)
	}

	bbOptions := BlockBasedOptions(
		WithSetBlockSize(bbs.BlockSizeKb),
		WithSetBlockCache(bbs.BlockCacheSizeMb),
		WithSetFilterPolicy(bbs.BloomFilterBits),
		WithSetCacheIndexAndFilterBlocks(bbs.CacheIndexAndFilterBlocks),
		WithSetCacheIndexAndFilterBlocksWithHighPriority(bbs.CacheIndexAndFilterBlocksWithHighPriority),
		WithSetNoBlockCache(bbs.DisableCache),
		WithSetPinL0FilterAndIndexBlocksInCache(bbs.PinL0FilterAndIndexBlocksInCache),
		WithSetPinTopLevelIndexAndFilter(bbs.PinTopLevelFilterAndIndex),
		WithSetFormatVersion(bbs.FormatVersion),

		WithSetUseDeltaEncoding(),
		WithSetWholeKeyFiltering(),
	)

	defer bbOptions.Destroy()

	options.SetBlockBasedTableFactory(bbOptions)

	return options
}

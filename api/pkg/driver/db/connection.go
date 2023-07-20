package db

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/linxGnu/grocksdb"

	config "github.com/thnkrn/comet/api/pkg/config"
	domain "github.com/thnkrn/comet/api/pkg/domain"
	log "github.com/thnkrn/comet/api/pkg/driver/log"
)

func ConnectRocksdb(cfg config.Config, log log.Logger) (*domain.RocksDBPool, error) {
	paths := make(map[string]string, len(cfg.App.Databases))
	modes := make(map[string]string, len(cfg.App.Databases))
	opts := make(map[string]*grocksdb.Options, len(cfg.App.Databases))
	primaryRef := make(map[string]string)

	for _, rdb := range cfg.App.Databases {
		// NOTE: Set default value for database config
		rdb.FillDefaults()

		if (cfg.Comet.Prefork == true) && (rdb.Mode != domain.READ_ONLY) {
			return nil, errors.New("Prefork mode can only use when open all databases as read only mode")
		}

		// NOTE: Create build options for open Database
		buildOptions := BuildOptions(
			rdb.BlockBasedSettings,
			WithCompression(rdb.CompressionType),
			WithCompactionStyle(rdb.CompactionStyle),
			WithCreateIfMissing(rdb.CreateIfMissing),
			WithWriteBufferSize(rdb.WriteBufferSizeMb),
			WithMaxWriteBufferSize(rdb.MaxWriteBufferNumber),
			WithIncreaseParallelism(rdb.IncreaseParallelism),
			WithOptimizeUniversalStyleCompaction(rdb.OptimizeUniversalStyleCompaction),
			WithCompactionReadaheadSize(rdb.CompactionReadAheadSizeMb),
			WithOptimizeForPointLookup(rdb.OptimizeForPointLookUp),
			WithSetOptimizeFiltersForHits(rdb.OptimizeFilterForHits),
			WithSetUseDirectReads(rdb.UseDirectReads),
			WithSetUseDirectIOForFlushAndCompaction(rdb.UseDirectIoForFlushAndCompaction),
			WithSetMaxBackgroundJobs(rdb.MaxBackgroundJobs),
			WithSetDisableAutoCompactions(rdb.DisableAutoCompaction),
			WithSetStatsDumpPeriodSec(rdb.StatsDumpPeriodSec),
			WithSetNumLevels(rdb.NumLevels),
			WithSetAllowMmapReads(rdb.AllowMmapReads),
			WithSetAllowMmapWrites(rdb.AllowMmapWrites),
			WithSetMaxLogFileSize(rdb.MaxLogFileSize),
			WithSetSkipStatsUpdateOnDBOpen(rdb.SkipStatsUpdateOnDBOpen),
			WithSetSkipCheckingSSTFileSizesOnDBOpen(rdb.SkipStatsUpdateOnDBOpen),

			// NOTE: performance tuing, REF: https://ceph.io/en/news/blog/2022/rocksdb-tuning-deep-dive/#performance-comparison, https://www.confluent.io/blog/how-to-tune-rocksdb-kafka-streams-state-stores-performance/
			WithSetMaxOpenFiles(),
			WithSetRecycleLogFileNum(),
			WithSetTargetFileSizeBase(),
			WithSetMaxBytesForLevelBase(),
			WithSetMaxTotalWalSize(),
			WithSetCompressionOptionsMaxDictBufferBytes(),
			WithSetMaxFileOpeningThreads(),
		)

		if rdb.Mode == domain.READ_ONLY {
			buildOptions.SetInfoLogLevel(grocksdb.ErrorInfoLogLevel)
		}

		defer buildOptions.Destroy()

		// NOTE: Create missing directory with permission
		if _, err := os.Stat(rdb.Path); err != nil {
			if os.IsNotExist(err) {
				log.Info(fmt.Sprintf("Create missing directory at: %s, with exception %s", rdb.Path, err))
				dir := filepath.Dir(rdb.Path)
				mkDirErr := os.MkdirAll(dir, 0777)
				if mkDirErr != nil {
					log.Error(fmt.Sprintf("Failed to mkdir at: %s, with exception %s", rdb.Path, mkDirErr))
					return nil, err
				}
			} else {
				log.Error(fmt.Sprintf("Failed to crearte directory at: %s, with exception %s", rdb.Path, err))
				return nil, err
			}
		}

		paths[rdb.Name] = rdb.Path
		modes[rdb.Name] = rdb.Mode
		opts[rdb.Name] = buildOptions

		if rdb.Mode == domain.SECONDARY && rdb.PrimaryRef != "" {
			primaryRef[rdb.Name] = rdb.PrimaryRef
		}
	}

	pool, err := domain.NewRocksDBPool(paths, modes, primaryRef, opts)
	// NOTE: Remove closing pool for now
	// defer pool.Close()

	return pool, err
}

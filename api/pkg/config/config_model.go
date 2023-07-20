package config

type Config struct {
	Comet CometConfig `mapstructure:"comet" json:"comet" validate:"required"`
	App   AppConfig   `mapstructure:"app" json:"app" validate:"required"`
}
type CometConfig struct {
	Prefork bool        `mapstructure:"prefork"`
	Recover bool        `mapstructure:"recover"`
	Log     LogConfig   `mapstructure:"log" validate:"required"`
	Debug   DebugConfig `mapstructure:"debug"`
}

type LogConfig struct {
	Tracing bool      `mapstructure:"tracing"`
	Zap     ZapConfig `mapstructure:"zap" validate:"required"`
}

type ZapConfig struct {
	Level            string        `mapstructure:"level" json:"level" validate:"required"`
	Encoding         string        `mapstructure:"encoding" json:"encoding" validate:"required"`
	OutputPaths      []string      `mapstructure:"outputPaths" json:"outputPaths" validate:"required"`
	ErrorOutputPaths []string      `mapstructure:"errorOutputPaths" json:"errorOutputPaths" validate:"required"`
	EncoderConfig    EncoderConfig `mapstructure:"encoderConfig" json:"encoderConfig" validate:"required"`
}

type EncoderConfig struct {
	MessageKey   string `mapstructure:"messageKey" json:"messageKey" validate:"required"`
	LevelKey     string `mapstructure:"levelKey" json:"levelKey" validate:"required"`
	LevelEncoder string `mapstructure:"levelEncoder" json:"levelEncoder" validate:"required"`
}

type DebugConfig struct {
	SSTWriterPath string `mapstructure:"sstWriterPath" json:"sstWriterPath"`
}

type AppConfig struct {
	BackupPath string      `mapstructure:"backupPath" json:"backupPath" validate:"required"`
	IngestPath string      `mapstructure:"ingestPath" json:"ingestPath" validate:"required"`
	Databases  []Databases `mapstructure:"databases" json:"databases" validate:"required,dive"`
}

type Databases struct {
	Path            string `mapstructure:"path" json:"path" validate:"required"`
	Name            string `mapstructure:"name" json:"name" validate:"required"`
	Mode            string `mapstructure:"mode" json:"mode" validate:"required"`
	Source          string `mapstructure:"source" json:"source"`
	PrimaryRef      string `mapstructure:"primaryRef" json:"primaryRef" validate:"required_if=Mode secondary"`
	CompressionType string `mapstructure:"compressionType" json:"compressionType" validate:"omitempty,oneof=SNAPPY_COMPRESSION ZLIB_COMPRESSION BZ2_COMPRESSION LZ4_COMPRESSION LZ4H_COMPRESSION XPRESS_COMPRESSION ZSTD_COMPRESSION NO_COMPRESSION"`
	CompactionStyle string `mapstructure:"compactionType" json:"compactionType" validate:"omitempty,oneof=LEVEL FIFO UNIVERSAL"`

	CreateIfMissing                     *bool `mapstructure:"createIfMissing" json:"createIfMissing"`
	OptimizeFilterForHits               *bool `mapstructure:"optimizeFilterForHits" json:"optimizeFilterForHits"`
	UseDirectReads                      *bool `mapstructure:"useDirectReads" json:"useDirectReads"`
	UseDirectIoForFlushAndCompaction    *bool `mapstructure:"useDirectIoForFlushAndCompaction" json:"useDirectIoForFlushAndCompaction"`
	DisableAutoCompaction               *bool `mapstructure:"disableAutoCompaction" json:"disableAutoCompaction"`
	AllowMmapReads                      *bool `mapstructure:"allowMmapReads" json:"allowMmapReads"`
	AllowMmapWrites                     *bool `mapstructure:"allowMmapWrites" json:"allowMmapWrites"`
	SkipStatsUpdateOnDBOpen             *bool `mapstructure:"skipStatsUpdateOnDbOpen" json:"skipStatsUpdateOnDbOpen"`
	SetSkipCheckingSSTFileSizesOnDBOpen *bool `mapstructure:"skipCheckingSstFileSizeOnDbOpen" json:"skipCheckingSstFileSizeOnDbOpen"`

	StatsDumpPeriodSec               *int `mapstructure:"statsDumpPeriodSec" json:"statsDumpPeriodSec" validate:"omitempty,gte=0"`
	WriteBufferSizeMb                *int `mapstructure:"writeBufferSizeMb" json:"writeBufferSizeMb" validate:"omitempty,gte=0"`
	OptimizeUniversalStyleCompaction *int `mapstructure:"optimizeUniversalStyleCompaction" json:"optimizeUniversalStyleCompaction" validate:"omitempty,gte=0"`
	CompactionReadAheadSizeMb        *int `mapstructure:"compactionReadAheadSizeMb" json:"compactionReadAheadSizeMb" validate:"omitempty,gte=0"`
	OptimizeForPointLookUp           *int `mapstructure:"optimizeForPointLookUp" json:"optimizeForPointLookUp" validate:"omitempty,gte=0"`
	MaxLogFileSize                   *int `mapstructure:"maxLogFileSize" json:"maxLogFileSize" validate:"omitempty,gte=0"`
	MaxWriteBufferNumber             *int `mapstructure:"maxWriteBufferNumber" json:"maxWriteBufferNumber" validate:"omitempty,gte=0"`
	IncreaseParallelism              *int `mapstructure:"increaseParallelism" json:"increaseParallelism" validate:"omitempty,gte=0"`
	MaxBackgroundJobs                *int `mapstructure:"maxBackgroundJobs" json:"maxBackgroundJobs" validate:"omitempty,gte=0"`
	NumLevels                        *int `mapstructure:"numLevels" json:"numLevels"  validate:"omitempty,gte=0"`

	BlockBasedSettings *BlockBasedSettings `mapstructure:"blockBasedSettings" json:"blockBasedSettings"`
}

func (dbs *Databases) FillDefaults() {
	if dbs.CompressionType == "" {
		dbs.CompressionType = "SNAPPY_COMPRESSION"
	}

	if dbs.CompactionStyle == "" {
		dbs.CompactionStyle = "UNIVERSAL"
	}
}

type BlockBasedSettings struct {
	BlockSizeKb      *int `mapstructure:"blockSizeKb" json:"blockSizeKb" validate:"omitempty,gte=0"`
	BlockCacheSizeMb *int `mapstructure:"blockCacheSizeMb" json:"blockCacheSizeMb" validate:"omitempty,gte=0"`
	FormatVersion    *int `mapstructure:"formatVersion" json:"formatVersion" validate:"omitempty"`

	BloomFilterBits *float64 `mapstructure:"bloomFilterBits" json:"bloomFilterBits"`

	DisableCache                              *bool `mapstructure:"disableCache" json:"disableCache"`
	CacheIndexAndFilterBlocks                 *bool `mapstructure:"cacheIndexAndFilterBlocks" json:"cacheIndexAndFilterBlocks"`
	CacheIndexAndFilterBlocksWithHighPriority *bool `mapstructure:"cacheIndexAndFilterBlocksWithHighPriority" json:"cacheIndexAndFilterBlocksWithHighPriority"`
	PinL0FilterAndIndexBlocksInCache          *bool `mapstructure:"pinL0FilterAndIndexBlocksInCache" json:"pinL0FilterAndIndexBlocksInCache"`
	PinTopLevelFilterAndIndex                 *bool `mapstructure:"pinTopLevelFilterAndIndex" json:"pinTopLevelFilterAndIndex"`
}

func (c *Config) FindDatabaseConfig(dbName string) *Databases {
	for i := 0; i < len(c.App.Databases); i++ {
		if c.App.Databases[i].Name == dbName {
			return &c.App.Databases[i]
		}
	}
	return nil
}

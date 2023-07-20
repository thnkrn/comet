package config

type Config struct {
	CometPuller CometPullerConfig `mapstructure:"cometPuller" json:"cometPuller" validate:"required"`
	App         AppConfig         `mapstructure:"app" json:"app" validate:"required"`
	GoogleCloud GoogleCloudConfig `mapstructure:"googleCloud" json:"googleCloud" validate:"required"`
}
type CometPullerConfig struct {
	Log LogConfig `mapstructure:"log" json:"log" validate:"required"`
}

type LogConfig struct {
	Tracing bool      `mapstructure:"tracing" json:"tracing"`
	Zap     ZapConfig `mapstructure:"zap" json:"zap" validate:"required"`
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

type AppConfig struct {
	APIUrl            string      `mapstructure:"apiUrl" json:"apiUrl" validate:"required"`
	APIToken          string      `mapstructure:"apiToken" json:"apiToken" validate:"required"`
	IngestStagingPath string      `mapstructure:"ingestStagingPath" json:"ingestStagingPath" validate:"required"`
	SSTBuckets        []SSTBucket `mapstructure:"sstBuckets" json:"sstBuckets" validate:"required"`
}

type SSTBucket struct {
	Name                string `mapstructure:"name" json:"name" validate:"required"`
	Source              string `mapstructure:"source" json:"source" validate:"required"`
	Schedule            string `mapstructure:"schedule" json:"schedule" validate:"required"`
	IngestOnlyLatestDir bool   `mapstructure:"ingestOnlyLatestDirectory" json:"ingestOnlyLatestDirectory" validate:"required"`
}

type GoogleCloudConfig struct {
	BucketName             string `mapstructure:"bucketName" json:"bucketName" validate:"required"`
	ProjectId              string `mapstructure:"projectId" json:"projectId" validate:"required"`
	ServiceAccountLocation string `mapstructure:"serviceAccountLocation" json:"serviceAccountLocation" validate:"required"`
}

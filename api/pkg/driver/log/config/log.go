package config

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	config "github.com/thnkrn/comet/api/pkg/config"
)

func ProvidZapLogger(cfg config.Config) (*zap.Logger, error) {
	jcfg, err := json.Marshal(cfg.Comet.Log.Zap)
	if err != nil {
		return nil, err
	}

	var zcfg zap.Config
	if err := json.Unmarshal(jcfg, &zcfg); err != nil {
		return nil, err
	}

	zcfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger := zap.Must(zcfg.Build())
	defer logger.Sync()

	return logger, nil
}

package logger

import (
	"errors"
	"syscall"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	LOG_LEVEL_KEY  = "log_level"
	LOG_FORMAT_KEY = "log_format"
)

type LogConfig struct {
	Level  string
	Format string
}

func LogPFlags() *pflag.FlagSet {
	f := pflag.NewFlagSet("log", pflag.ContinueOnError)
	f.String(LOG_LEVEL_KEY, "info", "log level")
	f.String(LOG_FORMAT_KEY, "json", "log format")

	return f
}

func LogConfigFromViper() LogConfig {
	return LogConfig{
		Level:  viper.GetString(LOG_LEVEL_KEY),
		Format: viper.GetString(LOG_FORMAT_KEY),
	}
}

func SetGlobalZapLogger(cfg LogConfig) {
	var logger *zap.Logger

	var zapCfg zap.Config

	if cfg.Format != "json" {
		zapCfg = zap.NewDevelopmentConfig()
	} else {
		zapCfg = zap.NewProductionConfig()
		zapCfg.EncoderConfig.TimeKey = "time"
	}

	err := zapCfg.Level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		panic(err)
	}

	logger, _ = zapCfg.Build()

	defer func() {
		if err := logger.Sync(); err != nil && !errors.Is(err, syscall.EINVAL) {
			panic(err)
		}
	}()

	zap.ReplaceGlobals(logger)

	zap.L().Info("logger config", zap.Dict("config",
		zap.String(LOG_LEVEL_KEY, cfg.Level),
		zap.String(LOG_FORMAT_KEY, cfg.Format),
	))
}

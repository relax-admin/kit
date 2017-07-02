package config

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func Read(env string, config interface{}) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if env != "" {
		f, err := os.Open("config." + env + ".yml")
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		defer f.Close()
		viper.MergeConfig(f)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

type LoggerConfig struct {
	Level             string
	DisableCaller     bool
	DisableStacktrace bool
	OutputPath        string
	ErrorOutputPath   string
}

func (config LoggerConfig) Build() *zap.Logger {
	var zapConfig zap.Config

	switch strings.ToLower(config.Level) {
	case "debug":
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.Level.SetLevel(zap.DebugLevel)
	default:
		zapConfig = zap.NewProductionConfig()
		switch strings.ToLower(config.Level) {
		case "info":
			zapConfig.Level.SetLevel(zap.InfoLevel)
		case "warn":
			zapConfig.Level.SetLevel(zap.WarnLevel)
		case "error":
			zapConfig.Level.SetLevel(zap.ErrorLevel)
		case "panic":
			zapConfig.Level.SetLevel(zap.PanicLevel)
		case "fatal":
			zapConfig.Level.SetLevel(zap.FatalLevel)
		}
	}

	if config.OutputPath != "" {
		zapConfig.OutputPaths = []string{config.OutputPath}
	}

	if config.ErrorOutputPath != "" {
		zapConfig.ErrorOutputPaths = []string{config.ErrorOutputPath}
	}

	zapConfig.DisableCaller = config.DisableCaller
	zapConfig.DisableStacktrace = config.DisableStacktrace

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

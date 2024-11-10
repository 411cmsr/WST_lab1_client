package logger

import (
	"go.uber.org/zap"
)

func NewLogger(config zap.Config) (*zap.Logger, error) {
	return config.Build()
}

func NewLoggerConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"soapclient.log"}
	return config
}

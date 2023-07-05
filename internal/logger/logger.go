package logger

import (
	"we-connect-test/config"

	"go.uber.org/zap"
)

func NewLogger(cfg *config.Cfg) (*zap.Logger, error) {
	return zap.NewProduction()
}

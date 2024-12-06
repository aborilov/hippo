package logger

import (
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(*zap.Config)

// NewLogger creates a new logr.Logger with the provided options.
func NewLogger(opts ...Option) (logr.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		zapcore.RFC3339TimeEncoder(time.UTC(), encoder)
	}
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.DisableStacktrace = false

	for _, opt := range opts {
		opt(&cfg)
	}

	zl, err := cfg.Build()
	if err != nil {
		return logr.Discard(), err
	}

	logger := zapr.NewLoggerWithOptions(zl)
	return logger, nil
}

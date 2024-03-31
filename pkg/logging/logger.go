package logging

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const loggerKey = contextKey("logger")

var (
	defaultLogger     *zap.SugaredLogger
	defaultLoggerOnce *sync.Once
)

var config = &Config{
	Encoding:    "console",
	Level:       zapcore.InfoLevel,
	Development: true,
}

type Config struct {
	Encoding    string
	Level       zapcore.Level
	Development bool
}

func SetConfig(c *Config) {
	config = &Config{Encoding: c.Encoding,
		Level:       c.Level,
		Development: c.Development,
	}
}

func SetLevel(level zapcore.Level) {
	config.Level = level
}

// Create a new logger with the appropriate log level.
func NewLogger(config *Config) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg := zap.Config{
		Encoding:         config.Encoding,
		EncoderConfig:    encoderConfig,
		Level:            zap.NewAtomicLevelAt(config.Level),
		Development:      config.Development,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	return logger.Sugar()
}

// DefaultLogger returns the default logger for the package.
func DefaultLogger() *zap.SugaredLogger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLogger(config)
	})
	return defaultLogger
}

// WithLogger creates a new context with the provided logger attached.
func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	if gCtx, ok := ctx.(*gin.Context); ok {
		ctx = gCtx.Request.Context()
	}
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext attempts to return the logger stored in the context. If no
// logger exists, a default logger is returned.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return DefaultLogger()
	}
	if gCtx, ok := ctx.(*gin.Context); ok && gCtx != nil {
		ctx = gCtx.Request.Context()
	}
	if logger, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return logger
	}
	return DefaultLogger()
}

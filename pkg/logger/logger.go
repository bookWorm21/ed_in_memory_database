package logger

import (
	"io"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel -
type LogLevel string

const (
	// ErrorLevel -
	ErrorLevel LogLevel = "error"
	// WarnLevel -
	WarnLevel LogLevel = "warn"
	// InfoLevel -
	InfoLevel LogLevel = "info"
	// DebugLevel -
	DebugLevel LogLevel = "debug"
	// PanicLevel -
	PanicLevel LogLevel = "panic"
	// FatalLevel -
	FatalLevel LogLevel = "fatal"
)

// Params -
type Params struct {
	ServiceName string
	LogDir      string
	LogLevel    LogLevel
}

// logger - обертка над zap.Logger
type logger struct {
	logger *zap.SugaredLogger
}

// New - конструирует реализацию Logger
func New(params Params, out io.Writer) (Logger, error) {
	topicErrors := zapcore.AddSync(out)
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= toZapLevel(params.LogLevel)
	})

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, topicErrors, highPriority),
	)

	zapLogger := zap.New(core)

	return &logger{logger: zapLogger.Sugar()}, nil
}

func (l logger) Debugf(message string, args ...any) {
	l.logger.Debugf(message, args)
}

func (l logger) Infof(message string, args ...any) {
	l.logger.Infof(message, args)
}

func (l logger) Warnf(message string, args ...any) {
	l.logger.Warnf(message, args)
}

func (l logger) Errorf(message string, args ...any) {
	l.logger.Errorf(message, args)
}

func (l logger) Fatalf(message string, args ...any) {
	l.logger.Fatalf(message, args)
}

func (l logger) Withf(key Field, value any) Logger {
	return &logger{logger: l.logger.With(key, value)}
}

func (l logger) Sync() error {
	return l.logger.Sync()
}

func toZapLevel(level LogLevel) zapcore.Level {
	switch LogLevel(strings.ToLower(string(level))) {
	case ErrorLevel:
		return zap.ErrorLevel
	case WarnLevel:
		return zap.WarnLevel
	case InfoLevel:
		return zap.InfoLevel
	case DebugLevel:
		return zap.DebugLevel
	case PanicLevel:
		return zap.PanicLevel
	case FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

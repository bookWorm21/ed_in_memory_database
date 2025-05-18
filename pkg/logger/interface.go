package logger

type Logger interface {
	Debugf(message string, args ...any)
	Infof(message string, args ...any)
	Warnf(message string, args ...any)
	Errorf(message string, args ...any)
	Fatalf(message string, args ...any)
	Withf(key Field, value any) Logger
	Sync() error
}

package logger

// Logger -
type Logger interface {
	// Debugf -
	Debugf(message string, args ...any)
	// Infof -
	Infof(message string, args ...any)
	// Warnf -
	Warnf(message string, args ...any)
	// Errorf -
	Errorf(message string, args ...any)
	// Fatalf -
	Fatalf(message string, args ...any)
	// Withf -
	Withf(key Field, value any) Logger
	// Sync -
	Sync() error
}

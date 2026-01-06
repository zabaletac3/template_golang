package logger

var defaultLogger Logger

func SetDefault(l Logger) {
	defaultLogger = l
}

func Default() Logger {
	return defaultLogger
}

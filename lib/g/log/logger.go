package log

const (
	LOG_TRACE LogLevel = 1
	LOG_DEBUG          = 2
	LOG_INFO           = 3
	LOG_WARN           = 4
	LOG_ERROR          = 5

	LOG_ALL LogLevel = 0
	LOG_NONE         = 9999
	LOG_DEFAULT      = LOG_INFO

	TAG_TRACE = "TRACE"
	TAG_DEBUG = "DEBUG"
	TAG_INFO  = "INFO"
	TAG_WARN  = "WARN"
	TAG_ERROR = "ERROR"
)

type LogLevel uint32

type Logger interface {
	Trace(v ...interface{})
	Tracef(format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

var logger Logger

func init() {
	logger = NewFileBasedLogger().SetLogLevel(LOG_DEFAULT).SetCallerDepth(1)
}

func Trace(v ...interface{}) {
	logger.Trace(v...)
}

func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

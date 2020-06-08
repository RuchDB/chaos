package log

/************************* Logger Interface *************************/

const (
	LOG_TRACE LogLevel = 1
	LOG_DEBUG          = 2
	LOG_INFO           = 3
	LOG_WARN           = 4
	LOG_ERROR          = 5

	LOG_ALL     LogLevel = 0
	LOG_NONE             = 9999
	LOG_DEFAULT          = LOG_INFO

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

package log

import (
	"fmt"
	"strings"

	"github.com/RuchDB/chaos/lib/g/util"
)

/************************* File-Based Logger *************************/

type fileLogger struct {
	logLevel LogLevel

	callerDepth int
}

func NewFileBasedLogger() *fileLogger {
	return &fileLogger{
		logLevel: LOG_DEFAULT,

		callerDepth: 0,
	}
}

func (flog *fileLogger) SetLogLevel(level LogLevel) {
	flog.logLevel = level
}

func (flog *fileLogger) SetCallerDepth(depth int) {
	flog.callerDepth = depth
}

func (flog *fileLogger) write(msg string) {
	//! TODO: ...
	fmt.Print(msg)
}

func (flog *fileLogger) log(level LogLevel, tag string, v ...interface{}) {
	if level >= flog.logLevel {
		msg := fmt.Sprintf("%s %s\n", flog.prefix(tag), fmt.Sprint(v...))
		flog.write(msg)
	}
}

func (flog *fileLogger) logf(level LogLevel, tag string, format string, v ...interface{}) {
	if level >= flog.logLevel {
		body := fmt.Sprintf(format, v...)

		var msg string
		if len(body) > 0 && body[len(body)-1] == '\n' {
			msg = fmt.Sprintf("%s %s", flog.prefix(tag), body)
		} else {
			msg = fmt.Sprintf("%s %s\n", flog.prefix(tag), body)
		}

		flog.write(msg)
	}
}

func (flog *fileLogger) prefix(tag string) string {
	// 2020-01-01 00:00:00 [G1] [INFO] [main.go:10 main]
	var builder strings.Builder

	time := util.FormatTime(util.Now())
	builder.WriteString(time)

	if goid, ok := util.Goid(); ok {
		builder.WriteString(fmt.Sprintf(" [G%d]", goid))
	}

	builder.WriteString(fmt.Sprintf(" [%s]", tag))

	// !!! NOTE: Take care NOT to refactor (inline/extract) this function as well as its caller in `fileLogger`.
	//           OTHERWISE the CALLER FRAME may be NOT what we need.
	if fn, file, line, ok := util.Caller(3 + flog.callerDepth); ok {
		builder.WriteString(fmt.Sprintf(" [%s:%d %s]", file, line, fn))
	}

	return builder.String()
}

func (flog *fileLogger) Trace(v ...interface{}) {
	flog.log(LOG_TRACE, TAG_TRACE, v...)
}

func (flog *fileLogger) Tracef(format string, v ...interface{}) {
	flog.logf(LOG_TRACE, TAG_TRACE, format, v...)
}

func (flog *fileLogger) Debug(v ...interface{}) {
	flog.log(LOG_DEBUG, TAG_DEBUG, v...)
}

func (flog *fileLogger) Debugf(format string, v ...interface{}) {
	flog.logf(LOG_DEBUG, TAG_DEBUG, format, v...)
}

func (flog *fileLogger) Info(v ...interface{}) {
	flog.log(LOG_INFO, TAG_INFO, v...)
}

func (flog *fileLogger) Infof(format string, v ...interface{}) {
	flog.logf(LOG_INFO, TAG_INFO, format, v...)
}

func (flog *fileLogger) Warn(v ...interface{}) {
	flog.log(LOG_WARN, TAG_WARN, v...)
}

func (flog *fileLogger) Warnf(format string, v ...interface{}) {
	flog.logf(LOG_WARN, TAG_WARN, format, v...)
}

func (flog *fileLogger) Error(v ...interface{}) {
	flog.log(LOG_ERROR, TAG_ERROR, v...)
}

func (flog *fileLogger) Errorf(format string, v ...interface{}) {
	flog.logf(LOG_ERROR, TAG_ERROR, format, v...)
}

/************************* File Logger Builder *************************/

type fileLoggerBuilder struct {
	logLevel LogLevel

	callerDepth int
}

func NewFileBasedLoggerBuilder() *fileLoggerBuilder {
	return &fileLoggerBuilder{
		logLevel: LOG_DEFAULT,

		callerDepth: 0,
	}
}

func (builder *fileLoggerBuilder) SetLogLevel(level LogLevel) *fileLoggerBuilder {
	builder.logLevel = level
	return builder
}

func (builder *fileLoggerBuilder) SetCallerDepth(depth int) *fileLoggerBuilder {
	builder.callerDepth = depth
	return builder
}

func (builder *fileLoggerBuilder) Build() Logger {
	return &fileLogger{
		logLevel: builder.logLevel,

		callerDepth: builder.callerDepth,
	}
}

package log

import (
	"fmt"
	"strings"

	"github.com/RuchDB/chaos/util"
)

type fileLog struct {
	logLevel LogLevel

	callerDepth int
}

func NewFileBasedLogger() *fileLog {
	return &fileLog{
		logLevel: LOG_DEFAULT,

		callerDepth: 0,
	}
}

func (flog *fileLog) SetLogLevel(level LogLevel) *fileLog {
	flog.logLevel = level
	return flog
}

func (flog *fileLog) SetCallerDepth(depth int) *fileLog {
	flog.callerDepth = depth
	return flog
}

func (flog *fileLog) write(msg string) {
	fmt.Print(msg)
}

func (flog *fileLog) log(level LogLevel, tag string, v ...interface{}) {
	if level >= flog.logLevel {
		msg := fmt.Sprintf("%s %s\n", flog.prefix(tag), fmt.Sprint(v...))
		flog.write(msg)
	}
}

func (flog *fileLog) logf(level LogLevel, tag string, format string, v ...interface{}) {
	if level >= flog.logLevel {
		body := fmt.Sprintf(format, v...)

		var msg string
		if len(body) > 0 && body[len(body) - 1] == '\n' {
			msg = fmt.Sprintf("%s %s", flog.prefix(tag), body)
		} else {
			msg = fmt.Sprintf("%s %s\n", flog.prefix(tag), body)
		}

		flog.write(msg)
	}
}

func (flog *fileLog) prefix(tag string) string {
	// 2020-01-01 00:00:00 [INFO] [main.go:10 main] [G1]
	var builder strings.Builder

	time := util.FormatTime(util.Now())
	builder.WriteString(time)

	builder.WriteString(fmt.Sprintf(" [%s]", tag))

	// !!! NOTE: Take care NOT to refactor (inline/extract) this function as well as its caller in `fileLog`.
	//           OR the CALLER FRAME may be NOT what we need.
	if fn, file, line, ok := util.Caller(3 + flog.callerDepth); ok {
		builder.WriteString(fmt.Sprintf(" [%s:%d %s]", file, line, fn))
	}

	if goid, ok := util.Goid(); ok {
		builder.WriteString(fmt.Sprintf(" [G%d]", goid))
	}

	return builder.String()
}

func (flog *fileLog) Trace(v ...interface{}) {
	flog.log(LOG_TRACE, TAG_TRACE, v...)
}

func (flog *fileLog) Tracef(format string, v ...interface{}) {
	flog.logf(LOG_TRACE, TAG_TRACE, format, v...)
}

func (flog *fileLog) Debug(v ...interface{}) {
	flog.log(LOG_DEBUG, TAG_DEBUG, v...)
}

func (flog *fileLog) Debugf(format string, v ...interface{}) {
	flog.logf(LOG_DEBUG, TAG_DEBUG, format, v...)
}

func (flog *fileLog) Info(v ...interface{}) {
	flog.log(LOG_INFO, TAG_INFO, v...)
}

func (flog *fileLog) Infof(format string, v ...interface{}) {
	flog.logf(LOG_INFO, TAG_INFO, format, v...)
}

func (flog *fileLog) Warn(v ...interface{}) {
	flog.log(LOG_WARN, TAG_WARN, v...)
}

func (flog *fileLog) Warnf(format string, v ...interface{}) {
	flog.logf(LOG_WARN, TAG_WARN, format, v...)
}

func (flog *fileLog) Error(v ...interface{}) {
	flog.log(LOG_ERROR, TAG_ERROR, v...)
}

func (flog *fileLog) Errorf(format string, v ...interface{}) {
	flog.logf(LOG_ERROR, TAG_ERROR, format, v...)
}

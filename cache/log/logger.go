package log

import (
	"github.com/RuchDB/chaos/lib/g/log"
	"github.com/RuchDB/chaos/lib/g/net"
)

var Logger log.Logger

func init() {
	Logger = log.NewFileBasedLoggerBuilder().
		SetLogLevel(log.LOG_INFO).
		Build()

	net.RegisterLogger(Logger)
}

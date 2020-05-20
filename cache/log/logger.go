package log

import (
	"github.com/RuchDB/chaos/lib/g/log"
	"github.com/RuchDB/chaos/lib/g/net"
)

var logger log.Logger

func init() {
	logger = log.NewFileBasedLoggerBuilder().
		SetLogLevel(log.LOG_INFO).
		Build()
	
	net.RegisterLogger(logger)
}

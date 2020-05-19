package net

import (
	"github.com/RuchDB/chaos/lib/g/log"
)

var logger log.Logger = log.NewDummyLoggerBuilder().Build()

func RegisterLogger(l log.Logger) {
	logger = l
}

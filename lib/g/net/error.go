package net

import (
	"net"
)

func IsNetError(err error) bool {
	_, ok := err.(net.Error)
	return ok
}

func IsNetOpError(err error) bool {
	_, ok := err.(*net.OpError)
	return ok
}

func IsTimeoutError(err error) bool {
	opErr, ok := err.(*net.OpError)
	return ok && opErr.Timeout()
}

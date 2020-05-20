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
	netErr, ok := err.(net.Error)
	return ok && netErr.Timeout()
}

func IsTemporaryError(err error) bool {
	netErr, ok := err.(net.Error)
	return ok && netErr.Temporary()
}

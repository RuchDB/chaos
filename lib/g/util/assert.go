package util

import (
	"fmt"
)

func Assert(condition bool, args ...interface{}) {
	if !condition {
		panic(fmt.Sprint(args...))
	}
}

func Assertf(condition bool, format string, args ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(format, args...))
	}
}

func AssertEqual(expected, actual interface{}, args ...interface{}) {
	Assert(expected == actual, args...)
}

func AssertEqualf(expected, actual interface{}, format string, args ...interface{}) {
	Assertf(expected == actual, format, args...)
}

package util

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

// [func, file, line] of caller
func Caller(depth int) (string, string, int, bool) {
	pcs := make([]uintptr, 1)
	if n := runtime.Callers(2+depth, pcs); n != 1 {
		return "", "", 0, false
	}

	// Caller frame
	frame, _ := runtime.CallersFrames(pcs).Next()

	// Trim function name: github.com/RuchDB/chaos/main.main
	fn := frame.Function
	if dot := strings.LastIndexByte(fn, '.'); dot >= 0 {
		fn = fn[dot+1:]
	}

	// Trim file name: /home/user/workspace/RuchDB/chaos/main.go
	file := frame.File
	if slash := strings.LastIndexByte(file, '/'); slash >= 0 {
		file = file[slash+1:]
	}

	return fn, file, frame.Line, true
}

// Goroutine ID
func Goid() (int64, bool) {
	// Stack: goroutine 1 [running]: ...
	buf := make([]byte, 20)
	if n := runtime.Stack(buf, false); n <= 10 {
		return 0, false
	}

	space := bytes.IndexByte(buf[10:], ' ')
	if space <= 0 {
		return 0, false
	}
	buf = buf[10 : 10+space]

	gid, err := strconv.ParseInt(string(buf), 10, 64)
	if err != nil {
		return 0, false
	}

	return gid, true
}

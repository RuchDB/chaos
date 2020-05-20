package util

import (
	"syscall"
)

func GetHardOpenFileLimit() (int, error) {
	var rlimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		return 0, err
	}

	return int(rlimit.Max), nil
}

func GetSoftOpenFileLimit() (int, error) {
	var rlimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		return 0, err
	}

	return int(rlimit.Cur), nil
}

func SetOpenFileLimit(limit int) (int, error) {
	var rlimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		return 0, err
	}

	if limit <= 0 {
		return int(rlimit.Cur), nil
	}

	rlimit.Cur = MinInt(limit, int(rlimit.Max))
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		return 0, err
	}

	return int(rlimit.Cur), nil
}

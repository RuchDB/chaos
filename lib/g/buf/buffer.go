package buf

import (
	"errors"
)

const (
	BUF_SIZE_MIN     = 8
	BUF_SIZE_MAX     = 1024 * 1024 * 64 // 64MB
	BUF_SIZE_DEFAULT = 128
)

var (
	ERR_BUF_NO_DATA    = errors.New("no data")
	ERR_BUF_NOT_ENOUGH = errors.New("no enough data")
	ERR_BUF_LARGE_SIZE = errors.New("too large size")
)

type Buffer interface {
	Len() int
	Cap() int

	Clear()
	Reset()

	Peak(n int) ([]byte, error)
	PeakExact(n int) ([]byte, error)
	PeakAll() ([]byte, error)

	Read(n int) ([]byte, error)
	ReadExact(n int) ([]byte, error)
	ReadAll() ([]byte, error)

	Write(bs []byte) error
}

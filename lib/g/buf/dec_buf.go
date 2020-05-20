package buf

import (
	"github.com/RuchDB/chaos/lib/g/util"
)

const (
	DECODE_BUFFER_STABLE_SCALE = 4
)

type decodeBuffer struct {
	buf  []byte
	head int
	tail int
	size int

	minSize int
	maxSize int
}

// Both minSize & maxSize SHOULD be power of 2
func NewDecodeBuffer(minSize, maxSize int) *decodeBuffer {
	util.Assert(minSize <= maxSize, "minSize SHOULD NOT be larger than maxSize")

	min := util.MinInt(minSize, BUF_SIZE_MIN)
	max := util.MaxInt(maxSize, BUF_SIZE_MAX)

	return &decodeBuffer{
		buf:  make([]byte, min),
		head: 0,
		tail: 0,
		size: min,

		minSize: min,
		maxSize: max,
	}
}

func (buf *decodeBuffer) Len() int {
	return buf.tail - buf.head
}

func (buf *decodeBuffer) Cap() int {
	return buf.size
}

func (buf *decodeBuffer) Clear() {
	buf.head = 0
	buf.tail = 0
}

func (buf *decodeBuffer) Reset() {
	if buf.size > buf.minSize {
		buf.buf = make([]byte, buf.minSize)
		buf.size = buf.minSize
	}
	buf.head = 0
	buf.tail = 0
}

func (buf *decodeBuffer) Peak(n int) ([]byte, error) {
	if buf.Len() == 0 {
		return nil, ERR_BUF_NO_DATA
	}

	n = util.MinInt(n, buf.Len())
	return buf.buf[buf.head : buf.head+n], nil
}

func (buf *decodeBuffer) PeakExact(n int) ([]byte, error) {
	if buf.Len() == 0 {
		return nil, ERR_BUF_NO_DATA
	}
	if buf.Len() < n {
		return nil, ERR_BUF_NOT_ENOUGH
	}

	return buf.buf[buf.head : buf.head+n], nil
}

func (buf *decodeBuffer) PeakAll() ([]byte, error) {
	if buf.Len() == 0 {
		return nil, ERR_BUF_NO_DATA
	}

	return buf.buf[buf.head:buf.tail], nil
}

func (buf *decodeBuffer) Read(n int) ([]byte, error) {
	if buf.Len() == 0 {
		return nil, ERR_BUF_NO_DATA
	}

	n = util.MinInt(n, buf.Len())
	bs := buf.buf[buf.head : buf.head+n]
	buf.head = buf.head + n

	buf.shrink()

	return bs, nil
}

func (buf *decodeBuffer) ReadExact(n int) ([]byte, error) {
	if buf.Len() == 0 {
		return nil, ERR_BUF_NO_DATA
	}
	if buf.Len() < n {
		return nil, ERR_BUF_NOT_ENOUGH
	}

	bs := buf.buf[buf.head : buf.head+n]
	buf.head = buf.head + n

	buf.shrink()

	return bs, nil
}

func (buf *decodeBuffer) ReadAll() ([]byte, error) {
	if buf.Len() == 0 {
		return nil, ERR_BUF_NO_DATA
	}

	bs := buf.buf[buf.head:buf.tail]
	buf.head = 0
	buf.tail = 0

	buf.shrink()

	return bs, nil
}

func (buf *decodeBuffer) shrink() {
	// If buffer size is too large, shrink it
	if buf.size > buf.minSize*DECODE_BUFFER_STABLE_SCALE {
		// New buffer size
		newSize := buf.minSize
		blen := buf.Len()
		for newSize < blen {
			newSize *= 2
		}

		// Create a smaller buffer
		if newSize < buf.size {
			newBuf := make([]byte, newSize)
			copy(newBuf, buf.buf[buf.head:buf.tail])

			buf.buf = newBuf
			buf.head = 0
			buf.tail = blen
			buf.size = newSize
		}
	}
}

func (buf *decodeBuffer) Write(bs []byte) error {
	// Remain enough buffer
	if err := buf.remain(len(bs)); err != nil {
		return err
	}

	// Move data to the beginning
	if buf.size-buf.tail < len(bs) {
		blen := buf.Len()
		copy(buf.buf, buf.buf[buf.head:buf.tail])

		buf.head = 0
		buf.tail = blen
	}

	copy(buf.buf[buf.tail:], bs)
	buf.tail += len(bs)

	return nil
}

func (buf *decodeBuffer) remain(size int) error {
	size = buf.Len() + size
	if size > buf.maxSize {
		return ERR_BUF_LARGE_SIZE
	}

	// Expand/Shrink buffer
	newSize := buf.size
	for newSize < size {
		newSize *= 2
	}
	for newSize/2 >= size {
		newSize /= 2
	}

	// Create buffer
	if newSize != size {
		newBuf := make([]byte, newSize)
		copy(newBuf, buf.buf[buf.head:buf.tail])

		blen := buf.Len()
		buf.buf = newBuf
		buf.head = 0
		buf.tail = blen
		buf.size = newSize
	}

	return nil
}

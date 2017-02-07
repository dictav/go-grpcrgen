package myrouter

import (
	"errors"
	"fmt"
	"io"

	"github.com/google/flatbuffers/go"
)

var (
	// DefaultMaxChannels is number of buffer chennels. It is used to receive HTTP Reuest's body. Default value is 100.
	DefaultMaxChannels = 100

	// DefaultMaxBufferSize is limit of buffer size. It is used to receive HTTP Reuest's body. Default value is 4KB.
	DefaultMaxBufferSize = 4 * 1024
)

var bufCh chan []byte

type tabler interface {
	Table() flatbuffers.Table
}

// InitBuffer creates buffers. This function is called by readBody if no buffers are created. If you care about the performance at runtime, call this method when initializing the application.
func InitBuffer() {
	bufCh = make(chan []byte, DefaultMaxBufferSize)
	for i := 0; i < DefaultMaxBufferSize; i++ {
		bufCh <- make([]byte, DefaultMaxBufferSize)
	}
}

func newBuilderForFlatbuffersRawCodec(buf []byte) *flatbuffers.Builder {
	return &flatbuffers.Builder{Bytes: buf}
}

func readBody(r io.Reader) ([]byte, error) {
	if cap(bufCh) == 0 {
		InitBuffer()
	}

	buf := <-bufCh
	defer func() {
		bufCh <- buf
	}()

	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}

	if n < 5 {
		str := fmt.Sprintf("Invalid request body: %v", buf)
		return nil, errors.New(str)
	}

	return buf[:n], nil
}

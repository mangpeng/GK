package TCP

import (
	"errors"
	"math"
)

// receiveBuffer
// ChunkBuffer for writing bytes received from socket.
type receiveBuffer struct {
	chunk  []byte
	offset int
}

const (
	receiveChunkBufferSize = math.MaxUint16
)

// initReceiveBuffer
// Initialize receiveBuffer as max size.
func (b *receiveBuffer) initReceiveBuffer() {
	b.chunk = make([]byte, receiveChunkBufferSize)
}

// write
// Copy new datum to receiveBuffer.
// if not enough to space to copy, append byte slice.
// todo if don't read bytes while write bytes, receiveBuffer is getting big forever? how can i process this?
func (b *receiveBuffer) write(buf []byte) {
	l := len(buf)
	if n := copy(b.chunk[b.offset:], buf); n < l {
		b.chunk = append(b.chunk, buf[n:]...)
	}

	b.offset = b.offset + len(buf)
}

// peek
// Return byte slice requested without reading receiveBuffer.
func (b *receiveBuffer) peek(size int) ([]byte, error) {
	if size > b.offset {
		return nil, errors.New("buffer size is not enough to read the size requested")
	}

	return b.chunk[:size], nil
}

// read
// Read the data by the requested size and adjust receiveBuffer offset by the requested size.
func (b *receiveBuffer) read(buf []byte, size int) error {
	if size > b.offset {
		return errors.New("buffer size is not enough to read the size requested")
	}

	if len(buf) < size {
		return errors.New("the requested size is bigger than the size of receive buffer")
	}

	// todo when read is requested, buffer data have to be copied everytime.. is it right?
	b.offset = b.offset - size
	copy(buf, b.chunk[:size])
	copy(b.chunk, b.chunk[size:])

	return nil
}

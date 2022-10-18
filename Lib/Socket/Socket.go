package Socket

import (
	"errors"
	"net"
)

type socketBuffer struct {
	data   []byte
	offset int
}

type TCP struct {
	connection net.Conn
	connected  bool
	buffer     socketBuffer
}

type Listener struct {
	ln       net.Listener
	flagStop bool
}

func (b *socketBuffer) initSocketBuffer() {
	b.data = make([]byte, 65536)
}

func (b *socketBuffer) write(p []byte) {
	l := len(p)
	if n := copy(b.data[b.offset:], p); n < l {
		b.data = append(b.data, p[n:]...)
	}

	b.offset = b.offset + len(p)
}

func (b *socketBuffer) peek(size int) ([]byte, error) {
	if size > b.offset {
		return nil, errors.New("overflow")
	}

	return b.data[:size], nil
}

// read buffer is copied from receive buffer to buffer you request
func (b *socketBuffer) read(buffer []byte, size int) error {
	if size > b.offset {
		return errors.New("overflow")
	}

	if len(buffer) < size {
		panic("buffer size is not enough to copy")
	}

	// TODO : 데이터를 읽을 때 마다 버퍼를 전체 복사를 한다... 개선필요.
	// [읽은 데이터][남은 공간 .........] => [남은 공간........         ] 이와 같이 당기기 위해서
	b.offset = b.offset - size
	copy(buffer, b.data[:size])
	copy(b.data, b.data[size:])

	return nil
}

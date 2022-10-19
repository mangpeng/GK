package TCP

import (
	"fmt"
	"math"
	"net"
)

// Session
// It contains a TCP connection objects and i/o data.
type Session struct {
	connection net.Conn
	connected  bool
	buffer     receiveBuffer
}

const (
	readBufferSize = math.MaxUint16
)

// TryConnect
// Try to tcp-connection.
func (s *Session) TryConnect(addr string, port uint) (bool, error) {
	var err error
	host := fmt.Sprint(addr, ":", port)
	s.connection, err = net.Dial("tcp", host)

	if err != nil {
		return false, err
	}

	s.connected = true
	s.buffer.initReceiveBuffer()
	return true, nil
}

// TryDisconnect
// Try to tcp-disconnection.
func (s *Session) TryDisconnect() (bool, error) {
	err := s.connection.Close()
	if err != nil {
		return false, err
	}

	s.connected = false
	return true, nil
}

// IsConnected
// Return the state of connection : true or false.
func (s *Session) IsConnected() bool {
	return s.connected
}

//// DelayClose 곧 끊김 ㅋ
//func (t *TCP) DelayClose() {
//	closeTimer := time.NewTimer(time.Second * 2)
//	go func() {
//		<-closeTimer.C
//		closeTimer.Stop()
//		t.Close()
//	}()
//}

// IOReceiveHandler
// todo It need to check later.
func (s *Session) IOReceiveHandler(success func(), fail func(error)) {
	readBuffer := make([]byte, readBufferSize)
	for {
		l, err := s.connection.Read(readBuffer)
		if err != nil {
			if l == 0 {
				s.connected = false
				fail(err)
			}
			break
		}

		if l > 0 {
			s.buffer.write(readBuffer[:l])
			success()
		}
	}
}

func (s *Session) Send(buf []byte) (int, error) {
	return s.connection.Write(buf)
}

func (s *Session) Peek(size int) ([]byte, error) {
	return s.buffer.peek(size)
}

func (s *Session) Read(buf []byte, size int) error {
	return s.buffer.read(buf, size)
}

// GetLocalAddr
// Return host local address.
func (s *Session) GetLocalAddr() string {
	return s.connection.LocalAddr().String()
}

// GetRemoteAddr
// Return quest(peer) address.
func (s *Session) GetRemoteAddr() string {
	return s.connection.RemoteAddr().String()
}

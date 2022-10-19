package TCP

import (
	"fmt"
	"math"
	"net"
)

// TCPSession
// It contains a TCP connection objects and i/o data.
type TCPSession struct {
	connection net.Conn
	connected  bool
	buffer     receiveBuffer
}

const (
	readBufferSize = math.MaxUint16
)

// TryConnect
// Try to tcp-connection.
func (ts *TCPSession) TryConnect(address string, port uint) (bool, error) {
	var err error
	host := fmt.Sprint(address, ":", port)
	ts.connection, err = net.Dial("tcp", host)

	if err != nil {
		return false, err
	}

	ts.connected = true
	ts.buffer.initReceiveBuffer()
	return true, nil
}

// TryDisconnect
// Try to tcp-disconnection.
func (ts *TCPSession) TryDisconnect() (bool, error) {
	err := ts.connection.Close()
	if err != nil {
		return false, err
	}

	ts.connected = false
	return true, nil
}

// IsConnected
// Return the state of connection : true or false.
func (ts *TCPSession) IsConnected() bool {
	return ts.connected
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

// ConnectionHandler
// todo It need to check later.
func (ts *TCPSession) OnReceivedIOHandler(onSuccess func(), onFail func(error)) {
	readBuffer := make([]byte, readBufferSize)
	for {
		numOfBytes, err := ts.connection.Read(readBuffer)
		if err != nil {
			if numOfBytes == 0 {
				ts.connected = false
				onFail(err)
			}
			break
		}

		if numOfBytes > 0 {
			ts.buffer.write(readBuffer[:numOfBytes])
			onSuccess()
		}
	}
}

func (ts *TCPSession) Send(data []byte) (int, error) {
	return ts.connection.Write(data)
}

func (ts *TCPSession) Peek(size int) ([]byte, error) {
	return ts.buffer.peek(size)
}

func (ts *TCPSession) Read(buf []byte, size int) error {
	return ts.buffer.read(buf, size)
}

// GetLocalAddr
// Return host local address.
func (ts *TCPSession) GetLocalAddr() net.Addr {
	return ts.connection.LocalAddr()
}

// GetRemoteAddr
// Return quest(peer) address.
func (ts *TCPSession) GetRemoteAddr() net.Addr {
	return ts.connection.RemoteAddr()
}

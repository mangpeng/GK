package TCP

import (
	"fmt"
	"net"
)

// Listener
// It is for accepting from clients.
type Listener struct {
	listener net.Listener
	stop     bool
}

// Listen is for server
func (l *Listener) Listen(port uint) error {
	str := fmt.Sprintf("0.0.0.0:%d", port)
	ln, err := net.Listen("tcp", str)
	if err != nil {
		return err
	}

	l.listener = ln
	l.stop = false

	return nil
}

// AsyncAccept
// Accept clients by go routine.
func (l *Listener) AsyncAccept(call func(*Session)) {
	go func() {
		for {
			c, _ := l.listener.Accept()
			if l.stop {
				break
			}
			s := new(Session)
			s.connection = c
			s.connected = true
			s.buffer.initReceiveBuffer()

			call(s)
		}
	}()
}

func (l *Listener) IsStopped() bool {
	return l.stop
}

func (l *Listener) TryStopAccept() (bool, error) {
	if !l.stop {
		l.stop = true
		err := l.listener.Close()
		return true, err
	}

	return false, nil
}

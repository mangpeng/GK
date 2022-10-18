package Packet

import (
	Packet "Client/packet"
	"encoding/binary"
	"fmt"
	"nm.go/go-server/library/socket"
	"sync"
	"sync/atomic"
)

type Session struct {
	Host       string
	Connection *socket.TCP
	AccountId  string
	Wg         WaitGroupCount
	Result     interface{}
}

type WaitGroupCount struct {
	sync.WaitGroup
	count int64
}

func (wg *WaitGroupCount) Add(delta int) {
	atomic.AddInt64(&wg.count, int64(delta))
	wg.WaitGroup.Add(delta)
}

func (wg *WaitGroupCount) Done() {
	atomic.AddInt64(&wg.count, -1)
	wg.WaitGroup.Done()
}

func (wg *WaitGroupCount) GetCount() int {
	return int(atomic.LoadInt64(&wg.count))
}

const headerSize int = 4

var (
	session       *Session
	receiveBuffer []byte

	isInit           bool = false
	packetHandlerMap map[uint16]func(*Session, []byte)
)

func Init() {
	isInit = true
	receiveBuffer = make([]byte, 65536)
	packetHandlerMap = make(map[uint16]func(*Session, []byte))

	InitPacketHandler()
}

func InitPacketHandler() {
	// 패킷이 추가 될때 마다 패킷 핸들러를 추가 해야함. => 자동화 코드로 변경 필요.
	packetHandlerMap[uint16(Packet.ENUMSTC_MESSAGE)] = OnReceiveMessagePacket

	packetHandlerMap[uint16(Packet.ENUMSTC_LOGIN)] = OnReceiveLoginPacket
	packetHandlerMap[uint16(Packet.ENUMSTC_SIGNUP)] = OnReceiveSignupPacket

	packetHandlerMap[uint16(Packet.ENUMSTC_ALARM_LIST)] = OnReceiveAlarmListPacket
	packetHandlerMap[uint16(Packet.ENUMSTC_ALARM_ADD)] = OnReceiveAlarmAddPacket
	packetHandlerMap[uint16(Packet.ENUMSTC_ALARM_DELETE)] = OnReceiveAlarmDeletePacket
	packetHandlerMap[uint16(Packet.ENUMSTC_ALARM_MODIFY)] = OnReceiveAlarmModifyPacket
}

func GetSession() *Session {
	if session == nil {

		session = &Session{Connection: new(socket.TCP)}
	}
	return session
}

func (s *Session) Connect(address string, port uint) bool {

	if !isInit {
		panic("you need to call Init() before calling this function")
	}

	return s.Connection.Connect(address, port)
}

func (s *Session) StartReceive() {
	s.Connection.ConnectionHandler(func() {
		s.processPacket()
	}, nil)
}

func (s *Session) SendSync(buf []byte) {

	s.Wg.Add(1)
	go s.Connection.Send(buf)
	s.Wg.Wait()
}

func (s *Session) SendAsync(buf []byte) {
	s.Connection.Send(buf)
}

func (s *Session) processPacket() {
	for {
		id, packet := s.extractPacket()

		if packet == nil {
			break
		}

		s.parsePacket(id, packet)
	}
}

func (s *Session) extractPacket() (uint16, []byte) {
	// [id(2)][size(2)[ ............   ]
	// header(4) : id(2) + size(2)
	// size : id(2) + size(2) + data(....)

	rawHeader, err := s.Connection.Peek(headerSize)
	if err != nil {
		return 0, nil
	}
	id := binary.LittleEndian.Uint16(rawHeader)

	rawSize := rawHeader[2:4]
	size := binary.LittleEndian.Uint16(rawSize)

	err = s.Connection.Read(receiveBuffer, int(size))
	if err != nil {
		return 0, nil
	}

	return id, receiveBuffer[headerSize:size]
}

func (s *Session) parsePacket(id uint16, packet []byte) {
	val, exists := packetHandlerMap[id]
	if exists {
		val(s, packet)
	} else {
		fmt.Println("Can't process this packet id\n", id)
	}

	if s.Wg.GetCount() != 0 {
		s.Wg.Done()
	}
}

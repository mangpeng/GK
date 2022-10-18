package Packet

import (
	Packet "Client/packet"
	"Client/src/Util"
	"encoding/binary"
	"fmt"
	flatbuffers "github.com/google/flatbuffers/go"
)

type FnPacketGen func(*flatbuffers.Builder) (Packet.ENUM, flatbuffers.UOffsetT)

func MakePacket(generator FnPacketGen) []byte {
	builder := flatbuffers.NewBuilder(0)

	id, p := generator(builder)

	builder.Finish(p)
	buf := builder.FinishedBytes()

	const HeaderSize int = 4
	size := len(buf) + HeaderSize
	header := make([]byte, HeaderSize)
	binary.LittleEndian.PutUint16(header[:2], uint16(id))
	binary.LittleEndian.PutUint16(header[2:4], uint16(size))
	packet := append(header, buf...)

	return packet
}

func OnReceiveMessagePacket(session *Session, buf []byte) {
	fmt.Printf("OnReceiveMessagePacket\n")
	fmt.Println(string(buf))
	//packet := Packet.GetRootAsSTC_TEST1(buf, 0)
	//Alert.ShowPopup("test", string(packet.Message()))
}

func OnReceiveLoginPacket(session *Session, buf []byte) {
	//fmt.Printf("OnReceiveLoginPacket\n")
	packet := Packet.GetRootAsSTC_LOGIN(buf, 0)
	if packet.Loginok() {
		Util.ConsolePrintln(Util.SUCCESS, "Success to login")
		session.Result = true
	} else {
		Util.ConsolePrintln(Util.FAIL, "Failed to login")
		session.Result = false
		session.AccountId = ""
	}
}

func OnReceiveSignupPacket(session *Session, buf []byte) {
	//fmt.Printf("OnReceiveSignupPacket\n")
	packet := Packet.GetRootAsSTC_SIGNUP(buf, 0)
	if packet.Signupok() {
		Util.ConsolePrintln(Util.SUCCESS, "Succeed to make account")
		session.Result = true
	} else {
		Util.ConsolePrintln(Util.FAIL, "Failed to make account")
		session.Result = false
	}
}

func OnReceiveAlarmListPacket(session *Session, buf []byte) {
	//fmt.Printf("OnReceiveAlarmListPacket\n")
	packet := Packet.GetRootAsSTC_ALARM_LIST(buf, 0)
	if packet.AlarmsLength() == 0 {
		Util.ConsolePrintln(Util.FAIL, "There is no single registered alarm")
	} else {
		fmt.Println("+-----------------------------------------------------------------------------+")
		fmt.Printf("| %3s  %-20s %-50s|\n", "IDX", "TITLE", "CONTENT")
		fmt.Println("+-----------------------------------------------------------------------------+")
		for i := 0; i < packet.AlarmsLength(); i++ {
			alarm := Packet.Alarm{}
			packet.Alarms(&alarm, i)
			fmt.Printf("| %3d  %-20s %-50s|\n", alarm.Id(), string(alarm.Title()), string(alarm.Contents()))
		}
		fmt.Println("+-----------------------------------------------------------------------------+")
	}
}

func OnReceiveAlarmAddPacket(session *Session, buf []byte) {
	//fmt.Printf("OnReceiveAlarmAddPacket\n")
	packet := Packet.GetRootAsSTC_ALARM_ADD(buf, 0)
	if packet.Addok() {
		Util.ConsolePrintln(Util.SUCCESS, "Succeed to add alarm")
		session.Result = true
	} else {
		Util.ConsolePrintln(Util.FAIL, "Failed to add alarm")
		session.Result = false
	}
}

func OnReceiveAlarmDeletePacket(session *Session, buf []byte) {
	//fmt.Printf("OnReceiveAlarmDeletePacket\n")
	packet := Packet.GetRootAsSTC_ALARM_DELETE(buf, 0)
	if packet.Deleteok() {
		Util.ConsolePrintln(Util.SUCCESS, "Succeed to delete alarm")
		session.Result = true
	} else {
		Util.ConsolePrintln(Util.FAIL, "Failed to delete alarm")
		session.Result = false
	}
}

func OnReceiveAlarmModifyPacket(session *Session, buf []byte) {
	//fmt.Printf("OnReceiveAlarmModifyPacket\n")
	packet := Packet.GetRootAsSTC_ALARM_MODIFY(buf, 0)
	if packet.Modifyok() {
		Util.ConsolePrintln(Util.SUCCESS, "Succeed to modify alarm")
		session.Result = true
	} else {
		Util.ConsolePrintln(Util.FAIL, "Failed to modify alarm")
		session.Result = false
	}
}

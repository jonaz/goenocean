package goenocean

import "fmt"

type TelegramRps struct {
	packet
	senderId      [4]byte
	destinationId [4]byte
	data          byte
	status        byte
}

func NewTelegramRps() *TelegramRps {
	header := &header{}
	return &TelegramRps{
		packet:        packet{syncByte: 0x55, header: header},
		senderId:      [4]byte{0x00, 0x00, 0x00, 0x00}, //this can default to 00 since usb300 will add its own senderid
		destinationId: [4]byte{0xff, 0xff, 0xff, 0xff},
		status:        0,
		data:          0}
}
func (p *TelegramRps) Process() {
	length := len(p.packet.data)

	p.status = p.packet.data[length-1]
	copy(p.senderId[:], p.packet.data[length-5:length-1])
	copy(p.destinationId[:], p.packet.optData[1:5])

	//1 byte data only for RPS
	p.data = p.packet.data[1]
}

func (p *TelegramRps) Data() []byte {
	return []byte{p.data}
}

func (p *TelegramRps) SetTelegramData(data byte) {
	p.data = data
}

func (p *TelegramRps) SenderId() [4]byte {
	return p.senderId
}
func (p *TelegramRps) SetSenderId(data [4]byte) {
	p.senderId = data
}
func (p *TelegramRps) SetStatus(data byte) {
	p.status = data
}
func (p *TelegramRps) Encode() []byte {

	// 1 byte data + 4 byte sender id + 1 byte status
	data := []byte{TelegramTypeRps}
	data = append(data, p.data)
	data = append(data, p.senderId[:]...)
	data = append(data, p.status)

	p.packet.data = data

	//SubTelNum + Destination Id + dBm + security Level
	optData := []byte{0x03}
	optData = append(optData, p.destinationId[:]...)
	optData = append(optData, 0xff)
	optData = append(optData, 0x00)
	p.packet.optData = optData

	p.packet.SetPacketType(0x01)

	return p.packet.Encode()
}

func (p *TelegramRps) Action() string {

	//THIS WORKS. but must be refactored since it is part of EEP F6-02-01 and not the general RPS RORG type!
	fmt.Printf("raw action: %b\n", p.data)
	n := p.data >> 5
	fmt.Printf("bit action: %b\n", n)
	fmt.Printf("EB: %d\n", hasBit(p.data, 4)) // this also works :)
	switch n {
	case 0:
		return "Button A-ON"
	case 1:
		return "Button A-OFF"
	case 2:
		return "Button B-ON"
	case 3:
		return "Button B-OFF"

	}
	return "INGET"
}

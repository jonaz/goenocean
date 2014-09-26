package goenocean

import "fmt"

const (
	EepRps = 0xf6
	Eep1bs = 0xd5
	Eep4bs = 0xa5
	Eepvld = 0xd2
)

type PacketEepRps struct {
	Packet
	senderId      []byte
	destinationId []byte
	data          byte
	status        byte
}

func NewPacketEepRps() *PacketEepRps {
	header := &header{}
	return &PacketEepRps{Packet: Packet{syncByte: 0x55, header: header}, senderId: nil, destinationId: nil, data: 0}
}
func (p *PacketEepRps) Process() {
	length := len(p.Packet.data)

	p.status = p.Packet.data[length-1]
	p.senderId = p.Packet.data[length-5 : length-1]
	p.destinationId = p.Packet.optData[1:5]

	p.data = p.Packet.data[1]
}

func (p *PacketEepRps) Data() []byte {
	return []byte{p.data}
}

func (p *PacketEepRps) SenderId() []byte {
	return p.senderId
}

func (p *PacketEepRps) Action() string {

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
func hasBit(n byte, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
func reverse(x byte) byte {
	x = (x&0x55)<<1 | (x&0xAA)>>1
	x = (x&0x33)<<2 | (x&0xCC)>>2
	x = (x&0x0F)<<4 | (x&0xF0)>>4
	return x
}

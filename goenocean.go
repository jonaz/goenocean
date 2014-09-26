package goenocean

import (
	"errors"
	"fmt"
)

func Decode(data []byte) (p PacketInterface, err error) {
	if data[0] != 0x55 {
		return nil, errors.New("goenocean.Decode: must start with 0x55")
	}

	p = getPacket(data[4], data[6])
	p.SetSyncByte(data[0])
	p.SetHeaderFromBytes(data[1:5])
	p.SetHeaderCrc(data[5])

	p.SetData(data[6 : 6+p.Header().dataLength])
	p.SetOptData(data[6+p.Header().dataLength : 6+int(p.Header().dataLength)+int(p.Header().optDataLength)])
	p.SetDataCrc(data[len(data)-1])

	err = p.ValidateCrc()

	p.Process()
	return
}

func getPacket(packetType, rorg byte) PacketInterface {
	if packetType == PacketTypeRadioErp1 {
		switch rorg {
		case EepRps:
			return NewPacketEepRps()
		}
	}

	//return default packet
	return NewPacket()
}

func PrintHex(a []byte) { // {{{
	fmt.Println(ToHex(a))
} // }}}

func ToHex(a []byte) string { // {{{
	b := fmt.Sprintf("% x", a)
	return b
} // }}}

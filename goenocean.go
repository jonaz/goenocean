package goenocean

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	//"github.com/tarm/goserial"
)

func Decode(data []byte) (p *packet, err error) {
	if data[0] != 0x55 {
		return nil, errors.New("goenocean.Decode: must start with 0x55")
	}

	p = NewPacket()
	p.syncByte = data[0]
	p.header.dataLength = binary.BigEndian.Uint16(data[1:3])
	p.header.optDataLength = data[3]
	p.header.packetType = data[4]
	//p.headerCrc = crc(data[1:5])
	p.headerCrc = data[5]

	p.data = data[6 : 6+p.header.dataLength]
	p.optData = data[6+p.header.dataLength : 6+int(p.header.dataLength)+int(p.header.optDataLength)]
	p.dataCrc = data[len(data)-1]

	//TODO create a pkg.ValidateCrc and run it here

	return
}

func Read(rd io.Reader) []byte {
	//TODO use this example receivePacket when reading from serial https://github.com/kleckse/enocean/blob/master/esp3.py

	buf := make([]byte, 1)
	rawPacket := make([]byte, 1)
	state := 0

	for {
		rd.Read(buf)
		if buf[0] == 0x55 {
			state = 0
		}

		switch state {
		case 0: //0x55
			state = 1
			rawPacket = append(rawPacket, buf...)
			buf = make([]byte, 4) //read the header

		case 1: //header
			state = 2
			rawPacket = append(rawPacket, buf...)
			buf = make([]byte, 1) //read the crc header

		case 2: //crc header
			state = 3
			rawPacket = append(rawPacket, buf...)
			buf = make([]byte, binary.BigEndian.Uint16(rawPacket[1:2])) //read the opt data

		case 3: //data
			state = 4
			rawPacket = append(rawPacket, buf...)
			buf = make([]byte, (rawPacket[3])) //read the opt data

		case 4: //optional data
			state = 5
			rawPacket = append(rawPacket, buf...)
			buf = make([]byte, 1) //read the data crc

		case 5: //data crc
			state = 0
			rawPacket = append(rawPacket, buf...)
			buf = make([]byte, 1)
			break
			//default:
			//return nil
		}

	}

	return rawPacket
}

func printHex(a []byte) { // {{{
	fmt.Println(toHex(a))
} // }}}

func toHex(a []byte) string { // {{{
	b := fmt.Sprintf("% x", a)
	return b
} // }}}

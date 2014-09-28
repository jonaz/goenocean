package goenocean

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"

	"github.com/tarm/goserial"
)

func Serial(send chan Packet, recv chan Packet) {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 57600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	go readPackets(s, func(data []byte) {
		reciever(data, recv)
	})

	go sender(s, send)
}

func sender(data io.ReadWriter, send chan Packet) {

	for p := range send {
		_, err := data.Write(p.Encode())
		if err != nil {
			log.Fatal(err)
		}
	}

}

func reciever(data []byte, recv chan Packet) {
	p, err := Decode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	recv <- p
}

func readPackets(rd io.ReadWriter, f func([]byte)) {
	//TODO use this example receivePacket when reading from serial https://github.com/kleckse/enocean/blob/master/esp3.py

	buf := make([]byte, 1)
	var rawPacket []byte
	state := 0
	var length int

	for {
		readLen, err := rd.Read(buf)
		if err != nil {
			fmt.Println("ERROR reading:", err)
			continue
		}

		fmt.Printf("% x ", buf)
		//continue

		if readLen > 0 && buf[0] == 0x55 {
			rawPacket = []byte{}
			state = 1
		}

		rawPacket = append(rawPacket, buf...)

		switch state {
		case 1: //header
			if len(rawPacket) > 5 {
				length = int(binary.BigEndian.Uint16(rawPacket[1:3]))
				state = 2
			}

		case 2: // the rest!
			if len(rawPacket) > 5+length+int(rawPacket[3]) {
				state = 3
			}

		case 3: //data crc
			state = 0
			f(rawPacket)
		}

	}

}

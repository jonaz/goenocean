package goenocean

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"

	"github.com/tarm/goserial"
)

type Encoder interface {
	Encode() []byte
}

type request struct {
	Encoder
	responseChannel chan Packet
}

type Request interface {
	Encoder
	ResponseChannel() chan Packet
}

func (r *request) ResponseChannel() chan Packet {
	return r.responseChannel
}

func NewRequest(p Encoder) Request {
	return &request{p, make(chan Packet)}
}

func Serial(send chan Request, recv chan Packet) {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 57600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	response := make(chan Packet)

	go readPackets(s, func(data []byte) {
		reciever(data, recv, response)
	})

	go sender(s, send, response)
}

func sender(data io.ReadWriter, send chan Request, response chan Packet) {

	for r := range send {
		_, err := data.Write(r.Encode())
		response <- nil
		r.ResponseChannel() <- <-response //TODO test this. might work :)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func reciever(data []byte, recv chan Packet, response chan Packet) {
	p, err := Decode(data)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", p.Header())
	fmt.Printf("Data: %#v\n", p.Data())
	if err != nil {
		fmt.Println(err)
		return
	}

	select {
	case <-response:
		response <- p
	default:
		recv <- p
	}
}

func readPackets(rd io.ReadWriter, f func([]byte)) {

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

		//TODO add debug here seelog
		//fmt.Printf("% x ", buf)
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

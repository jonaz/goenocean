package goenocean

import (
	"encoding/binary"
	"io"

	log "github.com/cihub/seelog"
	"github.com/tarm/goserial"
)

type Encoder interface {
	Encode() []byte
}

func Serial(send chan Encoder, recv chan Packet) {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 57600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Critical(err)
	}

	go readPackets(s, func(data []byte) {
		reciever(data, recv)
	})

	go sender(s, send)
}

func sender(data io.ReadWriter, send chan Encoder) {

	for p := range send {
		_, err := data.Write(p.Encode())
		if err != nil {
			log.Critical(err)
		}
	}

}

func reciever(data []byte, recv chan Packet) {
	p, err := Decode(data)
	log.Debugf("%#v\n", p)
	log.Debugf("%#v\n", p.Header())
	log.Debugf("Data: %#v\n", p.Data())
	if err != nil {
		log.Error("Decode failed :", err)
		return
	}
	recv <- p
}

func readPackets(rd io.ReadWriter, f func([]byte)) {

	buf := make([]byte, 1)
	var rawPacket []byte
	state := 0
	var length int

	for {
		readLen, err := rd.Read(buf)
		if err != nil {
			log.Error("ERROR reading:", err)
			continue
		}

		//TODO add debug here seelog
		log.Infof("% x ", buf)
		//continue

		if readLen > 0 && buf[0] == 0x55 && state == 0 {
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

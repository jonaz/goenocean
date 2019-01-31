package goenocean

import (
	"bytes"
	"encoding/binary"
	"io"
	"time"

	log "github.com/cihub/seelog"
	serial "github.com/tarm/goserial"
)

type Encoder interface {
	Encode() []byte
}

func Serial(dev string, send chan Encoder, recv chan Packet) error {
	c := &serial.Config{Name: dev, Baud: 57600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}

	response := make(chan Packet, 100)
	go readPackets(s, func(data []byte) {
		reciever(data, recv, response)
	})

	go sender(s, send, response)
	return nil
}

func sender(data io.Writer, send chan Encoder, response chan Packet) {

	for p := range send {
		gotResponse := make(chan struct{})
		go waitForResponse(gotResponse, response)
		_, err := data.Write(p.Encode())
		//Dont send next until we have a response from the last one
		<-gotResponse
		if err != nil {
			log.Critical(err)
		}
	}
}

func waitForResponse(weGotResponse chan struct{}, response chan Packet) {
	select {
	case p := <-response:
		log.Debugf("We got response after send: % x\n", p.Encode())
		if !bytes.Equal(p.Data(), []byte{0}) {
			log.Errorf("We got RESPONSE error after send: % x\n", p.Encode())
		}
		weGotResponse <- struct{}{}
		return
	case <-time.After(time.Second * 2):
		log.Error("We got TIMOUT after send")
		weGotResponse <- struct{}{}
		return
	}

}

func reciever(data []byte, recv chan Packet, resp chan Packet) {
	p, err := Decode(data)
	log.Debugf("%#v\n", p)
	log.Debugf("%#v\n", p.Header())
	log.Debugf("Data: %#v\n", p.Data())
	if err != nil {
		log.Error("Decode failed :", err)
		return
	}
	if p.PacketType() == PacketTypeResponse {
		select { // Nonblocking send just in case
		case resp <- p:
		default:
		}
	}
	recv <- p
}

func readPackets(rd io.Reader, f func([]byte)) {

	buf := make([]byte, 1)
	var rawPacket []byte
	state := 0
	var length int

	for {
		readLen, err := rd.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Error("ERROR reading:", err)
			continue
		}

		log.Debugf("% x ", buf)

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

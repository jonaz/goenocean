package goenocean

import (
	"bytes"
	"testing"
)

type ReaderStub struct {
}

func (rs *ReaderStub) Read(p []byte) (n int, err error) {

	return 1, nil
}
func TestReadPackets(t *testing.T) {

	rd := bytes.NewBuffer([]byte{0x00, 0x01, 0x55, 0x00, 0x07, 0x07, 0x01, 0x7a, 0xf6, 0x70, 0x01, 0x8d, 0xf8, 0xbd, 0x30, 0x01, 0xff, 0xff, 0xff, 0xff, 0x55, 0x00, 0x8b, 0x02, 0x04})

	readPackets(rd, func(data []byte) {
		expected := []byte{0x55, 0x00, 0x07, 0x07, 0x01, 0x7a, 0xf6, 0x70, 0x01, 0x8d, 0xf8, 0xbd, 0x30, 0x01, 0xff, 0xff, 0xff, 0xff, 0x55, 0x00, 0x8b}
		if !bytes.Equal(data, expected) {
			t.Errorf("Wrong packet content\nGot\t%s\nExpected\t%s", ToHex(data), ToHex(expected))
			return
		}

		_, err := Decode(data)
		if err != nil {
			t.Error("Error decoding packet", ToHex(data))
		}

	})
	//if p.OutputValue() != 55 {
	//t.Errorf("OutputValue wrong expected: %s got %s", 55, p.OutputValue())
	//}
}
func TestReadPacketsType2(t *testing.T) {

	// OK packet and another one right after
	//55 00 01 00 02 65 00 00 55 00 07 07 01 7a f6 70 01 8d 88 bd 30 01 ff ff ff ff 4f 00 a9
	rd := bytes.NewBuffer([]byte{0x55, 0x00, 0x01, 0x00, 0x02, 0x65, 0x00, 0x00, 0x55, 0x00, 0x07, 0x07, 0x01, 0x7a, 0xf6, 0x70, 0x01, 0x8d, 0x88, 0xbd, 0x30, 0x01, 0xff, 0xff, 0xff, 0xff, 0x4f, 0x00, 0xa9})

	cnt := 0
	readPackets(rd, func(data []byte) {
		t.Logf("Packet: %s\n", ToHex(data))
		_, err := Decode(data)
		if err != nil {
			t.Error("Error decoding packet", ToHex(data))
		}
		cnt++

	})
	if cnt != 2 {
		t.Error("Expected to have found 2 packets")
	}
}

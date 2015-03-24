package goenocean

import (
	"bytes"
	"fmt"
	"testing"
)

type ReaderStub struct {
}

func (rs *ReaderStub) Read(p []byte) (n int, err error) {

	return 1, nil
}
func TestReadPackets(t *testing.T) {

	rd := bytes.NewBuffer([]byte{0x55, 0x00, 0x07, 0x07, 0x01, 0x7a, 0xf6, 0x70, 0x01, 0x8d, 0xf8, 0xbd, 0x30, 0x01, 0xff, 0xff, 0xff, 0xff, 0x55, 0x00, 0x8b})

	readPackets(rd, func(data []byte) {
		fmt.Println(ToHex(data))
	})
	//if p.OutputValue() != 55 {
	//t.Errorf("OutputValue wrong expected: %s got %s", 55, p.OutputValue())
	//}
}

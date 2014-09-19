package goEnocean

import "testing"

func TestEncode(t *testing.T) {

	p := NewPacket()
	p.Data = []byte{0x01, 0x02, 0x03}
	p.OptData = []byte{0x04, 0x05, 0x06}
	encoded := p.Encode()
	if toHex(encoded) != "55 00 03 03 00 82 01 02 03 04 05 06 2f" {
		t.Errorf("encoding failed: %s", toHex(encoded))
	}
}
func TestEncodeCO_WR_RESET(t *testing.T) {

	p := NewPacket()
	p.setPacketType(0x05)
	p.Data = []byte{0x02}
	encoded := p.Encode()
	if toHex(encoded) != "55 00 01 00 05 70 02 0e" {
		t.Errorf("encoding failed: %s", toHex(encoded))
	}
}
func TestEncodeCO_RD_IDBASE(t *testing.T) {
	p := NewPacket()
	p.setPacketType(0x05)
	p.Data = []byte{0x08}
	encoded := p.Encode()
	if toHex(encoded) != "55 00 01 00 05 70 08 38" {
		t.Errorf("encoding failed: %s", toHex(encoded))
	}
}

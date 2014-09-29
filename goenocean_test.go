package goenocean

import "testing"

func TestEncode(t *testing.T) {

	p := NewPacket()
	p.SetData([]byte{0x01, 0x02, 0x03})
	p.SetOptData([]byte{0x04, 0x05, 0x06})
	encoded := p.Encode()
	if ToHex(encoded) != "55 00 03 03 00 82 01 02 03 04 05 06 2f" {
		t.Errorf("encoding failed: %s", ToHex(encoded))
	}
}
func TestEncodeCO_WR_RESET(t *testing.T) {

	p := NewPacket()
	p.SetPacketType(0x05)
	p.SetData([]byte{0x02})
	encoded := p.Encode()
	if ToHex(encoded) != "55 00 01 00 05 70 02 0e" {
		t.Errorf("encoding failed: %s", ToHex(encoded))
	}
}
func TestEncodeCO_RD_IDBASE(t *testing.T) {
	p := NewPacket()
	p.SetPacketType(0x05)
	p.SetData([]byte{0x08})
	encoded := p.Encode()
	if ToHex(encoded) != "55 00 01 00 05 70 08 38" {
		t.Errorf("encoding failed: %s", ToHex(encoded))
	}
}
func TestSetGetPacketType(t *testing.T) {
	p := NewPacket()
	p.SetPacketType(PacketTypeRadioErp1)
	if p.PacketType() != PacketTypeRadioErp1 {
		t.Errorf("wrong PacketType: %s", p.PacketType())
	}
}
func TestDecodeBrokenPackage(t *testing.T) {
	pkg, err := Decode([]byte{0x00})
	if err != nil {
		return
	}
	t.Errorf("package Decode failed: %s", pkg)
}
func TestDecode(t *testing.T) {
	p := NewPacket()
	p.SetPacketType(0x05)
	p.SetData([]byte{0x02})
	encoded := p.Encode()

	pkg, err := Decode(encoded)
	if err != nil {
		t.Errorf("Decoding failed with error: %s", err)
	}
	if !p.Equal(pkg) {
		//if !reflect.DeepEqual(p, pkg) {
		t.Errorf("Packets not equal: \n%+v\n%+v", p, pkg)
	}
}

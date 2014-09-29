package goenocean

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeRpsTelegram(t *testing.T) {
	p := NewTelegramRps()
	p.SetSenderId([4]byte{0xfe, 0xfe, 0x74, 0x9b})
	p.SetTelegramData(0x70)
	p.SetStatus(3)

	encoded := p.Encode()

	if ToHex(encoded) != "55 00 07 07 01 7a f6 70 fe fe 74 9b 03 03 ff ff ff ff ff 00 17" {
		t.Errorf("encoding failed: %s", ToHex(encoded))
	}
}

func TestRpsTelegramData(t *testing.T) {
	p := NewTelegramRps()
	p.SetSenderId([4]byte{0xfe, 0xfe, 0x74, 0x9b})
	p.SetTelegramData(0x70)
	p.SetStatus(3)

	if !bytes.Equal(p.TelegramData(), []byte{0x70}) {
		t.Errorf("wrong data failed: %v != %v", p.Data(), []byte{0x70})
	}
}

func TestDecodeRpsTelegram(t *testing.T) {
	p := NewTelegramRps()
	p.SetSenderId([4]byte{0xfe, 0xfe, 0x74, 0x9b})
	p.SetTelegramData(0x70)
	p.SetStatus(3)

	encoded := p.Encode()

	pkg, err := Decode(encoded)
	if err != nil {
		t.Errorf("Decoding failed with error: %s", err)
	}
	//Here we also test that our Equal does the same job as reflect.DeepEqual
	if !p.Equal(pkg) && !reflect.DeepEqual(p, pkg) {
		t.Errorf("Packets not equal: \n%+v\n%+v\n%+v\n%+v", p, pkg, p.Header(), pkg.Header())
	}
}

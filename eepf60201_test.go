package goenocean

import (
	"fmt"
	"testing"
)

func TestEepf60201(t *testing.T) {
	p := NewEepF60201()
	p.SetSenderId([4]byte{0xfe, 0xfe, 0x74, 0x9b})
	p.SetTelegramData([]byte{0x70}) // ON
	p.SetStatus(3)

	p.SetT21(true)
	p.SetNu(true)

	//encoded := p.Encode()

	fmt.Println("repeat:", p.RepeatCount())
	fmt.Println("T21:", p.T21())
	fmt.Println("NU:", p.Nu())
	fmt.Println("EB:", p.EnergyBow())
	fmt.Printf("status: %b\n", p.Status())
	//fmt.Println(p.Action())

	//if ToHex(encoded) != "55 00 07 07 01 7a f6 70 fe fe 74 9b 03 03 ff ff ff ff ff 00 17" {
	//t.Errorf("encoding failed: %s", ToHex(encoded))
	//}
}

package goenocean

import "fmt"

type EepF60201 struct {
	*TelegramRps
}

func NewEepF60201() *EepF60201 { // {{{
	return &EepF60201{NewTelegramRps()}
} // }}}

func (p *EepF60201) RepeatCount() uint8 { // {{{
	return p.status & 0x0f
}                                                 // }}}
func (p *EepF60201) SetRepeatCount(count uint8) { // {{{
	//TODO find out how to set bit 210 here http://stackoverflow.com/questions/23192262/how-would-you-set-and-clear-a-single-bit-in-go http://stackoverflow.com/questions/4439078/how-do-you-set-only-certain-bits-of-a-byte-in-c-without-affecting-the-rest
	p.status &^= 0x0f        //zero first 4 bits
	p.status |= count & 0x0f //set the 4 bits from count
} // }}}

func (p *EepF60201) T21() bool { // {{{
	data := (p.status >> 5) & 0x01
	if data == 1 {
		return true
	}
	return false
}                                       // }}}
func (p *EepF60201) SetT21(data bool) { // {{{
	if data {
		p.status = setBit(p.status, 5)
		return
	}
	p.status = clearBit(p.status, 5)
} // }}}

func (p *EepF60201) Nu() bool { // {{{
	data := (p.status >> 4) & 0x01
	if data == 1 {
		return true
	}
	return false
}                                      // }}}
func (p *EepF60201) SetNu(data bool) { // {{{
	if data {
		p.status = setBit(p.status, 4)
		return
	}
	p.status = clearBit(p.status, 4)
}                                      // }}}
func (p *EepF60201) EnergyBow() bool { // {{{
	data := (p.data >> 4) & 0x01
	if data == 1 {
		return true
	}
	return false
} // }}}

func (p *EepF60201) R1B0() bool {
	if p.r1() == 3 {
		return true
	}
	return false
}

func (p *EepF60201) r1() uint {
	n := p.data >> 5
	return uint(n)
}

func (p *EepF60201) Action() string {
	//@flags = {:t21 => (@status >> 5) & 0x01, :nu => (@status >> 4) & 0x01 }
	fmt.Printf("raw action: %b\n", p.data)
	n := p.data >> 5
	fmt.Printf("bit action: %b\n", n)
	switch n {
	case 0:
		return "Button A-ON"
	case 1:
		return "Button A-OFF"
	case 2:
		return "Button B-ON"
	case 3:
		return "Button B-OFF"

	}
	return "INGET"
}

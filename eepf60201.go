package goenocean

import "fmt"

type EepF60201 struct {
	*TelegramRps
}

func NewEepF60201() *EepF60201 { // {{{
	return &EepF60201{NewTelegramRps()}
} // }}}

func (p *EepF60201) SetTelegram(t *TelegramRps) { // {{{
	p.TelegramRps = t
} // }}}

func (p *EepF60201) T21() bool { // {{{
	data := (p.Status() >> 5) & 0x01
	if data == 1 {
		return true
	}
	return false
}                                       // }}}
func (p *EepF60201) SetT21(data bool) { // {{{
	if data {
		p.SetStatus(setBit(p.Status(), 5))
		return
	}
	p.SetStatus(clearBit(p.Status(), 5))
} // }}}

func (p *EepF60201) Nu() bool { // {{{
	data := (p.Status() >> 4) & 0x01
	if data == 1 {
		return true
	}
	return false
}                                      // }}}
func (p *EepF60201) SetNu(data bool) { // {{{
	if data {
		p.SetStatus(setBit(p.Status(), 4))
		return
	}
	p.SetStatus(clearBit(p.Status(), 4))
}                                      // }}}
func (p *EepF60201) EnergyBow() bool { // {{{
	data := (p.TelegramData()[0] >> 4) & 0x01
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
	n := p.TelegramData()[0] >> 5
	return uint(n)
}

func (p *EepF60201) R2B1() bool {
	if p.r2() == 2 {
		return true
	}
	return false
}
func (p *EepF60201) R2B0() bool {
	if p.r2() == 3 {
		return true
	}
	return false
}
func (p *EepF60201) r2() uint {
	if p.TelegramData()[0]&0x01 == 1 {
		n := (p.TelegramData()[0] >> 1) & 0x07
		return uint(n)
	}
	return 0xff
}

func (p *EepF60201) Action() string {
	//@flags = {:t21 => (@status >> 5) & 0x01, :nu => (@status >> 4) & 0x01 }
	fmt.Printf("raw action: %b\n", p.TelegramData())
	n := p.TelegramData()[0] >> 5
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

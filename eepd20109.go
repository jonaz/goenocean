package goenocean

type EepD20109 struct {
	*TelegramVld
}

func NewEepD20109() *EepD20109 { // {{{
	return &EepD20109{NewTelegramVld()}
} // }}}

func (p *EepD20109) SetTelegram(t *TelegramVld) { // {{{
	p.TelegramVld = t
} // }}}

func (p *EepD20109) CommandId() uint8 {
	return p.data[0] & 0x0f
}
func (p *EepD20109) OutputValue() uint8 {
	return p.data[2] & 0x7f
}
func (p *EepD20109) LocalControl() uint8 {
	return (p.data[2] & 0x80) >> 7
}
func (p *EepD20109) IOChannel() uint8 {
	return p.data[1] & 0x1f
}

//OLD

func (p *EepD20109) RepeatCount() uint8 { // {{{
	return p.status & 0x0f
}                                                 // }}}
func (p *EepD20109) SetRepeatCount(count uint8) { // {{{
	p.status &^= 0x0f        //zero first 4 bits
	p.status |= count & 0x0f //set the 4 bits from count
} // }}}

func (p *EepD20109) T21() bool { // {{{
	data := (p.status >> 5) & 0x01
	if data == 1 {
		return true
	}
	return false
}                                       // }}}
func (p *EepD20109) SetT21(data bool) { // {{{
	if data {
		p.status = setBit(p.status, 5)
		return
	}
	p.status = clearBit(p.status, 5)
} // }}}

func (p *EepD20109) Nu() bool { // {{{
	data := (p.status >> 4) & 0x01
	if data == 1 {
		return true
	}
	return false
}                                      // }}}
func (p *EepD20109) SetNu(data bool) { // {{{
	if data {
		p.status = setBit(p.status, 4)
		return
	}
	p.status = clearBit(p.status, 4)
} // }}}
//func (p *EepD20109) EnergyBow() bool { // {{{
//data := (p.data >> 4) & 0x01
//if data == 1 {
//return true
//}
//return false
//} // }}}

//func (p *EepD20109) R1B0() bool {
//if p.r1() == 3 {
//return true
//}
//return false
//}

//func (p *EepD20109) r1() uint {
//n := p.data >> 5
//return uint(n)
//}

//func (p *EepD20109) R2B1() bool {
//if p.r2() == 2 {
//return true
//}
//return false
//}
//func (p *EepD20109) R2B0() bool {
//if p.r2() == 3 {
//return true
//}
//return false
//}

//func (p *EepD20109) r2() uint {
//if p.data&0x01 == 1 {
//n := (p.data >> 1) & 0x07
//return uint(n)
//}
//return 0xff
//}

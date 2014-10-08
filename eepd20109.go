package goenocean

type EepD20109 struct {
	TelegramVld
}

func NewEepD20109() *EepD20109 { // {{{
	return &EepD20109{NewTelegramVld()}
} // }}}

func (p *EepD20109) SetTelegram(t TelegramVld) { // {{{
	p.TelegramVld = t
} // }}}

func (p *EepD20109) CommandId() uint8 {
	return p.TelegramData()[0] & 0x0f
}
func (p *EepD20109) SetCommandId(id uint8) {
	//TODO also initialize TelegramData here with the correct size for correct commandId
	//p.SetTelegramData([]byte{0, 0, 0})
	// 1 : 3 bytes
	// 2 : 4 bytes
	// 3 : 2 bytes
	// 4 : 3 bytes
	// 5 : 6 bytes
	// 6 : 2 bytes
	// 7 : 6 bytes
	//return p.TelegramData()[0] & 0x0f
}
func (p *EepD20109) OutputValue() uint8 {
	return p.TelegramData()[2] & 0x7f
}
func (p *EepD20109) SetOutputValue(count uint8) {
	//TODO validate only 0-100 DECIMAL here
	tmp := p.TelegramData()
	tmp[2] &^= 0x7f        //zero first 4 bits
	tmp[2] |= count & 0x7f //set the 7 bits from count
	p.SetTelegramData(tmp)
}
func (p *EepD20109) LocalControl() uint8 {
	return (p.TelegramData()[2] & 0x80) >> 7
}
func (p *EepD20109) IOChannel() uint8 {
	return p.TelegramData()[1] & 0x1f
}

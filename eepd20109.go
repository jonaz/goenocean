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
func (p *EepD20109) OutputValue() uint8 {
	return p.TelegramData()[2] & 0x7f
}
func (p *EepD20109) LocalControl() uint8 {
	return (p.TelegramData()[2] & 0x80) >> 7
}
func (p *EepD20109) IOChannel() uint8 {
	return p.TelegramData()[1] & 0x1f
}

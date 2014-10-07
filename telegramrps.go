package goenocean

type TelegramRps struct {
	*telegram
}

func NewTelegramRps() *TelegramRps {
	t := &TelegramRps{telegram: NewTelegram()}
	t.telegramType = TelegramTypeRps
	return t
}

func (p *TelegramRps) RepeatCount() uint8 { // {{{
	return p.Status() & 0x0f
}                                                   // }}}
func (p *TelegramRps) SetRepeatCount(count uint8) { // {{{
	p.status &^= 0x0f        //zero first 4 bits
	p.status |= count & 0x0f //set the 4 bits from count
} // }}}

package goenocean

type telegramRps struct {
	*telegram
}
type TelegramRps interface {
	Telegram
	RepeatCount() uint8
	SetRepeatCount(uint8)
}

func NewTelegramRps() TelegramRps {
	t := &telegramRps{telegram: NewTelegram()}
	t.telegramType = TelegramTypeRps
	return t
}

func (p *telegramRps) RepeatCount() uint8 { // {{{
	return p.Status() & 0x0f
}                                                   // }}}
func (p *telegramRps) SetRepeatCount(count uint8) { // {{{
	p.status &^= 0x0f        //zero first 4 bits
	p.status |= count & 0x0f //set the 4 bits from count
} // }}}

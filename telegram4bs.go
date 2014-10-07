package goenocean

type Telegram4bs struct {
	*telegram
}

func NewTelegram4bs() *Telegram4bs {
	t := &Telegram4bs{telegram: NewTelegram()}
	t.telegramType = TelegramType4bs
	return t
}
func (p *Telegram4bs) TelegramData() []byte {
	return p.data
}

func (p *Telegram4bs) SetTelegramData(data []byte) {
	p.data = data
}

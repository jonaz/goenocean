package goenocean

type TelegramVld struct {
	*telegram
}

func NewTelegramVld() *TelegramVld {
	t := &TelegramVld{telegram: NewTelegram()}
	t.telegramType = TelegramTypeVld
	return t
}

func (p *TelegramVld) TelegramData() []byte {
	return p.data
}

func (p *TelegramVld) SetTelegramData(data []byte) {
	p.data = data
}

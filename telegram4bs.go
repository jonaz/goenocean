package goenocean

type Telegram4bs struct {
	*telegram
}

func NewTelegram4bs() *Telegram4bs {
	t := &Telegram4bs{telegram: NewTelegram()}
	t.telegramType = TelegramType4bs
	return t
}

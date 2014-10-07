package goenocean

type telegram4bs struct {
	*telegram
}
type Telegram4bs interface {
	Telegram
}

func NewTelegram4bs() Telegram4bs {
	t := &telegram4bs{telegram: NewTelegram()}
	t.telegramType = TelegramType4bs
	return t
}

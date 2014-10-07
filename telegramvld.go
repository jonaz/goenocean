package goenocean

type TelegramVld struct {
	*telegram
}

func NewTelegramVld() *TelegramVld {
	t := &TelegramVld{telegram: NewTelegram()}
	t.telegramType = TelegramTypeVld
	return t
}

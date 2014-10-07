package goenocean

type telegramVld struct {
	*telegram
}

type TelegramVld interface {
	Telegram
}

func NewTelegramVld() TelegramVld {
	t := &telegramVld{telegram: NewTelegram()}
	t.telegramType = TelegramTypeVld
	return t
}

package goenocean

type telegram4bs struct {
	*telegram
}
type Telegram4bs interface {
	Telegram
}

func NewTelegram4bs() Telegram4bs {
	t := &telegram4bs{telegram: NewTelegram()}
	t.SetTelegramData(make([]byte, 4))
	t.telegramType = TelegramType4bs
	//we dont want to be in learning mode
	t.setLearn(false)
	return t
}
func (p *telegram4bs) setLearn(lrn bool) {
	var data uint8
	if lrn {
		data = 0
	} else {
		data = 1
	}
	tmp := p.TelegramData()
	tmp[3] &^= 0x08
	tmp[3] |= (data << 3) & 0x08
	p.SetTelegramData(tmp)
	// 0 : teach in telegram
	// 1 : data telegram
}

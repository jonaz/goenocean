package goenocean

type Telegram4bsLearn struct {
	Telegram
}

// A5-38-08 CENTRAL COMMAND
func NewTelegram4bsLearn() *Telegram4bsLearn { // {{{
	t := &Telegram4bsLearn{NewTelegram4bs()}
	t.SetLearn(true)
	return t
} // }}}

func (p *Telegram4bsLearn) SetTelegram(t Telegram) { // {{{
	p.Telegram = t
} // }}}

func (p *Telegram4bsLearn) Learn() bool { // {{{
	learnBit := (p.TelegramData()[3] >> 3) & 0x01
	if learnBit == 0 {
		return true
	}
	return false
	// 0 : teach in telegram
	// 1 : data telegram
}                                               // }}}
func (p *Telegram4bsLearn) SetLearn(lrn bool) { // {{{
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
} // }}}

func (p *Telegram4bsLearn) LearnFunc() byte {
	return (p.TelegramData()[0] & 0xf8) >> 2
}
func (p *Telegram4bsLearn) SetLearnFunc(lfunc byte) {
	tmp := p.TelegramData()
	tmp[0] &^= 0xf8               //zero bits
	tmp[0] |= (lfunc << 2) & 0xf8 // set bits
	p.SetTelegramData(tmp)
}

func (p *Telegram4bsLearn) LearnType() byte { // {{{
	tmp := p.TelegramData()
	byte0 := (tmp[0] & 0x03) << 5
	byte1 := (tmp[1] & 0xf8) >> 3
	return byte0 | byte1
}                                                     // }}}
func (p *Telegram4bsLearn) SetLearnType(ltype byte) { // {{{
	tmp := p.TelegramData()

	byte0 := (ltype & 0x60) >> 5 //01100000
	byte1 := (ltype & 0x1f) << 3 //00011111

	//byte0
	tmp[0] &^= 0x03          //zero bits
	tmp[1] &^= 0xf8          //zero bits
	tmp[0] |= (byte0) & 0x03 // set bits
	tmp[1] |= (byte1) & 0xf8 // set bits

	//also set the learn type bit to true here
	tmp[3] = tmp[3] | 0x80

	p.SetTelegramData(tmp)
} // }}}

package goenocean

type EepA53808 struct {
	//*Telegram4bs
	Telegram
}

// A5-38-08 CENTRAL COMMAND
func NewEepA53808() *EepA53808 { // {{{
	t := &EepA53808{NewTelegram4bs()}
	//Must default to empty bytearray
	t.SetTelegramData(make([]byte, 4))
	return t
} // }}}

func (p *EepA53808) SetTelegram(t Telegram) { // {{{
	p.Telegram = t
} // }}}

func (p *EepA53808) Command() uint8 {
	return p.TelegramData()[0]
}
func (p *EepA53808) SetCommand(cmd uint8) {
	tmp := p.TelegramData()
	tmp[0] = cmd
	p.SetTelegramData(tmp)
}

//TODO check if Learn can be moved to Telegram4bs since its always at the same bit
func (p *EepA53808) Learn() bool {
	// 0 : teach in telegram
	// 1 : data telegram
	return true
}
func (p *EepA53808) SetLearn(lrn bool) {
	// 0 : teach in telegram
	// 1 : data telegram
}

//func (p *EepA53808) TariffInfo() uint8 {
//ti := (p.TelegramData()[3] & 0x0f0) >> 4
//return ti
//}

//func (p *EepA53808) DataType() string {
//d := (p.TelegramData()[3] >> 2) & 0x01

//switch d {
//case 0:
//return "kWh"
//case 1:
//return "W"
//}
//return "Unknown"
//}

//func (p *EepA53808) dividor() int64 {
//d := p.TelegramData()[3] & 0x03
//switch d {
//case 0:
//return 1
//case 1:
//return 10
//case 2:
//return 100
//case 3:
//return 1000
//}
//return 1
//}

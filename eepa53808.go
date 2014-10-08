package goenocean

type EepA53808 struct {
	//*Telegram4bs
	Telegram
}

// A5-38-08 CENTRAL COMMAND
func NewEepA53808() *EepA53808 { // {{{
	t := &EepA53808{NewTelegram4bs()}
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
func (p *EepA53808) Time() uint8 {
	return p.TelegramData()[1]
}
func (p *EepA53808) SetTime(time uint8) {
	tmp := p.TelegramData()
	tmp[1] = time
	p.SetTelegramData(tmp)
}
func (p *EepA53808) DimValue() uint8 {
	return p.TelegramData()[1]
}
func (p *EepA53808) SetDimValue(val uint8) {
	tmp := p.TelegramData()
	tmp[1] = val
	p.SetTelegramData(tmp)
}
func (p *EepA53808) RampTime() uint8 {
	return p.TelegramData()[2]
}
func (p *EepA53808) SetRampTime(val uint8) {
	tmp := p.TelegramData()
	tmp[2] = val
	p.SetTelegramData(tmp)
}
func (p *EepA53808) SwitchingCommand() uint8 {
	return p.TelegramData()[3] & 0x01
}
func (p *EepA53808) SetSwitchingCommand(cmd uint8) {
	tmp := p.TelegramData()
	tmp[3] &^= 0x01      //zero first 1 bits
	tmp[3] |= cmd & 0x01 //set the 1 bits from count
	p.SetTelegramData(tmp)
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

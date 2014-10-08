package goenocean

type EepD20109 struct {
	TelegramVld
}

func NewEepD20109() *EepD20109 { // {{{
	return &EepD20109{NewTelegramVld()}
} // }}}

func (p *EepD20109) SetTelegram(t TelegramVld) { // {{{
	p.TelegramVld = t
} // }}}

func (p *EepD20109) CommandId() uint8 {
	return p.TelegramData()[0] & 0x0f
}
func (p *EepD20109) SetCommandId(id uint8) {
	sizeMap := make(map[byte]int)
	sizeMap[0x01] = 3
	sizeMap[0x02] = 4
	sizeMap[0x03] = 2
	sizeMap[0x04] = 3
	sizeMap[0x05] = 6
	sizeMap[0x06] = 2
	sizeMap[0x07] = 7

	tmp := make([]byte, sizeMap[id])
	tmp[0] &^= 0x0f     //zero first 4 bits
	tmp[0] |= id & 0x0f //set the 4 bits from count
	p.SetTelegramData(tmp)
}
func (p *EepD20109) OutputValue() uint8 {
	return p.TelegramData()[2] & 0x7f
}
func (p *EepD20109) SetOutputValue(count uint8) {
	//TODO validate only 0-100 DECIMAL here
	tmp := p.TelegramData()
	tmp[2] &^= 0x7f        //zero first 4 bits
	tmp[2] |= count & 0x7f //set the 7 bits from count
	p.SetTelegramData(tmp)
}

//DimValue only valid for command id 1
func (p *EepD20109) DimValue() uint8 {
	return (p.TelegramData()[1] & 0xe0) >> 5
}

//Can be between 0 and 4.
func (p *EepD20109) SetDimValue(count uint8) {
	tmp := p.TelegramData()
	tmp[1] &^= 0xe0               //zero the bits
	tmp[1] |= (count << 5) & 0xe0 //set the bits from count
	p.SetTelegramData(tmp)
}

func (p *EepD20109) LocalControl() uint8 {
	return (p.TelegramData()[2] & 0x80) >> 7
}
func (p *EepD20109) IOChannel() uint8 {
	return p.TelegramData()[1] & 0x1f
}
func (p *EepD20109) SetIOChannel(count uint8) {
	tmp := p.TelegramData()
	tmp[1] &^= 0x1f        //zero first 5 bits
	tmp[1] |= count & 0x1f //set the 5 bits from count
	p.SetTelegramData(tmp)
}

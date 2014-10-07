package goenocean

import "math/big"

type EepA51201 struct {
	//*Telegram4bs
	Telegram
}

func NewEepA51201() *EepA51201 { // {{{
	return &EepA51201{NewTelegram4bs()}
} // }}}

func (p *EepA51201) SetTelegram(t Telegram) { // {{{
	p.Telegram = t
} // }}}

func (p *EepA51201) MeterReading() int64 {
	i := new(big.Int)
	i.SetBytes(p.TelegramData()[0:3])

	return i.Int64() / p.dividor()
}

func (p *EepA51201) TariffInfo() uint8 {
	ti := (p.TelegramData()[3] & 0x0f0) >> 4
	return ti
}

func (p *EepA51201) DataType() string {
	d := (p.TelegramData()[3] >> 2) & 0x01

	switch d {
	case 0:
		return "kWh"
	case 1:
		return "W"
	}
	return "Unknown"
}

func (p *EepA51201) dividor() int64 {
	d := p.TelegramData()[3] & 0x03
	switch d {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	}
	return 1
}

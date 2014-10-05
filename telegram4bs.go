package goenocean

type Telegram4bs struct {
	*packet
	senderId      [4]byte
	destinationId [4]byte
	data          []byte
	status        byte
}

func NewTelegram4bs() *Telegram4bs {
	return &Telegram4bs{
		packet:        NewPacket(),
		senderId:      [4]byte{0x00, 0x00, 0x00, 0x00}, //this can default to 00 since usb300 will add its own senderid
		destinationId: [4]byte{0xff, 0xff, 0xff, 0xff},
		status:        0,
		data:          []byte{}}
}
func (p *Telegram4bs) Process() {
	length := len(p.packet.data)

	p.status = p.packet.data[length-1]
	copy(p.senderId[:], p.packet.data[length-5:length-1])
	copy(p.destinationId[:], p.packet.optData[1:5])

	//1-14 byte data only for VLD
	//p.data = p.packet.data[1]
	p.data = p.packet.data[1:5]
}

func (p *Telegram4bs) DestinationId() [4]byte {
	return p.destinationId
}
func (p *Telegram4bs) TelegramData() []byte {
	return p.data
}

func (p *Telegram4bs) SetTelegramData(data []byte) {
	p.data = data
}

func (p *Telegram4bs) SenderId() [4]byte {
	return p.senderId
}
func (p *Telegram4bs) SetSenderId(data [4]byte) {
	p.senderId = data
}
func (p *Telegram4bs) SetStatus(data byte) {
	p.status = data
}

func (p *Telegram4bs) Status() byte {
	return p.status
}
func (p *Telegram4bs) Encode() []byte {

	// 1 byte data + 4 byte sender id + 1 byte status
	data := []byte{TelegramType4bs}
	data = append(data, p.data...)
	data = append(data, p.senderId[:]...)
	data = append(data, p.status)

	p.packet.data = data

	//SubTelNum + Destination Id + dBm + security Level
	optData := []byte{0x03}
	optData = append(optData, p.destinationId[:]...)
	optData = append(optData, 0xff)
	optData = append(optData, 0x00)
	p.packet.optData = optData

	p.packet.SetPacketType(0x01)

	return p.packet.Encode()
}

package goenocean

type TelegramVld struct {
	*packet
	senderId      [4]byte
	destinationId [4]byte
	data          []byte
	status        byte
}

func NewTelegramVld() *TelegramVld {
	return &TelegramVld{
		packet:        NewPacket(),
		senderId:      [4]byte{0x00, 0x00, 0x00, 0x00}, //this can default to 00 since usb300 will add its own senderid
		destinationId: [4]byte{0xff, 0xff, 0xff, 0xff},
		status:        0,
		data:          []byte{}}
}
func (p *TelegramVld) Process() {
	length := len(p.packet.data)

	p.status = p.packet.data[length-1]
	copy(p.senderId[:], p.packet.data[length-5:length-1])
	copy(p.destinationId[:], p.packet.optData[1:5])

	//1-14 byte data only for VLD
	//p.data = p.packet.data[1]
	p.data = p.packet.data[1 : length-5]
}

func (p *TelegramVld) DestinationId() [4]byte {
	return p.destinationId
}
func (p *TelegramVld) TelegramData() []byte {
	return p.data
}

func (p *TelegramVld) SetTelegramData(data []byte) {
	p.data = data
}

func (p *TelegramVld) SenderId() [4]byte {
	return p.senderId
}
func (p *TelegramVld) SetSenderId(data [4]byte) {
	p.senderId = data
}
func (p *TelegramVld) SetStatus(data byte) {
	p.status = data
}

func (p *TelegramVld) Status() byte {
	return p.status
}
func (p *TelegramVld) Encode() []byte {

	// 1 byte data + 4 byte sender id + 1 byte status
	data := []byte{TelegramTypeVld}
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

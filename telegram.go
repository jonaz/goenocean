package goenocean

const (
	TelegramTypeRps = 0xf6
	TelegramType1bs = 0xd5
	TelegramType4bs = 0xa5
	TelegramTypeVld = 0xd2
)

type Telegram interface {
	Packet
	Encode() []byte
	//SenderId() [4]byte
	SetSenderId([4]byte)
	TelegramData() []byte
	SetTelegramData([]byte)
	TelegramType() byte
	Status() byte
	SetStatus(byte)
}

type telegram struct {
	*packet
	senderId      [4]byte
	destinationId [4]byte
	data          []byte
	status        byte
	telegramType  byte
}

func NewTelegram() *telegram {
	return &telegram{senderId: [4]byte{0x00, 0x00, 0x00, 0x00}, //this can default to 00 since usb300 will add its own senderid
		packet:        NewPacket(),
		destinationId: [4]byte{0xff, 0xff, 0xff, 0xff},
		status:        0}
}
func (p *telegram) Encode() []byte {

	// 1 byte data + 4 byte sender id + 1 byte status
	data := []byte{p.TelegramType()}
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
func (p *telegram) Process() {
	length := len(p.packet.data)

	p.status = p.packet.data[length-1]
	copy(p.senderId[:], p.packet.data[length-5:length-1])
	copy(p.destinationId[:], p.packet.optData[1:5])

	//1 byte data only for RPS
	//p.data[0] = p.packet.data[1]
	p.data = p.packet.data[1 : length-5]
}

func (p *telegram) TelegramType() byte {
	return p.telegramType
}
func (p *telegram) DestinationId() [4]byte {
	return p.destinationId
}
func (p *telegram) TelegramData() []byte {
	return p.data
}

func (p *telegram) SetTelegramData(data []byte) {
	p.data = data
}

func (p *telegram) SenderId() [4]byte {
	return p.senderId
}
func (p *telegram) SetSenderId(data [4]byte) {
	p.senderId = data
}
func (p *telegram) SetStatus(data byte) {
	p.status = data
}
func (p *telegram) Status() byte {
	return p.status
}

func bits(b uint, subset ...uint) (r uint) {
	i := uint(0)
	for _, v := range subset {
		if b&(1<<v) > 0 {
			r = r | 1<<uint(i)
		}
		i++
	}
	return
}

// Check if a bit at pos is 1 or 0
func hasBit(n byte, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

// Sets the bit at pos in the integer n.
func setBit(n byte, pos uint) byte {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n byte, pos uint) byte {
	mask := ^(1 << pos)
	n &= byte(mask)
	return n
}

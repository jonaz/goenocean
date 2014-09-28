package goenocean

const (
	TelegramTypeRps = 0xf6
	TelegramType1bs = 0xd5
	TelegramType4bs = 0xa5
	TelegramTypeVld = 0xd2
)

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

func hasBit(n byte, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

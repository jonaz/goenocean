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

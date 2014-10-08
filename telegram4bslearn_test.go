package goenocean

import (
	"fmt"
	"testing"
)

func TestTelegram4bsLearn(t *testing.T) {
	p := NewTelegram4bsLearn()
	p.SetSenderId([4]byte{0xfe, 0xfe, 0x74, 0x9b})
	p.SetTelegramData([]byte{0x70, 0x10, 0x10, 0x11})
	p.SetStatus(3)
	p.SetLearn(false)
	if p.Learn() != false {
		t.Errorf("expected: %t got %t", false, p.Learn())
	}
}
func TestTelegram4bsLearnFunc(t *testing.T) {
	p := NewTelegram4bsLearn()
	p.SetSenderId([4]byte{0xfe, 0xfe, 0x74, 0x9b})
	p.SetTelegramData([]byte{0x70, 0x10, 0x10, 0x11})
	p.SetStatus(3)
	p.SetLearn(false)
	p.SetLearnFunc(0x38)
	fmt.Printf("%b\n", p.TelegramData()[0])
	if p.LearnFunc() != 0x38 {
		t.Errorf("expected: %t got %t", 0x38, p.LearnFunc())
	}
}

func TestTelegram4bsLearnType(t *testing.T) {
	p := NewTelegram4bsLearn()
	p.SetSenderId([4]byte{0xfe, 0xfe, 0x74, 0x9b})
	p.SetTelegramData([]byte{0x70, 0x10, 0x10, 0x11})
	p.SetStatus(3)
	p.SetLearn(false)
	p.SetLearnType(0x08)
	if p.LearnType() != 0x08 {
		t.Errorf("LearnType expected: %t got %t", 0x08, p.LearnType())
	}
}

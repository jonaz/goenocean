package goenocean

import "testing"

func TestEepD20109OutputValue(t *testing.T) {
	p := NewEepD20109()
	p.SetTelegramData(make([]byte, 3)) //TODO replace with SetCommandId when done
	p.SetOutputValue(55)
	if p.OutputValue() != 55 {
		t.Errorf("OutputValue wrong expected: %s got %s", 55, p.OutputValue())
	}
}

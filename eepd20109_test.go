package goenocean

import "testing"

func TestEepD20109OutputValue(t *testing.T) {
	p := NewEepD20109()
	p.SetCommandId(1)
	p.SetOutputValue(55)
	if p.OutputValue() != 55 {
		t.Errorf("OutputValue wrong expected: %s got %s", 55, p.OutputValue())
	}
}
func TestEepD20109DimValue(t *testing.T) {
	p := NewEepD20109()
	p.SetCommandId(1)
	p.SetOutputValue(55)
	p.SetDimValue(4)
	if p.OutputValue() != 55 {
		t.Errorf("DimValue wrong expected: %s got %s", 55, p.OutputValue())
	}
	if p.DimValue() != 4 {
		t.Errorf("DimValue wrong expected: %s got %s", 4, p.DimValue())
	}
}
func TestEepD20109IOChannel(t *testing.T) {
	p := NewEepD20109()
	p.SetCommandId(1)
	p.SetIOChannel(30)
	p.SetOutputValue(55)
	p.SetDimValue(4)
	if p.IOChannel() != 30 {
		t.Errorf("IOChannel wrong expected: %s got %s", 4, p.IOChannel())
	}
	p.SetIOChannel(3)
	if p.IOChannel() != 3 {
		t.Errorf("IOChannel wrong expected: %s got %s", 4, p.IOChannel())
	}
}

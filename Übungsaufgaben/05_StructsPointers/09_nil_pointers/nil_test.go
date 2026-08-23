package nilpointers

import "testing"

func TestReadTemperature(t *testing.T) {
	temp := Temperature{Celsius: 21.5}
	value, ok := ReadTemperature(&temp)
	if !ok || value != 21.5 {
		t.Errorf("ReadTemperature = (%v, %v), want (21.5, true)", value, ok)
	}
}

func TestReadTemperatureNil(t *testing.T) {
	value, ok := ReadTemperature(nil)
	if ok || value != 0 {
		t.Errorf("ReadTemperature(nil) = (%v, %v), want (0, false)", value, ok)
	}
}

func TestSetTemperature(t *testing.T) {
	temp := Temperature{Celsius: 20}
	if !SetTemperature(&temp, 25) {
		t.Fatal("SetTemperature should return true")
	}
	if temp.Celsius != 25 {
		t.Errorf("temp.Celsius = %v, want 25", temp.Celsius)
	}
}

func TestSetTemperatureNil(t *testing.T) {
	if SetTemperature(nil, 25) {
		t.Error("SetTemperature(nil, ...) should return false")
	}
}

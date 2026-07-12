package even

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("Ожидалось true для 2, но получено false")
	}

	if IsEven(3) {
		t.Error("Ожидалось false для 3, но получено true")
	}

	if !IsEven(0) {
		t.Error("Ожидалось true для 0, но получено false")
	}

	if !IsEven(-4) {
		t.Error("Ожидалось true для -4, но получено false")
	}

	if IsEven(-7) {
		t.Error("Ожидалось false для -7, но получено true")
	}
}

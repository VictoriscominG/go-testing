package safe

import "testing"

func TestMustAtOk(t *testing.T) {
	data := []int{1, 2, 3, 4}
	got := MustAt(data, 3)
	want := 4
	if got != want {
		t.Errorf("MustAt([]int{1, 2, 3, 4}, 3) = %d, but must be: %d", got, want)
	}
}

func TestMustAtPanicIndexIsNegative(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("we expected panic, but it didn't happen")
		}
		want := "index out of range"
		if r != want {
			t.Errorf("we have panic: %s, but panic must be: %s", r, want)
		}
	}()
	data := []int{1, 2, 3, 4}
	_ = MustAt(data, -2)
}

func TestMustAtPanicIndexOutOfRange(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("we expected panic, but it didn't happen")
		}
		want := "index out of range"
		if r != want {
			t.Errorf("we have panic: %s, but panic must be: %s", r, want)
		}
	}()
	data := []int{1, 2, 3, 4}
	_ = MustAt(data, 7)
}

package textutil

import (
	"strings"
	"testing"
)

func TestNormalizeSpaces(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"empty string", "", ""},
		{"spaces/tabs only", " \n  \t\n ", ""},
		{"mixed spaces and text", "  one\ttwo\nthree \rfour", "one two three four"},
		{"unicode and punctuation", "hello,  world! ", "hello, world!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NormalizeSpaces(tt.in)
			assertValue(t, got, tt.want)
		})
	}
}

func assertValue[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestCountLines(t *testing.T) {
	tests2 := []struct {
		name, in string
		want     int
		err      bool
	}{
		{"empty string", "", 0, false},
		{"one string without \n", "hello", 1, false},
		{"several string", "hello\nworld\n", 2, false},
		{"several string without \n in thee end", "hello\nworld", 2, false},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(tt.in)
			got, err := CountLines(r)
			assertError(t, err, tt.err)
			assertValue(t, got, tt.want)
		})
	}
}

func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if want {
		if got == nil {
			t.Error("expected an error, but got nil")
		}
	} else {
		if got != nil {
			t.Errorf("expected no error, but got: %v", got)
		}
	}
}

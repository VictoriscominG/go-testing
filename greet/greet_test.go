package greet

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
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

func TestHello_WithHelper(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		got, err := Hello("Go")
		assertError(t, err, false)
		assertString(t, got, "Hello, Go")

		got, err = Hello("World")
		assertError(t, err, false)
		assertString(t, got, "Hello, World")
	})

	t.Run("empty", func(t *testing.T) {
		_, err := Hello("")
		assertError(t, err, true)
	})

	t.Run("unicode", func(t *testing.T) {
		got, err := Hello("Гофер")
		assertError(t, err, false)
		assertString(t, got, "Hello, Гофер")
	})

	t.Run("spaces are preserved", func(t *testing.T) {
		got, err := Hello(" Go ")
		assertError(t, err, false)
		assertString(t, got, "Hello,  Go ")
	})
}

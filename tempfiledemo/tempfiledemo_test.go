package tempfiledemo

import (
	"os"
	"strings"
	"testing"
)

func TestWriteLinesToTemp(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		lines  []string
	}{
		{name: "one_word",
			prefix: "test_*.txt",
			lines:  []string{"hello"}},

		{name: "three_words",
			prefix: "test_*.txt",
			lines:  []string{"hello", "my", "litle", "friend"}},

		{name: "empty_file",
			prefix: "test_*.txt",
			lines:  []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := WriteLinesToTemp(tt.prefix, tt.lines)
			if err != nil {
				t.Fatalf("WriteLinesToTemp() error: %v", err)
			}
			defer os.Remove(path)

			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("os.ReadFile() error: %v", err)
			}
			got := string(data)
			want := strings.Join(tt.lines, "\n")
			assertString(t, got, want)
		})
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want: %q, but got: %q", want, got)
	}
}

package textstat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name, value string
		want        map[string]int
	}{
		{"empty string", "   ", map[string]int{}},
		{"repeat word", "Hello, world! hello.", map[string]int{"hello": 2, "world": 1}},
		{"multiple separators", "world,,!- world", map[string]int{"world": 2}},
		{"only separators", "!!!,,, ", map[string]int{}},
		{"numbers", "123 abc 123", map[string]int{"123": 2, "abc": 1}},
		{"unicode", "Привет hello Привет", map[string]int{"привет": 2, "hello": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordCount(tt.value)
			assert.Equal(t, tt.want, got, "WordCount mismatch for: %s", tt.name)
		})
	}
}

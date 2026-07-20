package sluggy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	tests := []struct {
		name       string
		gotString  string
		wantString string
	}{
		{
			name:       "spaces and punctuation",
			gotString:  "go,hello world",
			wantString: "go-hello-world",
		}, {
			name:       "duplicate separators",
			gotString:  "hello, world!",
			wantString: "hello-world",
		}, {
			name:       "unicode letters",
			gotString:  "Привет жителям планЕтЫ HD 137010 b",
			wantString: "привет-жителям-планеты-hd-137010-b",
		}, {
			name:       "empty string",
			gotString:  "",
			wantString: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Slug(tt.gotString)
			assert.Equal(t, tt.wantString, got)
		})
	}
}

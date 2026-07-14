package normalize

import (
	"fmt"
	"testing"
)

func TestClean(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"  hello    friend ", "hello friend"},
		{"Are YoU want to PlAy", "are you want to play"},
		{"i think we can do it", "i think we can do it"},
		{"  Hello,   world! 123  ", "hello, world! 123"},
	}
	for _, c := range cases {
		// имя подтеста формируем из входных данных
		name := fmt.Sprintf("%s", c.in)

		t.Run(name, func(t *testing.T) {
			got := Clean(c.in)
			if got != c.out {
				t.Errorf("we expected: %s, but we want: %s", got, c.out)
			}
		})
	}
}

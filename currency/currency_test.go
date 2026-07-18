package currency

import (
	"fmt"
	"testing"
)

func TestPriceIn(t *testing.T) {
	tests := []struct {
		name string
		pay  float64
		from string
		to   string
	}{
		{name: "base_price", pay: 20.7, from: "Evro", to: "Dollar"},
		{name: "base_price_2", pay: 121.0, from: "Dollar", to: "Evro"},
		{name: "negative_price", pay: -35.3, from: "Tenge", to: "Evro"},
		{name: "nil_price", pay: 0, from: "Dollar", to: "Lari"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockConverter{}
			got := PriceIn(tt.pay, tt.from, tt.to, mock)
			want := 42.0

			fmt.Printf("сумма: %.2f\n", tt.pay)
			fmt.Printf("from: %s\n", tt.from)
			fmt.Printf("to: %s\n", tt.to)
			fmt.Print("\n")

			if mock.calls != 1 {
				t.Errorf("expected 1 call to Convert, got %d", mock.calls)
			}

			if mock.lastAmount != tt.pay {
				t.Errorf("expected lastAmount=%f, got %f", tt.pay, mock.lastAmount)
			}

			if mock.lastFrom != tt.from {
				t.Errorf("expected from: %s, got: %s", tt.from, mock.lastFrom)
			}

			if mock.lastTo != tt.to {
				t.Errorf("expected to: %s, got: %s", tt.to, mock.lastTo)
			}

			if got != want {
				t.Errorf("mock_PriceIn expected: %f, got: %f", want, got)
			}
		})
	}

}

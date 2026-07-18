package currency

// Converter отвечает за пересчёт сумм между валютами.
type Converter interface {
	Convert(amount float64, from, to string) float64
}

// PriceIn пересчитывает цену через переданный Converter.
func PriceIn(amount float64, from, to string, c Converter) float64 {
	return c.Convert(amount, from, to)
}

type mockConverter struct {
	lastAmount float64
	lastFrom   string
	lastTo     string
	calls      int
}

func (m *mockConverter) Convert(amount float64, from, to string) float64 {
	m.calls++
	m.lastAmount = amount
	m.lastFrom = from
	m.lastTo = to
	return 42.0
}

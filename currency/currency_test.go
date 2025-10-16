package currency

import "testing"

type mockConverter struct{
	lastAmount float64
	lastFrom, lastTo string
	calls int
}

func (m *mockConverter) Convert(amount float64, from,to string) float64{
	m.lastAmount, m.lastFrom, m.lastTo = amount, from, to
	m.calls++
	return 42.0
}

func TestPriceIn_DelegatesAndReturns(t *testing.T){
	m := &mockConverter{}
	got := PriceIn(10.5, "EUR", "USD", m)
	if got != 42.0{
		t.Fatalf("got %v, want %v", got, 42.0)
	}

	if m.calls != 1{
		t.Fatalf("calls: got %d, want 1", m.calls)
	}

	if m.lastAmount != 10.5 || m.lastFrom != "EUR" || m.lastTo != "USD"{
		t.Fatalf("args mismatch: (%v,%s->%s)", m.lastAmount, m.lastFrom, m.lastTo)
	}
}

func TestPriceIn_ZeroAndNegative(t *testing.T){
	m := &mockConverter{}
	_ = PriceIn(0, "RUB", "USD", m)

	if m.lastAmount != 0 || m.lastFrom != "RUB" || m.lastTo != "USD"{
		t.Fatalf("args mismatch for zero: (%v,%s->%s)", m.lastAmount, m.lastFrom, m.lastTo)
	}

	_ = PriceIn(-7, "EUR", "USD", m)

	if m.lastAmount != -7 || m.lastFrom != "EUR" || m.lastTo != "USD"{
		t.Fatalf("args mismatch for negative: (%v,%s->%s)", m.lastAmount, m.lastFrom, m.lastTo)
	}
}
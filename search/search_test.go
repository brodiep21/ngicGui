package search

import "testing"

func TestSearchTaxRate(t *testing.T) {
	t.Run("Greeley zip returns correct value", func(t *testing.T) {
		zip := "80634"
		want := "7.01%"
		got, err := SearchForTaxRate(zip)
		if err != nil {
			t.Error(err)
		}

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
	t.Run("LA zip returns correct value", func(t *testing.T) {
		zip := "70650"
		want := "9.45%"
		got, err := SearchForTaxRate(zip)
		if err != nil {
			t.Error(err)
		}

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	})
}

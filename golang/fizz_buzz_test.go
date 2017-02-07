package fizz_buzz

import "testing"

func TestToStart(t *testing.T) {
	if fizz() != 1 {
		t.Errorf("%s", "Errouuurrr Feio")
	}
}

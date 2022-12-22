package actions

import (
	"testing"
)

func TestPets(t *testing.T) {
	got := charToInt('r')
	want := 18

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

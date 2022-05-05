package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("want: %d, but got: %d", want, got)
	}
}

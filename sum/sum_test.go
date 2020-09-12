package sum

import "testing"

func TestSum(t *testing.T) {
	got := Sum(1, 2)
	if got != 3 {
		t.Errorf("Sum(1,2) == %d, want 3", got)
	}
}

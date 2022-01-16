package algorithms

import (
	"testing"
)

func TestUtopianTree(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 1
	got = utopianTree(0)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 2
	got = utopianTree(1)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 7
	got = utopianTree(4)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func utopianTree(n int32) int32 {
	var h int32 = 1

	for p, l := 0, int(n); p < l; p++ {
		if p%2 == 0 {
			h *= 2
		} else {
			h++
		}
	}

	return h
}

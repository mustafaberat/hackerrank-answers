package algorithms

import (
	"strconv"
	"testing"
)

func TestFairRations(t *testing.T) {
	var want, got string

	want = "4"
	got = fairRations([]int32{4, 5, 6, 7})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "4"
	got = fairRations([]int32{2, 3, 4, 5, 6})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "NO"
	got = fairRations([]int32{1, 2})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func fairRations(B []int32) string {
	var (
		c int
		l = len(B) - 1
	)

	for i := 0; i < l; i++ {
		if B[i]%2 == 1 {
			B[i]++
			B[i+1]++
			c += 2
		}
	}

	if B[l]%2 == 1 {
		return "NO"
	}

	return strconv.Itoa(c)
}

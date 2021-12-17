package algorithms

import (
	"testing"
)

func TestDivisibleSumPairs(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 3
	got = divisibleSumPairs(5, []int32{1, 2, 3, 4, 5, 6})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 5
	got = divisibleSumPairs(3, []int32{1, 3, 2, 6, 1, 2})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func divisibleSumPairs(k int32, ar []int32) int32 {
	var pairCount int32 = 0

	for i, l := 0, len(ar); i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if (ar[i]+ar[j])%k == 0 {
				pairCount++
			}
		}

	}

	return pairCount
}

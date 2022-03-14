package algorithms

import (
	"math"
	"testing"
)

func TestMinimumDistances(t *testing.T) {
	var want, got int32

	want = 2
	got = minimumDistances([]int32{3, 2, 1, 2, 3})
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 3
	got = minimumDistances([]int32{7, 1, 3, 4, 1, 7})
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func minimumDistances(a []int32) int32 {
	var (
		res  int32 = math.MaxInt32
		diff int32 = 0
		i    int32 = 0
		l          = int32(len(a))
	)

	for ; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] == a[j] {
				diff = j - i
				if res > diff {
					res = diff
				}
				break
			}
		}
	}

	if res != math.MaxInt32 {
		return res
	}

	return -1
}

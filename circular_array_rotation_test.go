package algorithms

import (
	"testing"
)

func TestCircularArrayRotation(t *testing.T) {
	var want, got []int32

	want = []int32{5, 3}
	got = circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2})
	if isDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	var res = make([]int32, len(queries))

	a = rightRotateKTimes(a, k)

	for i, e := range queries {
		res[i] = a[e]
	}

	return res
}

func rightRotateKTimes(a []int32, k int32) []int32 {
	var (
		i   int32
		l   = int32(len(a))
		arr = make([]int32, l)
	)

	for i = 0; i < l; i++ {
		arr[(i+k)%l] = a[i]
	}

	return arr
}

func isDifferent(got, want []int32) bool {
	if len(got) != len(want) {
		return true
	}

	for i, limit := 0, len(want); i < limit; i++ {
		if want[i] != got[i] {
			return true
		}
	}
	return false
}

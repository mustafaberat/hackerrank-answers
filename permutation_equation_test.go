package algorithms

import (
	"testing"
)

func TestPermutationEquation(t *testing.T) {
	var (
		want []int32
		got  []int32
	)

	want = []int32{4, 2, 5, 1, 3}
	got = permutationEquation([]int32{5, 2, 1, 3, 4})
	if isDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func permutationEquation(p []int32) []int32 {
	var (
		res []int32
		i   int32 = 0
		l         = int32(len(p))
	)

	for ; i < l; i++ {
		res = append(res, findArrIndexPlusOne(p, findArrIndexPlusOne(p, i+1)))
	}

	return res
}

func findArrIndexPlusOne(p []int32, num int32) int32 {
	for i := range p {
		if p[i] == num {
			return int32(i + 1)
		}
	}
	return -1
}

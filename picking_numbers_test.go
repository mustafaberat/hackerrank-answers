package algorithms

import (
	"testing"
)

func TestPickingNumbers(t *testing.T) {
	var (
		want  int32
		got   int32
		input []int32
	)

	want = 3
	input = []int32{4, 6, 5, 3, 3, 1}
	got = pickingNumbers(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}
	want = 5
	input = []int32{1, 2, 2, 3, 1, 2}
	got = pickingNumbers(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}
}

func pickingNumbers(a []int32) int32 {
	var (
		max = 1
		arr []int32
	)

	for i, l := 0, len(a); i < l; i++ {
		arr = make([]int32, 0)
		arr = append(arr, a[i])

		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			if isDiffRange(arr, a[j]) {
				arr = append(arr, a[j])
			}
		}
		if max < len(arr) {
			max = len(arr)
		}
	}

	return int32(max)
}

func isDiffRange(arr []int32, num int32) bool {
	diffMaxCount := 1

	for i := range arr {
		if abs(arr[i]-num) > diffMaxCount {
			return false
		}
	}

	return true
}

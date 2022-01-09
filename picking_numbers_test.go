package algorithms

import (
	"sort"
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
		res        int32 = 0
		maxDiffNum int32 = 1
	)

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	for i, l := 0, len(a); i < l; i++ {
		for j := i + 1; j < l; j++ {
			if a[j]-a[i] > maxDiffNum {
				break
			}
			res = max(res, int32(j-i+1))
		}
	}

	return res
}

func max(nums ...int32) int32 {
	max := nums[0]

	for i, l := 1, len(nums); i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

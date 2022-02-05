package algorithms

import (
	"testing"
)

func TestEqualizeArray(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 2
	got = equalizeArray([]int32{3, 3, 2, 1, 3})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func equalizeArray(arr []int32) int32 {
	repeatCount := getMaxRepeatedNumberCount(arr)
	return int32(len(arr)) - repeatCount
}

func getMaxRepeatedNumberCount(arr []int32) int32 {
	var m = make(map[int32]int32)
	for i, l := 0, len(arr); i < l; i++ {
		m[arr[i]]++
	}

	var maxRepeatedNumCount int32 = 0
	for _, value := range m {
		if value > maxRepeatedNumCount {
			maxRepeatedNumCount = value
		}
	}

	return maxRepeatedNumCount
}

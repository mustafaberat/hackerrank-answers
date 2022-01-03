package algorithms

import (
	"testing"
)

func TestCutTheSticks(t *testing.T) {
	var (
		want  []int32
		input []int32
		got   []int32
	)

	want = []int32{3, 2, 1}
	input = []int32{1, 2, 3}
	got = cutTheSticks(input)
	if isDifferent(got, want) {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", input, got, want)
	}

	want = []int32{6, 4, 2, 1}
	input = []int32{5, 4, 4, 2, 2, 8}
	got = cutTheSticks(input)
	if isDifferent(got, want) {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", input, got, want)
	}

	want = []int32{8, 6, 4, 1}
	input = []int32{1, 2, 3, 4, 3, 3, 2, 1}
	got = cutTheSticks(input)
	if isDifferent(got, want) {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", input, got, want)
	}
}

func cutTheSticks(arr []int32) []int32 {
	var cutLen []int32

	for true {
		minStick := findMinNumberIntoArray(arr)
		cutLen = append(cutLen, int32(len(arr)))
		arr = cutSticks(arr, minStick)
		if len(arr) == 0 {
			break
		}
	}

	return cutLen
}

func cutSticks(arr []int32, minStickLen int32) []int32 {
	var newArr []int32

	for i := range arr {
		if arr[i]-minStickLen > 0 {
			newArr = append(newArr, arr[i]-minStickLen)
		}
	}

	return newArr
}

func findMinNumberIntoArray(arr []int32) int32 {
	min := arr[0]

	for i, l := 1, len(arr); i < l; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min
}

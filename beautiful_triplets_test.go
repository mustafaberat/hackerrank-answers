package algorithms

import "testing"

func TestBeautifulTriplets(t *testing.T) {
	var want, got int32

	want = 3
	got = beautifulTriplets(3, []int32{1, 2, 4, 5, 7, 8, 10})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func beautifulTriplets(d int32, arr []int32) int32 {
	var c int32 = 0

	for i, l := 0, len(arr); i < l; i++ {
		if isInArr(arr[i]+d, arr) && isInArr(arr[i]+d+d, arr) {
			c++
		}
	}

	return c
}

func isInArr(num int32, arr []int32) bool {
	var found = false

	for i := range arr {
		if arr[i] == num {
			found = true
			break
		}
	}

	return found
}

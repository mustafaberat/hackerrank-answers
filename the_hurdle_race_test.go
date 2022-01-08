package algorithms

import "testing"

func TestHurdleRace(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 5
	got = hurdleRace(0, []int32{4, 2})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func hurdleRace(k int32, height []int32) int32 {
	max := findMaxNumberInArray(height)

	if k <= max {
		return max - k
	}
	return 0
}

func findMaxNumberInArray(arr []int32) int32 {
	var max int32 = 0

	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}

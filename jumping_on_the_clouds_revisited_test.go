package algorithms

import (
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	var (
		got  int32
		want int32
	)

	want = 92
	got = jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 80
	got = jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func jumpingOnClouds(c []int32, k int32) int32 {
	var index int32 = 0
	var remainEnergy int32 = 100
	var l = int32(len(c))

	for true {
		index = (index + k) % l
		remainEnergy -= c[index]*2 + 1

		if index == 0 {
			break
		}
	}

	return remainEnergy
}

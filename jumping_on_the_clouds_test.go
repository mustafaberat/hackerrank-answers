package algorithms

import (
	"testing"
)

func TestJumpingOnCloud(t *testing.T) {
	var want, got int32

	want = 4
	got = jumpingOnCloud([]int32{0, 0, 1, 0, 0, 1, 0})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 3
	got = jumpingOnCloud([]int32{0, 0, 0, 0, 1, 0})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 3
	got = jumpingOnCloud([]int32{0, 0, 0, 1, 0, 0})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func jumpingOnCloud(c []int32) int32 {
	var (
		res int32 = 0
		l         = len(c)
	)

	for i := 0; i < l; i++ {
		if i+2 < l && c[i+2] == 0 {
			i++
		}
		res++
	}

	return res - 1
}

package algorithms

import (
	"testing"
)

func TestBirthday(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 2
	got = birthday([]int32{2, 2, 1, 3, 2}, 4, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 2
	got = birthday([]int32{1, 2, 1, 3, 2}, 3, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 0
	got = birthday([]int32{1, 1, 1, 1, 1, 1}, 3, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 1
	got = birthday([]int32{4}, 4, 1)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func birthday(s []int32, d int32, m int32) int32 {
	var (
		res int32 = 0
		i   int32 = 0
		lim       = int32(len(s))
	)

	for ; i < lim; i++ {
		t := i + m
		if t >= lim {
			t = lim
		}

		if sum := sumSlice(s[i:t]); sum == d {
			res++
		}
	}

	return res
}

func sumSlice(s []int32) int32 {
	var t int32 = 0
	for _, num := range s {
		t += num
	}
	return t
}

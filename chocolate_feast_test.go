package algorithms

import (
	"testing"
)

func TestChocolateFeast(t *testing.T) {
	var want, got int32

	want = 9
	got = chocolateFeast(15, 3, 2)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 6
	got = chocolateFeast(10, 2, 5)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func chocolateFeast(n int32, c int32, m int32) int32 {
	var (
		bar     = n / c
		wrapper = bar
		sum     = bar
	)

	for bar != 0 {
		bar = wrapper / m
		sum += bar
		wrapper = (wrapper % m) + bar
	}

	return sum
}

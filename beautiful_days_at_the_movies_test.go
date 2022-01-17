package algorithms

import (
	"testing"
)

func TestBeautifulDays(t *testing.T) {

	var (
		want int32
		got  int32
	)

	want = 2
	got = beautifulDays(20, 23, 6)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func beautifulDays(i int32, j int32, k int32) int32 {
	var beautifulDayCounter int32 = 0

	for t := i; t <= j; t++ {
		if (t-makeReverse(t))%k == 0 {
			beautifulDayCounter++
		}
	}

	return beautifulDayCounter
}

func makeReverse(n int32) int32 {
	var newInt int32 = 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}

package algorithms

import (
	"strconv"
	"testing"
)

func TestKaprekarNumbers(t *testing.T) {
	var want, got []int32

	want = []int32{297}
	got = kaprekarNumbers(100, 300)
	if isDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = []int32{1, 9, 45, 55, 99}
	got = kaprekarNumbers(1, 100)
	if isDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = []int32{1, 9, 45, 55, 99, 297, 703, 999, 2223, 2728, 4950, 5050, 7272, 7777, 9999, 17344, 22222, 77778, 82656, 95121, 99999}
	got = kaprekarNumbers(1, 99999)
	if isDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func kaprekarNumbers(p int32, q int32) []int32 {
	var res []int32

	for i := p; i <= q; i++ {
		if isKaprekarNumber(int(i)) {
			res = append(res, i)
		}
	}

	if len(res) > 0 {
		return res
	}
	return []int32{-1}
}

func isKaprekarNumber(n int) bool {
	nStr := strconv.Itoa(n * n)
	mid := len(nStr) / 2
	left, _ := strconv.Atoi(nStr[:mid])
	right, _ := strconv.Atoi(nStr[mid:])

	if left+right == n {
		return true
	}

	return false
}

func isDifferent(got, want []int32) bool {
	if len(got) != len(want) {
		return true
	}

	for i, limit := 0, len(want); i < limit; i++ {
		if want[i] != got[i] {
			return true
		}
	}
	return false
}

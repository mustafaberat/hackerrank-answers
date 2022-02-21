package algorithms

import "testing"

func TestTaumBday(t *testing.T) {
	var want, got int64

	want = 29
	got = taumBday(3, 5, 3, 4, 1)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 12
	got = taumBday(3, 6, 9, 1, 1)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 617318315833461267
	got = taumBday(742407782, 90529439, 847666641, 8651519, 821801924)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func taumBday(b int64, w int64, bc int64, wc int64, z int64) int64 {
	total := b * findMin(z+wc, bc)
	total += w * findMin(z+bc, wc)

	return total
}

func findMin(nums ...int64) int64 {
	min := nums[0]

	for i, l := 1, len(nums); i < l; i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}
taum_and_bday_test.go

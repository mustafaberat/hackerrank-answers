package algorithms

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	var (
		input       = []int32{18, 90, 90, 13, 90, 75, 90, 8, 90, 43}
		want  int32 = 5
	)
	got := birthdayCakeCandles(input)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func birthdayCakeCandles(candles []int32) int32 {
	var (
		max      int32 = 0
		countMax int32 = 0
	)

	for i, lim := 0, len(candles); i < lim; i++ {
		if candles[i] > max {
			max = candles[i]
			countMax = 0
		}
		if candles[i] == max {
			countMax++
		}
	}
	return countMax
}

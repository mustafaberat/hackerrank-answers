package algorithms

import (
	"strconv"
	"strings"
	"testing"
)

func TestFindDigits(t *testing.T) {
	var (
		want  int32
		input int32
		got   int32
	)

	input = 1012
	want = 3
	got = findDigits(input)
	if want != got {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", input, got, want)
	}

	input = 12
	want = 2
	got = findDigits(input)
	if want != got {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", input, got, want)
	}
}

func findDigits(n int32) int32 {
	var (
		divisorCount int32 = 0
		nAsString          = strconv.Itoa(int(n))
		numbers            = strings.Split(nAsString, "")
	)

	for i, l := 0, len(numbers); i < l; i++ {
		numAsInt, err := strconv.Atoi(numbers[i])
		if err != nil {
			return -1
		}

		numAsInt32 := int32(numAsInt)
		if numAsInt32 != 0 && n%numAsInt32 == 0 {
			divisorCount++
		}
	}

	return divisorCount
}

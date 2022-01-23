package algorithms

import (
	"testing"
)

func TestViralAdvertising(t *testing.T) {
	var (
		want  int32
		input int32
		got   int32
	)

	want = 24
	input = 5
	got = viralAdvertising(input)
	if got != want {
		t.Errorf("Input: %v Got: %v Want: %v \n", input, got, want)
	}

	want = 9
	input = 3
	got = viralAdvertising(input)
	if got != want {
		t.Errorf("Input: %v Got: %v Want: %v \n", input, got, want)
	}
}

func viralAdvertising(n int32) int32 {
	var (
		shared     int32 = 5
		liked      int32 = 0
		cumulative int32 = 0
	)

	for i, l := 0, int(n); i < l; i++ {
		liked = shared / 2
		cumulative += liked
		shared = 3 * liked
	}

	return cumulative
}

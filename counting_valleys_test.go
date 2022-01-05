package algorithms

import (
	"testing"
)

func TestCountingValleys(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 1
	got = countingValleys("UDDDUDUU")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 0
	got = countingValleys("UDUUUDUDDD")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 2
	got = countingValleys("DDUUDDDUDUUDUDDDUUDDUDDDUDDDUDUUDDUUDDDUDDDUDDDUUUDUDDDUDUDUDUUDDUDUDUDUDUUUUDDUDDUUDUUDUUDUUUUUUUUU")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func countingValleys(path string) int32 {
	var lvl, seaLvlCounter int32

	for i, l := 0, len(path); i < l; i++ {
		if path[i] == 'D' {
			lvl--
		} else {
			lvl++
		}

		//We came up to sea lvl
		if lvl == 0 && path[i] == 'U' {
			seaLvlCounter++
		}
	}

	return seaLvlCounter
}

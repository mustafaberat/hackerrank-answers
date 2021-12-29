package algorithms

import (
	"testing"
)

func TestMigratoryBirds(t *testing.T) {
	var (
		want  int32
		got   int32
		input []int32
	)

	want = 1
	input = []int32{1, 1, 2, 2, 3}
	got = migratoryBirds(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}

	want = 3
	input = []int32{2, 4, 3, 2, 3, 1, 2, 1, 3, 3}
	got = migratoryBirds(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}

	want = 4
	input = []int32{1, 4, 4, 4, 5, 3}
	got = migratoryBirds(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}

	want = 3
	input = []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	got = migratoryBirds(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}
}

func migratoryBirds(arr []int32) int32 {
	frequencyNums := make([]int32, 5)

	for i, l := 0, len(arr); i < l; i++ {
		frequencyNums[arr[i]-1] += 1
	}

	max, occ := frequencyNums[0], 1

	for i, num := range frequencyNums {
		if num > max {
			max = num
			occ = i + 1
		}
	}

	return int32(occ)
}

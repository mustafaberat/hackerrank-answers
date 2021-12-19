package algorithms

import "testing"

func TestBreakingRecords(t *testing.T) {
	var (
		want []int32
		got  []int32
	)

	want = []int32{2, 4}
	got = breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
	if isDifferent(want, got) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func breakingRecords(scores []int32) []int32 {
	var max, min, minumum, maximum int32 = 0, 0, scores[0], scores[0]

	for i, l := 1, len(scores); i < l; i++ {
		if scores[i] > maximum {
			maximum = scores[i]
			max++
		} else if scores[i] < minumum {
			minumum = scores[i]
			min++
		}
	}
	return []int32{max, min}
}

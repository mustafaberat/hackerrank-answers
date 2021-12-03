package algorithms

import "testing"

func TestGrading(t *testing.T) {
	inputs := []int32{73, 67, 38, 33}
	want := []int32{75, 67, 40, 33}
	got := grading(inputs)

	if isDifferent(got, want) {
		t.Errorf("Result is not correct. Got: %v Want: %v \n", got, want)
	}
}

func grading(inputs []int32) []int32 {
	var (
		results []int32
		_input  int32
	)

	for _, input := range inputs {
		_input = input
		var gradingLimit int32 = 3
		for true {
			if _input < 40-gradingLimit {
				results = append(results, _input)
				break
			} else if _input%5 == 0 {
				results = append(results, _input)
				break
			} else if gradingLimit > 1 {
				_input++
				gradingLimit--
			} else {
				results = append(results, input)
				break
			}
		}
	}
	return results
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

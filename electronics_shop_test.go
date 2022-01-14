package algorithms

import "testing"

func TestGetMoneySpent(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 9
	got = getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = -1
	got = getMoneySpent([]int32{4}, []int32{5}, 5)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var result int32 = -1

	for i := range keyboards {
		for j := range drives {
			if keyboards[i]+drives[j] > result && keyboards[i]+drives[j] <= b {
				result = keyboards[i] + drives[j]
			}
		}
	}

	return result
}

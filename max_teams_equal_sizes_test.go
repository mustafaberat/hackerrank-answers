package algorithms

import (
	"slices"
	"sort"
	"testing"
)

func TestMaxTeamsEqualSizes(t *testing.T) {
	var (
		want     int32
		got      int32
		teamSize []int32
		k        int32
	)

	teamSize = []int32{1, 2, 2, 3, 4}
	k = 2
	want = 4
	got = maxTeamsEqualSize(teamSize, k)

	if want != got {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", teamSize, got, want)
	}

	teamSize = []int32{1, 7, 3, 8}
	k = 1
	want = 2
	got = maxTeamsEqualSize(teamSize, k)

	if want != got {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", teamSize, got, want)
	}

	teamSize = []int32{1, 2, 3, 2, 2, 2}
	k = 1
	want = 5
	got = maxTeamsEqualSize(teamSize, k)

	if want != got {
		t.Errorf("Not correct. Input: %v Got: %v Want: %v \n", teamSize, got, want)
	}
}

func maxTeamsEqualSize(teamSize []int32, k int32) int32 {
	sort.Slice(teamSize, func(i, j int) bool { return teamSize[i] < teamSize[j] })
	reduceTime := k

	var results []int32

	teams := make([]int32, len(teamSize))
	copy(teams, teamSize)

	for i := 0; i < len(teams)-1; i++ {
		reduceTime = k
		copy(teams, teamSize)

		for j := i + 1; j < len(teams); j++ {
			if teams[j]-teams[i] != 0 && teams[j]-teams[i] <= k {
				teams[j] = teams[i]
				reduceTime -= 1

				if reduceTime == 0 {
					results = append(results, findSumOfDepracatedNumbers(teams))
					break
				}
			}

		}
	}

	return slices.Max(results)
}

func findSumOfDepracatedNumbers(teams []int32) int32 {
	sumDepracated := 1
	for i := 1; i < len(teams); i++ {
		if teams[i] == teams[i-1] {
			sumDepracated += 1
		}
	}

	return int32(sumDepracated)
}

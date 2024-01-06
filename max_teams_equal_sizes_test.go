// 1. Hackathon at HackerRank
// HackerRank is organizing a hackathon for all its employees.
// A hackathon is a team event, and there are n teams taking part. The number of employees in the ith team is denoted by teamSize[i]. In order to maintain uniformity, the team size of at most k teams can be
// reduced. Find the maximum number of teams of equal size that can be formed if team size is reduced optimally.
// Example
// There are n = 5 teams, team sizes are teamSize = [1, 2, 2, 3, 4], and the maximum number of teams whose size can be reduced, k = 2.
// The team size of the last 2 teams can be reduced to 2, thus teamSize = [1, 2, 2, 2, 2]. The maximum number of teams with equal size is 4.
// Function Description
// Complete the function equalizeTeamSize in the editor below.
// equalize TeamSize has the following parameters:
// int teamSize[n]: the number of employees in each team
// int k: the maximum number of teams whose size can be reduced
// Returns
// int: the maximum number of equal size teams possible

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

package algorithms

import (
	"testing"
)

func TestAcmTeam(t *testing.T) {
	var want, got []int32

	want = []int32{5, 1}
	got = acmTeam([]string{"10101", "11110", "00010"})
	if isDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = []int32{5, 2}
	got = acmTeam([]string{"10101", "11100", "11010", "00101"})
	if isDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func acmTeam(topic []string) []int32 {
	var (
		maxTopicCount      int32 = 0
		maxTeamCount int32 = 0
		andOpRes     int32
		l            = len(topic)
	)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			andOpRes = andOperation(topic[i], topic[j])
			if andOpRes > maxTopicCount {
				maxTopicCount = andOpRes
				maxTeamCount = 1
			} else if andOpRes == maxTopicCount {
				maxTeamCount++
			}
		}
	}

	return []int32{maxTopicCount, maxTeamCount}
}

func andOperation(s1, s2 string) int32 {
	var (
		l       = len(s1)
		c int32 = 0
	)

	for k := 0; k < l; k++ {
		if s1[k] == '1' || s2[k] == '1' {
			c++
		}
	}

	return c
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

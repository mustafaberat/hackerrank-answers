package algorithms

import "testing"

func TestAngryProfessor(t *testing.T) {
	var (
		want string
		got  string
	)

	want = "NO"
	got = angryProfessor(2, []int32{0, -1, 2, 1})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func angryProfessor(k int32, a []int32) string {
	var l int32 = 0

	for i := range a {
		if a[i] <= 0 {
			l++
		}
	}

	if l >= k {
		return "NO"
	}
	return "YES"
}

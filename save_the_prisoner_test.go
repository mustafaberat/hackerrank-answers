package algorithms

import "testing"

func TestSaveThePrisoner(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 2
	got = saveThePrisoner(5, 2, 1)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 3
	got = saveThePrisoner(5, 2, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 3
	got = saveThePrisoner(3, 7, 3)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 6
	got = saveThePrisoner(7, 19, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	return ((s - 1 + m - 1) % n) + 1
}

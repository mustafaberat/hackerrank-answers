package algorithms

import "testing"

func TestPageCount(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 0
	got = pageCount(5, 4)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 1
	got = pageCount(6, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 1
	got = pageCount(5, 3)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 1
	got = pageCount(6, 5)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func pageCount(n, p int32) int32 {
	front := p / 2
	back := (n - p) / 2

	if n%2 == 0 {
		back = (n - p + 1) / 2
	}

	if front < back {
		return front
	}
	return back
}

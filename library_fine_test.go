package algorithms

import "testing"

func TestLibraryFine(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 0
	got = libraryFine(2, 7, 1014, 1, 1, 1015)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 45
	got = libraryFine(9, 6, 2015, 6, 6, 2015)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1 > y2 {
		return 10000
	} else if y1 == y2 {
		if m1 > m2 {
			return 500 * (m1 - m2)
		} else if m1 == m2 {
			if d1 > d2 {
				return 15 * (d1 - d2)
			}
		}
	}
	return 0
}

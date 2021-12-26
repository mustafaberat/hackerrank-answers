package algorithms

import "testing"

func TestCatAndMouse(t *testing.T) {
	var (
		want string
		got  string
	)

	want = "Mouse C"
	got = catAndMouse(1, 3, 2)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "Cat B"
	got = catAndMouse(1, 2, 3)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func catAndMouse(x, y, z int32) string {
	var (
		catA = abs(x - z)
		catB = abs(y - z)
	)

	if catA > catB {
		return "Cat B"
	} else if abs(x-z) == abs(y-z) {
		return "Mouse C"
	}
	return "Cat A"
}


func abs(num int32) int {
	if num < 0 {
		return int(-num)
	}
	return int(num)
}

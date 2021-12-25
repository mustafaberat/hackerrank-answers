package algorithms

import (
	"strconv"
	"testing"
)

func TestBonAppetit(t *testing.T) {
	var (
		want string
		got  string
	)

	want = "5"
	got = bonAppetit([]int32{3, 10, 2, 9}, 1, 12)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "Bon Appetit"
	got = bonAppetit([]int32{3, 10, 2, 9}, 1, 7)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func bonAppetit(bill []int32, k int32, b int32) string {
	var sum int32

	for i := range bill {
		if int32(i) != k {
			sum += bill[i]
		}
	}

	bActual := sum / 2
	result := b - bActual

	if result == 0 {
		return "Bon Appetit"
	}
	return strconv.Itoa(int(result))
}

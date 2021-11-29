package algorithms

import "testing"

func TestCountAp(t *testing.T) {
	appleLen, orangeLen := countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
	if appleLen != 1 || orangeLen != 1 {
		t.Errorf("Not correct. %v %v \n", appleLen, orangeLen)
	}
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (int32, int32) {
	var (
		appleLen  int32 = 0
		orangeLen int32 = 0
	)

	for _, apple := range apples {
		i := apple + a
		if isBetweenHouse(i, s, t) {
			appleLen++
		}
	}

	for _, orange := range oranges {
		i := orange + b
		if isBetweenHouse(i, s, t) {
			orangeLen++
		}
	}

	return appleLen, orangeLen
}

func isBetweenHouse(num, s, t int32) bool {
	if s > num || t < num {
		return false
	}
	return true
}

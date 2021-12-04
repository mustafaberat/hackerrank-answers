package algorithms

/*
you are choreographing a circus show with various animals. For one act, you are given two kangaroos on a number line ready to jump in the positive direction (i.e, toward positive infinity).

The first kangaroo starts at location  and moves at a rate of  meters per jump.
The second kangaroo starts at location  and moves at a rate of  meters per jump.
You have to figure out a way to get both kangaroos at the same location at the same time as part of the show. If it is possible, return YES, otherwise return NO.
Function Description

Complete the function kangaroo in the editor below.

kangaroo has the following parameter(s):

int x1, int v1: starting position and jump distance for kangaroo 1
int x2, int v2: starting position and jump distance for kangaroo 2
Returns

string: either YES or NO
*/
import (
	"testing"
)

func TestKangaroo(t *testing.T) {
	want := "YES"
	got := kangaroo(0, 3, 4, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "NO"
	got = kangaroo(0, 2, 5, 3)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "NO"
	got = kangaroo(21, 6, 47, 3)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "NO"
	got = kangaroo(43, 2, 70, 2)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func kangaroo(x1, v1, x2, v2 int32) string {
	if isForwardAndFaster(x1, v1, x2, v2) {
		return "NO"
	} else {
		if v1-v2 == 0 {
			return "NO"
		} else if (x2-x1)%(v1-v2) == 0 {
			return "YES"
		}
	}
	return "NO"
}

func isForwardAndFaster(x1, v1, x2, v2 int32) bool {
	if v1 > v2 && x1 > x2 || v1 < v2 && x1 < x2 {
		return true
	}
	return false
}

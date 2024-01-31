//You are given a string num representing a large integer. An integer is good if it meets the following conditions:
// It is a substring of num with length 3.
// It consists of only one unique digit.
// Return the maximum good integer as a string or an empty string "" if no such integer exists.

// Note:
// A substring is a contiguous sequence of characters within a string.
// There may be leading zeroes in num or a good integer.

// Example 1:
// Input: num = "6777133339"
// Output: "777"
// Explanation: There are two distinct good integers: "777" and "333".
// "777" is the largest, so we return "777".

// Example 2:
// Input: num = "2300019"
// Output: "000"
// Explanation: "000" is the only good integer.

// Example 3:
// Input: num = "42352338"
// Output: ""
// Explanation: No substring of length 3 consists of only one unique digit. Therefore, there are no good integers.

package algorithms

import (
	"strconv"
	"testing"
)

func TestLargestGoodInteger(t *testing.T) {
	var want, got string

	want = "777"
	got = largestGoodInteger("6777133339")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "000"
	got = largestGoodInteger("2300019")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = ""
	got = largestGoodInteger("42352338")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}
func largestGoodInteger(num string) string {
	var sames []string

	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			sames = append(sames, num[i:i+3])
		}
	}

	if len(sames) == 0 {
		return ""
	}

	return largestGoodIntegerFindBigger(sames)
}

func largestGoodIntegerFindBigger(nums []string) string {
	max := 0

	for j := 0; j < len(nums); j++ {
		num, _ := strconv.Atoi(nums[j])
		if max < num {
			max = num
		}
	}

	if max == 0 {
		return "000"
	}

	return strconv.Itoa(max)
}

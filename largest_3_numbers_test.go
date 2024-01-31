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
	"strings"
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
	biggest := "0"
	found := false

	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			found = true
			biggest = maxStr(biggest, string(num[i]))
		}
	}

	if found == false {
		return ""
	}

	return strings.Repeat(biggest, 3)
}

func maxStr(num1, num2 string) string {
	num1Int, _ := strconv.Atoi(num1)
	num2Int, _ := strconv.Atoi(num2)

	if num1Int > num2Int {
		return num1
	}
	return num2
}

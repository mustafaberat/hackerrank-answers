// 1758. Minimum Changes To Make Alternating Binary String
// You are given a string s consisting only of the characters '0' and '1'. In one operation, you can change any '0' to '1' or vice versa.

// The string is called alternating if no two adjacent characters are equal. For example, the string "010" is alternating, while the string "0100" is not.

// Return the minimum number of operations needed to make s alternating.

// Example 1:

// Input: s = "0100"
// Output: 1
// Explanation: If you change the last character to '1', s will be "0101", which is alternating.
// Example 2:

// Input: s = "10"
// Output: 0
// Explanation: s is already alternating.
// Example 3:

// Input: s = "1111"
// Output: 2
// Explanation: You need two operations to reach "0101" or "1010".

// Constraints:

// 1 <= s.length <= 104
// s[i] is either '0' or '1'.
package algorithms

import (
	"math"
	"testing"
)

func TestMinOperations(t *testing.T) {
	var (
		want  int
		got   int
		input string
	)

	want = 1
	input = "0100"
	got = minOperations(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}

	want = 0
	input = "10"
	got = minOperations(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}

	want = 2
	input = "1111"
	got = minOperations(input)
	if got != want {
		t.Errorf("Not correct. Input: %v \t Got: %v Want: %v \n", input, got, want)
	}
}

func minOperations(s string) int {
	zeroChangeCount := 0
	l := len(s)

	for i := 0; i < l; i += 1 {
		if i%2 == 0 {
			if s[i] == '1' {
				zeroChangeCount += 1
			}
		} else {
			if s[i] == '0' {
				zeroChangeCount += 1
			}
		}
	}

	return int(math.Min(float64(zeroChangeCount), float64(l-zeroChangeCount)))
}

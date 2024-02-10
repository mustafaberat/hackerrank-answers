// Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. If there is no such substring return -1.
// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "aa"
// Output: 0
// Explanation: The optimal substring here is an empty substring between the two 'a's.

// Example 2:
// Input: s = "abca"
// Output: 2
// Explanation: The optimal substring here is "bc".

// Example 3:
// Input: s = "cbzxy"
// Output: -1
// Explanation: There are no characters that appear twice in s.
package algorithms

import (
	"testing"
)

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	var want, got int

	want = 0
	got = maxLengthBetweenEqualCharacters("aa")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 2
	got = maxLengthBetweenEqualCharacters("abca")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = -1
	got = maxLengthBetweenEqualCharacters("abcd")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func maxLengthBetweenEqualCharacters(s string) int {
	dict := make(map[byte]int)
	result := -1
	dict[s[0]] = 0

	for i := 1; i < len(s); i++ {
		if index, ok := dict[s[i]]; !ok {
			dict[s[i]] = i
		} else {
			distance := i - index - 1
			if distance > result {
				result = distance
			}
		}

	}

	return result
}

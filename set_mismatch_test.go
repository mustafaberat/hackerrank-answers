// You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.
// You are given an integer array nums representing the data status of this set after the error.
// Find the number that occurs twice and the number that is missing and return them in the form of an array.

// Example 1:
// Input: nums = [1,2,2,4]
// Output: [2,3]

// Example 2:
// Input: nums = [1,1]
// Output: [1,2]

package algorithms

import (
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	var want, got []int

	want = []int{2, 3}
	got = findErrorNums([]int{1, 2, 2, 4})
	if findErrorNumsIsDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = []int{1, 2}
	got = findErrorNums([]int{1, 1})
	if findErrorNumsIsDifferent(got, want) {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func findErrorNums(nums []int) []int {
	n := len(nums)
	counts := make([]int, n+1)
	duplicate, missing := 0, 0

	for _, num := range nums {
		counts[num]++
	}

	for i := 1; i <= n; i++ {
		if counts[i] == 2 {
			duplicate = i
		} else if counts[i] == 0 {
			missing = i
		}
	}

	return []int{duplicate, missing}
}

func findErrorNumsIsDifferent(got, want []int) bool {
	if len(got) != len(want) {
		return true
	}

	for i, limit := 0, len(want); i < limit; i++ {
		if want[i] != got[i] {
			return true
		}
	}
	return false
}

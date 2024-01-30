// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray
// whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

// Example 1:
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Example 2:
// Input: target = 4, nums = [1,4,4]
// Output: 1

// Example 3:
// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

package algorithms

import (
	"math"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var want, got int

	want = 2
	got = minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 1
	got = minSubArrayLen(4, []int{1, 4, 4})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 0
	got = minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt

	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			res = min(res, right-left+1)

			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}

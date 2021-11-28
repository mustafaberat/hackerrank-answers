package algorithms

import (
	"testing"
)

//Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

func TestMiniMaxSum(t *testing.T) {
	inputs := [][]int32{
		{1, 3, 5, 7, 9},
		{1, 2, 3, 4, 5},
		{769082435, 210437958, 673982045, 375809214, 380564127},
		{256741038, 623958417, 467905213, 714532089, 938071625},
	}

	outputs := [][]int64{
		{16, 24},
		{10, 14},
		{1640793344, 2199437821},
		{2063136757, 2744467344},
	}

	for i := 0; i < len(inputs); i++ {
		gotMinSum, gotMaxSum := miniMaxSum(inputs[i])
		wantMinSum, wantMaxSum := outputs[i][0], outputs[i][1]
		if gotMinSum != wantMinSum || gotMaxSum != wantMaxSum {
			t.Errorf("Not correct. Got: %v %v\t Want: %v %v\n", gotMinSum, gotMaxSum, wantMinSum, wantMaxSum)
		}
	}

}

func miniMaxSum(arr []int32) (int64, int64) {
	var (
		min       = int64(arr[0])
		max       = int64(arr[1])
		sum int64 = 0
	)

	for _, num := range arr {
		number := int64(num)
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
		sum += number
	}

	return sum - max, sum - min

}

package algorithms

import (
	"testing"
)

func TestDiagonalCalculator(t *testing.T) {
	input1 := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	result1 := 2

	if result1 != DiagonalCalculator(input1) {
		t.Error("Not correct: \t", result1)
	}

	input2 := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	result2 := 15

	if result2 != DiagonalCalculator(input2) {
		t.Error("Not correct: \t", result2)
	}

	input3 := [][]int32{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	result3 := 0

	if result3 != DiagonalCalculator(input3) {
		t.Error("Not correct: \t", result3)
	}
}

func DiagonalCalculator(matrix [][]int32) int {
	var (
		leftToRight, rightToLeft int32 = 0, 0
		length                         = len(matrix) - 1
	)
	for i := range matrix {
		leftToRight += matrix[i][i]
	}
	for i := range matrix {
		rightToLeft += matrix[i][length-i]
	}
	return abs(rightToLeft - leftToRight)
}

func abs(num int32) int {
	if num < 0 {
		return int(-num)
	}
	return int(num)
}

// User
// There is a puzzle using a rectangular grid. The upper left corner is at (row, column) = (0, 0). Each cell contains an integer. The score starts at O and is the sum of all the integers in each cell visited as the grid is traversed. Movement begins in either the top or the bottom row and stays within the bounds of the grid. Only 1 cell can be visited per row per direction. Determine the maximum achievable score.
// Movement for the two scenarios are as follows:
// 1. From a cell (i,j) = (0,p), i.e. in the top row:
// 。 (i+1, j-1)
// 。 (i+1, j)
// 。 (i+1, j+1)
// 2. From a cell (i,j) = (rows-1,q), i.e. in the bottom row:
// 。 (i-1,j-1)
// 。 (i-1, j)
// 。 (i-1, j+1)
// Example// board = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
// Function Description
// Complete the function maxPathSum in the editor below.
// maxPathSum has the following parameter(s):
// int board[n][m]: the values for the grid cells
// p: row 0 starting column
// q: row n-1 starting column
// Returns:
// int: the maximum achievable score from the two start positions"

package algorithms

import (
	"fmt"
	"testing"
)

func TestGridTraversal(t *testing.T) {
	var (
		want  int32
		got   int32
		p     int32
		q     int32
		board [][]int32
	)

	board = [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	p = 1
	q = 0
	want = 17
	got = maxPathSum(board, p, q)
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func maxPathSum(board [][]int32, p, q int32) int32 {
	var (
		val1 int32
		val2 int32
	)

	val1 = findMaxPath(board, 0, p, true)
	val2 = findMaxPath(board, int32(len(board)-1), q, false)

	return max(val1, val2)
}

func findMaxPath(board [][]int32, r, c int32, goDown bool) int32 {
	var sum int32 = 0

	for r >= 0 && r < int32(len(board)) && c >= 0 && c <= int32(len(board)) {
		if goDown {
			sum += board[r][c]
			r += 1

			if r < int32(len(board)) && c > 0 && board[r][c-1] > board[r][c] {
				c -= 1
			} else if r < int32(len(board)) && c < int32(len(board)-1) && board[r][c+1] > board[r][c] {
				c += 1
			}
		} else {
			sum += board[r][c]
			r -= 1

			fmt.Println("Sum:", sum, "r:", r, "c:", c)

			if r >= 0 && c < int32(len(board)) && board[r][c+1] > board[r][c] {
				c += 1
			} else if r >= 0 && c > 0 && board[r][c-1] > board[r][c] {
				c -= 1
			}
		}
	}

	return sum
}

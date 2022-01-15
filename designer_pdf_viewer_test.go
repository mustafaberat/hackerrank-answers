package algorithms

import (
	"strings"
	"testing"
)

func TestDesignerPdfViewer(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 28
	got = designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "zaba")
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func designerPdfViewer(h []int32, word string) int32 {
	var (
		ioa  int32
		high int32 = 0
	)

	for i, l := 0, len(word); i < l; i++ {
		ioa = indexOfAlphabet(strings.ToUpper(string(word[i])))
		high = max(high, h[ioa])
	}

	return high * int32(len(word))
}

func indexOfAlphabet(character string) int32 {
	var abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for j := 0; j < len(abc); j++ {
		if character == string(abc[j]) {
			return int32(j)
		}
	}

	return -1
}

func max(nums ...int32) int32 {
	max := nums[0]

	for i, l := 1, len(nums); i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

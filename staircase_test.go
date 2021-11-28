package algorithms

import "fmt"

func StairCase(n int32) {
	var (
		char  = "#"
		space int32
		ch    int32
		i     int32
	)

	for i = 0; i < n; i++ {
		for space = 0; space < n-i-1; space++ {
			fmt.Print(" ")
		}
		for ch = 0; ch <= i; ch++ {
			fmt.Print(char)
		}
		fmt.Println()
	}
}

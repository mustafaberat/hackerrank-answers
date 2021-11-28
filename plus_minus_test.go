
func plusMinus(arr []int32) {
    var (
        p float32=0
        n float32=0
        z float32=0
        l = len(arr)
    )
    
    for i:=0; i<l;i++{
        if arr[i] > 0{
            p++
        } else if arr[i] < 0{
            n++
        } else{
            z++
        }
    }
    fmt.Println(p/float32(l))
    fmt.Println(n/float32(l))
    fmt.Println(z/float32(l))
}

// ______________________
// ______________________
// ______________________
// ______________________
// ______________________
// ______________________

package algorithms

import (
	"testing"
)

func Test_PlustMinus(t *testing.T) {
	inputs := [][]int{
		{1, 1, 0, -1, -1},
		{-4, 3, -9, 0, 4, 1},
	}
	results := [][]float32{
		{0.40, 0.40, 0.20},
		{0.50, 0.33, 0.16},
	}
	for i := 0; i < len(inputs); i++ {
		pl, nl, zl := plusMinus(inputs[i])
		if !(pl == results[i][0] || nl == results[i][1] || zl == results[i][2]) {
			t.Error("Not correct result!")
		}
	}

}

func plusMinus(input []int) (float32, float32, float32) {
	var (
		p, n, z float32 = 0, 0, 0
		l               = len(input)
	)

	for i := 0; i < l; i++ {
		if input[i] == 0 {
			z++
		} else if input[i] > 0 {
			p++
		} else if input[i] < 0 {
			n++
		}
	}
	return p / float32(l), n / float32(l), z / float32(l)
}

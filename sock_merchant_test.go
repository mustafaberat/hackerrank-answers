package algorithms

import (
	"fmt"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 9
	got = sockMerchant([]int32{4, 5, 5, 5, 6, 6, 4, 1, 4, 4, 3, 6, 6, 3, 6, 1, 4, 5, 5, 5})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 3
	got = sockMerchant([]int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func sockMerchant(ar []int32) int32 {
	var (
		pairNumbers int32 = 0
		pairSocks   []int32
	)

	for _, sock := range ar {
		if !isInPair(pairSocks, sock) {
			pairSocks = append(pairSocks, sock)
			fmt.Println("-", pairSocks)
		} else {
			pairSocks = removePairSock(pairSocks, sock)
			fmt.Println("+", pairSocks)
			pairNumbers++
		}
	}
	return pairNumbers
}

func isInPair(socks []int32, sock int32) bool {
	for i := range socks {
		if socks[i] == sock {
			return true
		}
	}
	return false
}

func removePairSock(socks []int32, sock int32) []int32 {
	var newSocks []int32

	for _, s := range socks {
		if s != sock {
			newSocks = append(newSocks, s)
		}
	}

	return newSocks
}

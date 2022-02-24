package algorithms

import (
	"testing"
	"unicode"
)

func TestCamelcase(t *testing.T) {
	var (
		want int32
		got  int32
	)

	want = 5
	got = camelcase("saveChangesInTheEditor")
	if want != got {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

}

func camelcase(s string) int32 {
	var c int32 = 1

	for i, l := 0, len(s); i < l; i++ {
		if IsUpper(string(s[i])) {
			c++
		}
	}

	return c
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

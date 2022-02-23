package algorithms

import "testing"

func TestRepeatedString(t *testing.T) {
	var want, got int64

	want = 3
	got = repeatedString("bab", 8)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 4
	got = repeatedString("abcac", 10)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = 7
	got = repeatedString("aba", 10)
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func repeatedString(s string, n int64) int64 {
	var (
		letter     uint8 = 'a'
		l                = int64(len(s))
		repeatedC        = n / l
		remainingC       = n % l
	)

	counter := repeatedC * getLetterCountIntoText(s, letter)
	counter += getLetterCountIntoText(s[:remainingC], letter)

	return counter
}

func getLetterCountIntoText(text string, letter uint8) int64 {
	var counter int64 = 0

	for i := 0; i < len(text); i++ {
		if text[i] == letter {
			counter++
		}
	}

	return counter
}

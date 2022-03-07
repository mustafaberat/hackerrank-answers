package algorithms

import (
	"strconv"
	"testing"
)

func TestTimeConversion(t *testing.T) {
	var want, got string

	want = "12:45:54"
	got = timeConversion("12:45:54PM")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "19:05:45"
	got = timeConversion("07:05:45PM")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}

	want = "00:40:22"
	got = timeConversion("12:40:22AM")
	if got != want {
		t.Errorf("Not correct. Got: %v Want: %v \n", got, want)
	}
}

func timeConversion(s string) string {
	var (
		hour   = s[0:2]
		remain = s[2 : len(s)-2]
		format = s[len(s)-2:]
	)

	if format == "PM" {
		hourAsInt, _ := strconv.Atoi(hour)
		hour = strconv.Itoa((hourAsInt + 12) % 24)
		if hour == "0" {
			hour = "12"
		}
	} else if format == "AM" {
		if hour == "12" {
			hour = "00"
		}
	}

	return hour + remain
}

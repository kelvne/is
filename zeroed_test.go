package is

import (
	"testing"
)

func TestIntZero(t *testing.T) {
	cases := []*int{
		nil,
		new(int),
		createPointer(42),
		createPointer(0),
	}
	expected := []int{
		0,
		0,
		42,
		0,
	}

	for i, c := range cases {
		if got := IntZero(c); got != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], got)
		}
	}
}

func TestStringZero(t *testing.T) {
	cases := []*string{
		nil,
		new(string),
		createPointer("Hello"),
		createPointer(""),
		createPointer(" "),
	}
	expected := []string{
		"",
		"",
		"Hello",
		"",
		" ",
	}

	for i, c := range cases {
		if got := StringZero(c); got != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], got)
		}
	}
}

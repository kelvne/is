package is

import (
	"fmt"
	"testing"
)

func TestIntZero(t *testing.T) {
	t.Parallel()

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
		t.Run(fmt.Sprintf("expects %v to return the expected value %d", c, expected[i]), func(t *testing.T) {
			if got := IntZero(c); got != expected[i] {
				t.Errorf("Expected %d, got %d", expected[i], got)
			}
		})
	}
}

func TestStringZero(t *testing.T) {
	t.Parallel()

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
		t.Run(fmt.Sprintf("expects %v to return the expected value %s", c, expected[i]), func(t *testing.T) {
			if got := StringZero(c); got != expected[i] {
				t.Errorf("expected %s, got %s", expected[i], got)
			}
		})
	}
}

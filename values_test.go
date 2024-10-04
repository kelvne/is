package is

import (
	"testing"
)

type testValueStruct struct {
	name      string
	value     interface{}
	reference interface{}
	expected  bool
}

func TestValue(t *testing.T) {
	testCases := []testValueStruct{
		{"Works with pointers", createPointer(42), 0, true},
		{"Both integeres", 42, 0, true},
		{"Both strings", "hello", "", true},
		{"Both float64", 3.14, 0.0, true},
		{"Integers and strings", 42, "", false},
		{"Reversed string and integer", "hello", 0, false},
		{"Float64 and integer", 3.14, 0, false},
		{"Same struct type", struct{ Name string }{"test"}, struct{ Name string }{}, true},
		{"Different struct type", struct{ Name string }{"test"}, 0, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if result := Value(testCase.value, testCase.reference); result != testCase.expected {
				t.Log(testCase.name)
				t.Errorf(
					"value: %+v; reference: %+v; expected: %v; got: %v",
					testCase.value, testCase.reference, testCase.expected, result,
				)
			}
		})
	}
}

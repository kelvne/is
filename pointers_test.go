package is

import (
	"testing"
)

type testPointersStruct struct {
	name         string
	value        interface{}
	checkForNull bool
	expected     bool
}

func TestPointer(t *testing.T) {
	testCases := []testPointersStruct{
		{"value: nil pointer, checkForNull: true", (*int)(nil), true, false},
		{"value: nil pointer, checkForNull: false", (*int)(nil), false, true},
		{"value: non-nil pointer, checkForNull: true", new(int), true, true},
		{"value: non-nil pointer, checkForNull: false", new(int), false, true},
		{"value: non-pointer, checkForNull: true", 42, true, false},
		{"value: non-pointer, checkForNull: true", 42, false, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if result := Pointer(testCase.value, testCase.checkForNull); result != testCase.expected {
				t.Log(testCase.name)
				t.Errorf(
					"value: %+v; checkForNull: %+v; expected: %v; got: %v",
					testCase.value, testCase.checkForNull, result, testCase.expected,
				)
			}
		})
	}
}

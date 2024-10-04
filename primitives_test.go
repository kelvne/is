package is

import (
	"reflect"
	"testing"
)

type testPrimitivesStruct[T any] struct {
	name     string
	value    interface{}
	expected T
}

func createPointer[T any](value T) *T {
	return &value
}

func zeroValue[T any](value interface{}) T {
	return reflect.Zero(reflect.TypeOf(value)).Interface().(T)
}

func TestBool(t *testing.T) {
	zeroBool := zeroValue[bool](true)

	tests := []testPrimitivesStruct[bool]{
		{"Valid false boolean pointer", createPointer(false), false},
		{"Valid true boolean pointer", createPointer(true), true},
		{"Invalid true boolean pointer", false, zeroBool},
		{"Invalid true boolean pointer", true, zeroBool},
		{"Invalid value", "true", zeroBool},
		{"Invalid invalid structured value", struct{ name string }{}, zeroBool},
	}

	for _, testCases := range tests {
		t.Run(testCases.name, func(t *testing.T) {
			got, err := Bool(testCases.value)
			if got != testCases.expected {
				if expectedErr := NewNotPointerError(got); err != nil && reflect.TypeOf(err) != reflect.TypeOf(expectedErr) {
					t.Errorf("Expected error not returned. Expected: %v; Got: %v", expectedErr, err)
					return
				}

				t.Errorf("Expected value not returned. Expected: %v; Got: %v", testCases.expected, got)
			}
		})
	}
}

func TestByte(t *testing.T) {
	zeroByte := zeroValue[byte](byte(0))

	tests := []testPrimitivesStruct[byte]{
		{"Valid byte pointer", createPointer(byte(0)), byte(0)},
		{"Valid byte pointer", createPointer(byte(0b10001)), byte(0b10001)},
		{"Invalid byte pointer", byte(0), zeroByte},
		{"Invalid byte pointer", byte(0b11101), zeroByte},
		{"Invalid value", true, zeroByte},
		{"Invalid invalid structured value", struct{ name string }{}, zeroByte},
	}

	for _, testCases := range tests {
		t.Run(testCases.name, func(t *testing.T) {
			got, err := Byte(testCases.value)
			if got != testCases.expected {
				if expectedErr := NewNotPointerError(got); err != nil && reflect.TypeOf(err) != reflect.TypeOf(expectedErr) {
					t.Errorf("Expected error not returned. Expected: %v; Got: %v", expectedErr, err)
					return
				}

				t.Errorf("Expected value not returned. Expected: %v; Got: %v", testCases.expected, got)
			}
		})
	}
}

func TestInt(t *testing.T) {
	zeroInt := zeroValue[int](int(0))

	tests := []testPrimitivesStruct[int]{
		{"Valid integer pointer", createPointer(5), 5},
		{"Valid integer pointer", createPointer(1024), 1024},
		{"Invalid integer pointer", 5, zeroInt},
		{"Invalid integer pointer", 1024, zeroInt},
		{"Invalid value", "true", zeroInt},
		{"Invalid invalid structured value", struct{ name string }{}, zeroInt},
	}

	for _, testCases := range tests {
		t.Run(testCases.name, func(t *testing.T) {
			got, err := Int(testCases.value)
			if got != testCases.expected {
				if expectedErr := NewNotPointerError(got); err != nil && reflect.TypeOf(err) != reflect.TypeOf(expectedErr) {
					t.Errorf("Expected error not returned. Expected: %v; Got: %v", expectedErr, err)
					return
				}

				t.Errorf("Expected value not returned. Expected: %v; Got: %v", testCases.expected, got)
			}
		})
	}
}

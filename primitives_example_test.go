package is_test

import (
	"fmt"

	"github.com/kelvne/is"
)

func ExampleGeneric() {
	type sampleStruct struct {
		name string
	}

	validPointer := &sampleStruct{name: "example"}
	invalidPointer := "other value"
	var nilPointer *sampleStruct = nil

	value, err := is.Generic[sampleStruct](validPointer)
	fmt.Println(value, err)

	invalidValue, err := is.Generic[sampleStruct](invalidPointer)
	fmt.Println(invalidValue, err)

	nilTypedValue, err := is.Generic[sampleStruct](nilPointer)
	fmt.Println(nilTypedValue, err)

	nilValue, err := is.Generic[sampleStruct](nil)
	fmt.Println(nilValue, err)

	// Output: {example} <nil>
	// {} Provided value other value (string) is not a pointer of the expected type.
	// {} Provided value <nil> (*is_test.sampleStruct) is not a pointer of the expected type.
	// {} Provided value is nil
}

func ExampleBool() {
	validPointer := new(bool)
	*validPointer = true

	value, err := is.Bool(validPointer)
	fmt.Println(value, err)

	value, err = is.Bool("other")
	fmt.Println(value, err)

	// Output: true <nil>
	// false Provided value other (string) is not a pointer of the expected type.
}

func ExampleString() {
	validPointer := new(string)
	*validPointer = "example string"

	value, err := is.String(validPointer)
	fmt.Println(value, err)

	value, err = is.String(123)
	fmt.Println(value, err)

	// Output: example string <nil>
	//  Provided value 123 (int) is not a pointer of the expected type.
}

func ExampleInt() {
	validPointer := new(int)
	*validPointer = 123

	value, err := is.Int(validPointer)
	fmt.Println(value, err)

	value, err = is.Int("abc")
	fmt.Println(value, err)

	// Output: 123 <nil>
	// 0 Provided value abc (string) is not a pointer of the expected type.
}

func ExampleByte() {
	validPointer := new(byte)
	*validPointer = 0b11110001

	value, err := is.Byte(validPointer)
	fmt.Println(value, err)

	value, err = is.Byte("oops")
	fmt.Println(value, err)

	// Output: 241 <nil>
	// 0 Provided value oops (string) is not a pointer of the expected type.
}

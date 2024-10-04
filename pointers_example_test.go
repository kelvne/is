package is_test

import (
	"fmt"

	"github.com/kelvne/is"
)

func ExamplePointer() {
	var nilPointer *int = nil
	var nonNilPointer *int = new(int)

	fmt.Println(is.Pointer(nilPointer, true))
	fmt.Println(is.Pointer(nilPointer, false))
	fmt.Println(is.Pointer(nonNilPointer, true))
	fmt.Println(is.Pointer(nonNilPointer, false))
	// Output: false
	// true
	// true
	// true
}

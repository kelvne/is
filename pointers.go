package is

import "reflect"

// Pointer checks if the value is a pointer.
//
// If checkForNull is true, It returns true if the value is a pointer and has a valid memory address.
// If checkForNull is false it returns true if the value is a pointer but no valid memory address (nil).
// Otherwise, it returns false.
func Pointer(value interface{}, checkForNull bool) bool {
	valueOf := reflect.ValueOf(value)

	return valueOf.Kind() == reflect.Ptr && (checkForNull && !valueOf.IsNil() || !checkForNull)
}

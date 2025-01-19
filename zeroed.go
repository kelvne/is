package is

// GenericZero validates if the provided pointer is to a given generic type,
// and if it is, it returns the value.
// Otherwise, it returns just the zero value without an error.
func GenericZero[T any](value interface{}) T {
	var zeroOf T

	if Pointer(value, true) && Value(value, zeroOf) {
		return *value.(*T)
	}

	return zeroOf
}

// IntZero validates if the provided pointer is to an int,
// and if it is, it returns the value.
// Otherwise, it returns the zero value 0 without an error.
func IntZero(value interface{}) int {
	return GenericZero[int](value)
}

// StringZero validates if the provided pointer is to a string,
// and if it is, it returns the value.
// Otherwise, it returns the zero value "" without an error.
func StringZero(value interface{}) string {
	return GenericZero[string](value)
}

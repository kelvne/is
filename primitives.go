package is

func commonReturn[T any](value T, err error) (T, error) {
	return value, err
}

// Generic validates if the provided pointer is to a given generic type,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func Generic[T any](value interface{}) (T, error) {
	var zeroOf T

	if Pointer(value, true) && Value(value, zeroOf) {
		return *value.(*T), nil
	}

	return commonReturn(zeroOf, &NotPointerError{Value: value})
}

// Bool validates if the provided pointer is to a boolean,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func Bool(value interface{}) (bool, error) {
	return Generic[bool](value)
}

// Byte validates if the provided pointer is to a byte,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func Byte(value interface{}) (byte, error) {
	return Generic[byte](value)
}

// Int validates if the provided pointer is to an int,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func Int(value interface{}) (int, error) {
	return Generic[int](value)
}

// String validates if the provided pointer is to a string,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func String(value interface{}) (string, error) {
	return Generic[string](value)
}

package is

func commonReturn[T any](value T, err error) (T, error) {
	return value, err
}

// Bool validates if the provided pointer is to a boolean,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func Bool(value interface{}) (bool, error) {
	if Pointer(value, true) && Value(value, true) {
		return *value.(*bool), nil
	}

	return commonReturn(false, &NotPointerError{Value: value})
}

// Byte validates if the provided pointer is to a byte,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func Byte(value interface{}) (byte, error) {
	if Pointer(value, true) && Value(value, byte(0)) {
		return *value.(*byte), nil
	}

	return commonReturn(byte(0), &NotPointerError{Value: value})
}

// Int validates if the provided pointer is to a int,
// and if it is, it returns the value.
// Otherwise, it returns an NotPointerError.
func Int(value interface{}) (int, error) {
	if Pointer(value, true) && Value(value, int(0)) {
		return *value.(*int), nil
	}

	return commonReturn(int(0), &NotPointerError{Value: value})
}

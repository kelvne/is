package is

import "fmt"

// To make sure the NotPointerError implements the error interface
var _ (error) = (*NotPointerError)(nil)

// NotPointerError is an error type that indicates
// the provided value is not a pointer.
type NotPointerError struct {
	Value interface{}
}

// NewNotPointerError creates a NotPointerError and the value that caused it.
// The value is intentionally an interface since it is not required.
//
// Returns a NotPointerError.
func NewNotPointerError(value interface{}) NotPointerError {
	return NotPointerError{Value: value}
}

// Error returns the error message.
func (e NotPointerError) Error() string {
	if e.Value == nil {
		return "Provided value is nil"
	}

	return fmt.Sprintf(
		"Provided value %#+v (%T) is not a valid pointer of the expected type.",
		e.Value,
		e.Value,
	)
}

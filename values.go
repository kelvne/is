package is

import "reflect"

// Value [T any] checks if the value is similar type than the reference.
//
// If the value is assignable to the reference type, it returns true, otherwise false.
func Value[V any](value V, reference interface{}) bool {
	typeOf := reflect.TypeOf(reference)
	valueOf := reflect.ValueOf(value)

	if valueOf.Kind() == reflect.Ptr {
		return valueOf.Type().Elem().AssignableTo(typeOf)
	}

	return valueOf.Type().AssignableTo(typeOf)
}

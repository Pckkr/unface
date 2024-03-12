package unface

import (
	"errors"
	"reflect"
)

// ErrUnsupportedType is returned when unface is called with an unsupported type
var ErrUnsupportedType = errors.New("unsupported type")

// Unface converts src interface to dest type
func Unface[T any](src interface{}, dest *T) error {
	srcValue := reflect.ValueOf(src)
	if err := validateSrc(srcValue); err != nil {
		return err
	}

	destValue := reflect.ValueOf(dest)

	if srcValue.Type() == destValue.Elem().Type() {
		*dest = src.(T)
		return nil
	}

	return ErrUnsupportedType
}

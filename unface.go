package unface

import (
	"reflect"
)

// Unface converts src interface to dest type
func Unface[T any](src interface{}, dest *T) error {
	srcValue := reflect.ValueOf(src)
	if err := validateSrc(srcValue); err != nil {
		return err
	}

	return nil
}

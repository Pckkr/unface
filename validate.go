package unface

import (
	"errors"
	"reflect"
)

// ErrSrcInvalid is returned when unface is called with an invalid source
var ErrSrcInvalid = errors.New("invalid source")

// ErrSrcNil is returned when unface is called with a nil source
var ErrSrcNil = errors.New("nil source")

func isValid(v reflect.Value) bool {
	return v.IsValid()
}

func isNilPointer(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr && v.IsNil()
}

func validateSrc(src reflect.Value) error {
	if !isValid(src) {
		return ErrSrcInvalid
	}
	if isNilPointer(src) {
		return ErrSrcNil
	}
	return nil
}

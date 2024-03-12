package unface

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		var src int = 42
		if !isValid(reflect.ValueOf(src)) {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		var src interface{} = nil
		if isValid(reflect.ValueOf(src)) {
			t.Errorf("Expected false, got true")
		}
	})
}

func TestIsNilPointer(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var src *int
		if !isNilPointer(reflect.ValueOf(src)) {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("NotPointer", func(t *testing.T) {
		var src int = 42
		if isNilPointer(reflect.ValueOf(src)) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("NotNil", func(t *testing.T) {
		var src int = 42
		if isNilPointer(reflect.ValueOf(&src)) {
			t.Errorf("Expected false, got true")
		}
	})
}

func TestValidateSrc(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		var src int = 42
		if err := validateSrc(reflect.ValueOf(src)); err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		var src interface{} = nil
		if err := validateSrc(reflect.ValueOf(src)); err != ErrSrcInvalid {
			t.Errorf("Expected %v, got %v", ErrSrcInvalid, err)
		}
	})

	t.Run("Nil", func(t *testing.T) {
		var src *int
		if err := validateSrc(reflect.ValueOf(src)); err != ErrSrcNil {
			t.Errorf("Expected %v, got %v", ErrSrcNil, err)
		}
	})
}

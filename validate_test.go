package unface

import (
	"reflect"
	"testing"
)

// isValid

func TestIsValidValid(t *testing.T) {
	var src int = 42
	if !isValid(reflect.ValueOf(src)) {
		t.Errorf("Expected true, got false")
	}
}

func TestIsValidInvalid(t *testing.T) {
	var src interface{} = nil
	if isValid(reflect.ValueOf(src)) {
		t.Errorf("Expected false, got true")
	}
}

// isNilPointer

func TestIsNilPointerNil(t *testing.T) {
	var src *int
	if !isNilPointer(reflect.ValueOf(src)) {
		t.Errorf("Expected true, got false")
	}
}

func TestIsNilPointerNotPointer(t *testing.T) {
	var src int = 42
	if isNilPointer(reflect.ValueOf(src)) {
		t.Errorf("Expected false, got true")
	}
}

func TestIsNilPointerNotNil(t *testing.T) {
	var src int = 42
	if isNilPointer(reflect.ValueOf(&src)) {
		t.Errorf("Expected false, got true")
	}
}

// validateSrc

func TestValidateSrcValid(t *testing.T) {
	var src int = 42
	if err := validateSrc(reflect.ValueOf(src)); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestValidateSrcInvalid(t *testing.T) {
	var src interface{} = nil
	if err := validateSrc(reflect.ValueOf(src)); err != ErrSrcInvalid {
		t.Errorf("Expected %v, got %v", ErrSrcInvalid, err)
	}
}

func TestValidateSrcNil(t *testing.T) {
	var src *int
	if err := validateSrc(reflect.ValueOf(src)); err != ErrSrcNil {
		t.Errorf("Expected %v, got %v", ErrSrcNil, err)
	}
}

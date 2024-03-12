package unface

import "testing"

func TestUnfaceValidSrc(t *testing.T) {
	var src int = 42
	var dest int
	if err := Unface(src, &dest); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestUnfaceInvalidSrc(t *testing.T) {
	var src interface{} = nil
	var dest int
	if err := Unface(src, &dest); err != ErrSrcInvalid {
		t.Errorf("Expected error, got nil")
	}
}

func TestUnfaceNilSrc(t *testing.T) {
	var src *int = nil
	var dest int
	if err := Unface(src, &dest); err != ErrSrcNil {
		t.Errorf("Expected error, got nil")
	}
}

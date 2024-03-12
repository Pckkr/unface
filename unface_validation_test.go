package unface

import "testing"

func TestUnfaceValidation(t *testing.T) {
	t.Run("Invalid src, valid dest", func(t *testing.T) {
		var src interface{} = nil
		var dest int
		if err := Unface(src, &dest); err != ErrSrcInvalid {
			t.Errorf("Expected %v, got %v", ErrSrcInvalid, err)
		}
	})

	t.Run("Nil src, valid dest", func(t *testing.T) {
		var src *int = nil
		var dest int
		if err := Unface(src, &dest); err != ErrSrcNil {
			t.Errorf("Expected %v, got %v", ErrSrcNil, err)
		}
	})
}

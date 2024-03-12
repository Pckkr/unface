package unface

import (
	"reflect"
	"testing"
)

func helperSameTypeErrorCheck(t *testing.T, src, dest interface{}, err, expected error) {
	t.Helper()
	if err != expected {
		t.Errorf("Expected %v, got %v", expected, err)
	} else if err == nil && !reflect.DeepEqual(src, dest) {
		t.Errorf("Dest should be %v, got %v", src, dest)
	}
}

func TestUnfaceSameTypeValid(t *testing.T) {
	t.Run("uint8 to uint8", func(t *testing.T) {
		var src uint8 = 42
		var dest uint8
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("uint16 to uint16", func(t *testing.T) {
		var src uint16 = 42
		var dest uint16
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("uint32 to uint32", func(t *testing.T) {
		var src uint32 = 42
		var dest uint32
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("uint64 to uint64", func(t *testing.T) {
		var src uint64 = 42
		var dest uint64
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("uint to uint", func(t *testing.T) {
		var src uint = 42
		var dest uint
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("int8 to int8", func(t *testing.T) {
		var src int8 = 42
		var dest int8
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("int16 to int16", func(t *testing.T) {
		var src int16 = 42
		var dest int16
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("int32 to int32", func(t *testing.T) {
		var src int32 = 42
		var dest int32
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("int64 to int64", func(t *testing.T) {
		var src int64 = 42
		var dest int64
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("int to int", func(t *testing.T) {
		var src int = 42
		var dest int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("float32 to float32", func(t *testing.T) {
		var src float32 = 42
		var dest float32
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("float64 to float64", func(t *testing.T) {
		var src float64 = 42
		var dest float64
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("complex64 to complex64", func(t *testing.T) {
		var src complex64 = 42
		var dest complex64
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("complex128 to complex128", func(t *testing.T) {
		var src complex128 = 42
		var dest complex128
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("uintptr to uintptr", func(t *testing.T) {
		var src uintptr = 42
		var dest uintptr
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("bool to bool", func(t *testing.T) {
		var src bool = true
		var dest bool
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("string to string", func(t *testing.T) {
		var src string = "hello"
		var dest string
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("[]byte to []byte", func(t *testing.T) {
		var src []byte = []byte("hello")
		var dest []byte
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("[]int to []int", func(t *testing.T) {
		var src []int = []int{1, 2, 3}
		var dest []int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("map[int]string to map[int]string", func(t *testing.T) {
		var src map[int]string = map[int]string{1: "one", 2: "two"}
		var dest map[int]string
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("map[string]int to map[string]int", func(t *testing.T) {
		var src map[string]int = map[string]int{"one": 1, "two": 2}
		var dest map[string]int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})

	t.Run("struct to struct", func(t *testing.T) {
		type person struct {
			Name string
			Age  int
		}

		var src person = person{"John", 42}
		var dest person
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, nil)
	})
}

func TestUnfaceSameTypeInvalid(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		var src int = 42
		var dest string
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("string to int", func(t *testing.T) {
		var src string = "42"
		var dest int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("[]int to []string", func(t *testing.T) {
		var src []int = []int{1, 2, 3}
		var dest []string
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("map[int]string to map[string]int", func(t *testing.T) {
		var src map[int]string = map[int]string{1: "one", 2: "two"}
		var dest map[string]int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("map[string]int to int", func(t *testing.T) {
		var src map[string]int = map[string]int{"one": 1, "two": 2}
		var dest int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("struct to string", func(t *testing.T) {
		type person struct {
			Name string
			Age  int
		}

		var src person = person{"John", 42}
		var dest string
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("struct to int", func(t *testing.T) {
		type person struct {
			Name string
			Age  int
		}

		var src person = person{"John", 42}
		var dest int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("struct to []int", func(t *testing.T) {
		type person struct {
			Name string
			Age  int
		}

		var src person = person{"John", 42}
		var dest []int
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})

	t.Run("struct to map[int]string", func(t *testing.T) {
		type person struct {
			Name string
			Age  int
		}

		var src person = person{"John", 42}
		var dest map[int]string
		err := Unface(src, &dest)
		helperSameTypeErrorCheck(t, src, dest, err, ErrUnsupportedType)
	})
}

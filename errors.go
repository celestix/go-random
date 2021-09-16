package random

import "errors"

// A struct Error records a failed conversion.
type Error struct {
	Func string // the failing function
	Err  error  // the reason the conversion failed (e.g. ErrExceed, ErrUnsupported, etc.)
}

func (e *Error) Error() string {
	return "random." + e.Func + ": " + e.Err.Error()
}

// This error will be returned if 'n' i.e. number of elements is greater than 'a' i.e. slice.
var ErrExceed = errors.New("n exceeded a")

// This error will be returned if the input type is not supported by the function.
var ErrUnsupported = errors.New("unsupported type")

// This error will be returned if the ending number of range is smaller than starting one.
var ErrEndNumSmaller = errors.New("endNum must be greater than startNum")

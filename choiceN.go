package random

import (
	"fmt"
	"reflect"
)

// ChoiceN indexes its parameters and pick n random choices from them.
// n (type int) is the number of values to be randomly chosen.
// 'a' can be of any type: string, integer, floats, slices, bool, etc.
// It returns the randomly chosen values in a slice of type []interface{} and any write error encountered.
func ChoiceN(n int, a ...interface{}) (interface{}, error) {
	const fn = "ChoiceN"
	var err error
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	var cs []interface{}
	for i := 1; i <= n; i++ {
		if len(a) > 1 {
			if reflect.TypeOf(a[i-1]).Kind() == reflect.Slice {
				return nil, &Error{fn, fmt.Errorf("%w: slice", ErrUnsupported)}
			}
			if reflect.TypeOf(a[i-1]).Kind() == reflect.Array {
				return nil, &Error{fn, fmt.Errorf("%w: array", ErrUnsupported)}
			}
		}
		c := Choice(a...)
		cs = append(cs, c)
		a = removeChoiceFromSlice(c, a)
	}
	return cs, err
}

// ChoiceStringN indexes the string slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a' must be of type []string.
// It returns n randomly chosen values of type string.
func ChoiceStringN(n int, a []string) (interface{}, error) {
	const fn = "ChoiceStringN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceIntN indexes the int slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a' must be of type []int.
// It returns n randomly chosen values of type int.
func ChoiceIntN(n int, a []int) (interface{}, error) {
	const fn = "ChoiceIntN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceInt8N indexes the int8 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int8.
// It returns n randomly chosen values of type interface{}.
func ChoiceInt8N(n int, a []int8) (interface{}, error) {
	const fn = "ChoiceInt8N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceInt16N indexes the int16 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int16.
// It returns n randomly chosen values of type interface{}.
func ChoiceInt16N(n int, a []int16) (interface{}, error) {
	const fn = "ChoiceInt16N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceInt32N indexes the int32 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int32.
// It returns n randomly chosen values of type interface{}.
func ChoiceInt32N(n int, a []int32) (interface{}, error) {
	const fn = "ChoiceInt32N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceInt64N indexes the int64 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int64.
// It returns n randomly chosen values of type interface{}.
func ChoiceInt64N(n int, a []int64) (interface{}, error) {
	const fn = "ChoiceInt64N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceUintN indexes the uint slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint.
// It returns n randomly chosen values of type interface{}.
func ChoiceUintN(n int, a []uint) (interface{}, error) {
	const fn = "ChoiceUintN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceUint8N indexes the uint8 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint8.
// It returns n randomly chosen values of type interface{}.
func ChoiceUint8N(n int, a []uint8) (interface{}, error) {
	const fn = "ChoiceUint8N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceUint16N indexes the uint16 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint16.
// It returns n randomly chosen values of type interface{}.
func ChoiceUint16N(n int, a []uint16) (interface{}, error) {
	const fn = "ChoiceUint16N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceUint32N indexes the uint32 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint32.
// It returns n randomly chosen values of type interface{}.
func ChoiceUint32N(n int, a []uint32) (interface{}, error) {
	const fn = "ChoiceUint32N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceUint64N indexes the uint64 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint64.
// It returns n randomly chosen values of type interface{}.
func ChoiceUint64N(n int, a []uint64) (interface{}, error) {
	const fn = "ChoiceUint64N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceFloat32N indexes the float32 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []float32.
// It returns n randomly chosen values of type interface{}.
func ChoiceFloat32N(n int, a []float32) (interface{}, error) {
	const fn = "ChoiceFloat32N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// ChoiceFloat64N indexes the float64 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []float64.
// It returns n randomly chosen values of type interface{}.
func ChoiceFloat64N(n int, a []float64) (interface{}, error) {
	const fn = "ChoiceFloat64N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	intf := make([]interface{}, len(a))
	for i := range a {
		intf[i] = a[i]
	}
	return ChoiceN(n, intf...)
}

// removeChoiceFromSlice is one of the inner functions of this package.
// This function is used to remove a choice (type interface{}) from a slice of interfaces (type []interface{})
// It returns a slice of type []interface{}.
func removeChoiceFromSlice(c interface{}, a []interface{}) []interface{} {
	for x, i := range a {
		if i == c {
			a[x] = a[len(a)-1]
			a = a[:len(a)-1]
		}
	}
	return a
}

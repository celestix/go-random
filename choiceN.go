package random

// ChoiceN indexes its parameters and pick n random choices from them.
// n (type int) is the number of values to be randomly chosen.
// 'a' can be of any particular type: T
// It returns the randomly chosen values in a slice of type []T and any error encountered.
func ChoiceN[T any](n int, a ...T) ([]T, error) {
	const fn = "ChoiceN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceAnyN indexes its parameters and pick n random choices from them.
// n (type int) is the number of values to be randomly chosen.
// 'a' can be of any particular type: string, integer, floats, slices, bool, etc.
// It returns the randomly chosen values in a slice of type []T and any error encountered.
func ChoiceAnyN(n int, a ...any) ([]any, error) {
	const fn = "ChoiceN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceStringN indexes the string slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a' must be of type []string.
// It returns n randomly chosen values of type string.
func ChoiceStringN(n int, a []string) ([]string, error) {
	const fn = "ChoiceStringN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceIntN indexes the int slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a' must be of type []int.
// It returns n randomly chosen values of type int.
func ChoiceIntN(n int, a []int) ([]int, error) {
	const fn = "ChoiceIntN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceInt8N indexes the int8 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int8.
// It returns n randomly chosen values of type int8.
func ChoiceInt8N(n int, a []int8) ([]int8, error) {
	const fn = "ChoiceInt8N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceInt16N indexes the int16 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int16.
// It returns n randomly chosen values of type int16.
func ChoiceInt16N(n int, a []int16) ([]int16, error) {
	const fn = "ChoiceInt16N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceInt32N indexes the int32 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int32.
// It returns n randomly chosen values of type int32.
func ChoiceInt32N(n int, a []int32) ([]int32, error) {
	const fn = "ChoiceInt32N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceInt64N indexes the int64 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []int64.
// It returns n randomly chosen values of type int64.
func ChoiceInt64N(n int, a []int64) ([]int64, error) {
	const fn = "ChoiceInt64N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceUintN indexes the uint slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint.
// It returns n randomly chosen values of type uint.
func ChoiceUintN(n int, a []uint) ([]uint, error) {
	const fn = "ChoiceUintN"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceUint8N indexes the uint8 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint8.
// It returns n randomly chosen values of type uint8.
func ChoiceUint8N(n int, a []uint8) ([]uint8, error) {
	const fn = "ChoiceUint8N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceUint16N indexes the uint16 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint16.
// It returns n randomly chosen values of type uint16.
func ChoiceUint16N(n int, a []uint16) ([]uint16, error) {
	const fn = "ChoiceUint16N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceUint32N indexes the uint32 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint32.
// It returns n randomly chosen values of type uint32.
func ChoiceUint32N(n int, a []uint32) ([]uint32, error) {
	const fn = "ChoiceUint32N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceUint64N indexes the uint64 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []uint64.
// It returns n randomly chosen values of type uint64.
func ChoiceUint64N(n int, a []uint64) ([]uint64, error) {
	const fn = "ChoiceUint64N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceFloat32N indexes the float32 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []float32.
// It returns n randomly chosen values of type float32.
func ChoiceFloat32N(n int, a []float32) ([]float32, error) {
	const fn = "ChoiceFloat32N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

// ChoiceFloat64N indexes the float64 slice and pick n random choices from it.
// n (type int) is the number of values to be randomly chosen.
// Its parameter 'a'must be of type []float64.
// It returns n randomly chosen values of type float64.
func ChoiceFloat64N(n int, a []float64) ([]float64, error) {
	const fn = "ChoiceFloat64N"
	if n > len(a) {
		return nil, &Error{fn, ErrExceed}
	}
	return choiceN(n, a...)
}

func choiceN[T any](n int, a ...T) (cs []T, err error) {
	cs = make([]T, n)
	for i := 0; i < n; i++ {
		var index int
		index, err = Integer(0, len(a)-1)
		if err != nil {
			break
		}
		cs[i] = a[index]
		a = removeChoiceFromSlice(index, a)
	}
	return
}

// removeChoiceFromSlice is one of the inner functions of this package.
// This function is used to remove a choice (type interface{}) from a slice of interfaces (type []interface{})
// It returns a slice of type []interface{}.
func removeChoiceFromSlice[T any](index int, a []T) []T {
	for i := range a {
		if i == index {
			a[i] = a[len(a)-1]
			a = a[:len(a)-1]
			break
		}
	}
	return a
}

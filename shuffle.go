package random

import (
	"math/rand"
)

// Shuffle pseudo-randomizes the order of parameters.
// it can be of any type: string, integer, floats, slices, bool, etc.
// It returns the shuffled slice of type []interface{}.
func Shuffle(a ...interface{}) interface{} {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleStrings pseudo-randomizes the order of provided slice (type []string).
// parameter 'a' must be of type []string.
// It returns the shuffled slice of type []string.
func ShuffleStrings(a []string) []string {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt pseudo-randomizes the order of provided slice (type []int).
// parameter 'a' must be of type []int.
// It returns the shuffled slice of type []int.
func ShuffleInt(a []int) []int {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt8 pseudo-randomizes the order of provided slice (type []int8).
// parameter 'a' must be of type []int8.
// It returns the shuffled slice of type []int8.
func ShuffleInt8(a []int8) []int8 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt16 pseudo-randomizes the order of provided slice (type []int16).
// parameter 'a' must be of type []int16.
// It returns the shuffled slice of type []int16.
func ShuffleInt16(a []int16) []int16 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt32 pseudo-randomizes the order of provided slice (type []int32).
// parameter 'a' must be of type []int32.
// It returns the shuffled slice of type []int32.
func ShuffleInt32(a []int32) []int32 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleInt64 pseudo-randomizes the order of provided slice (type []int64).
// parameter 'a' must be of type []int64.
// It returns the shuffled slice of type []int64.
func ShuffleInt64(a []int64) []int64 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint pseudo-randomizes the order of provided slice (type []uint).
// parameter 'a' must be of type []uint.
// It returns the shuffled slice of type []uint.
func ShuffleUint(a []uint) []uint {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint8 pseudo-randomizes the order of provided slice (type []uint8).
// parameter 'a' must be of type []uint8.
// It returns the shuffled slice of type []uint8.
func ShuffleUint8(a []uint8) []uint8 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint16 pseudo-randomizes the order of provided slice (type []uint16).
// parameter 'a' must be of type []uint16.
// It returns the shuffled slice of type []uint16.
func ShuffleUint16(a []uint16) []uint16 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint32 pseudo-randomizes the order of provided slice (type []uint32).
// parameter 'a' must be of type []uint32.
// It returns the shuffled slice of type []uint32.
func ShuffleUint32(a []uint32) []uint32 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleUint64 pseudo-randomizes the order of provided slice (type []uint64).
// parameter 'a' must be of type []uint64.
// It returns the shuffled slice of type []uint64.
func ShuffleUint64(a []uint64) []uint64 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleFloat32 pseudo-randomizes the order of provided slice (type []float32).
// parameter 'a' must be of type []float32.
// It returns the shuffled slice of type []float32.
func ShuffleFloat32(a []float32) []float32 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

// ShuffleFloat64 pseudo-randomizes the order of provided slice (type []float64).
// parameter 'a' must be of type []float64.
// It returns the shuffled slice of type []float64.
func ShuffleFloat64(a []float64) []float64 {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	return a
}

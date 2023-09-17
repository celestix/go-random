package random

// ChoiceAny indexes its parameters and pick a random choice from them.
// it can be of any type: string, integer, floats, slices, bool, etc.
// It returns the randomly chosen value of input type.
func Choice[T any](a ...T) T {
	return a[rand.Intn(len(a))]
}

// ChoiceAny indexes its parameters and pick a random choice from them.
// it can be of any type: string, integer, floats, slices, bool, etc.
// It returns the randomly chosen value of type interface{}.
func ChoiceAny(a ...any) interface{} {
	return a[rand.Intn(len(a))]
}

// ChoiceString indexes the string slice and pick a random choice from it.
// Its parameter 'a' must be of type []string.
// It returns the randomly chosen value of type string.
func ChoiceString(a []string) string {
	return a[rand.Intn(len(a))]
}

// ChoiceInt indexes the int slice and pick a random choice from it.
// Its parameter 'a' must be of type []int.
// It returns the randomly chosen value of type int.
func ChoiceInt(a []int) int {
	return a[rand.Intn(len(a))]
}

// ChoiceInt8 indexes the int8 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int8.
// It returns the randomly chosen value of type int8.
func ChoiceInt8(a []int8) int8 {
	return a[rand.Intn(len(a))]
}

// ChoiceInt16 indexes the int16 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int16.
// It returns the randomly chosen value of type int16.
func ChoiceInt16(a []int16) int16 {
	return a[rand.Intn(len(a))]
}

// ChoiceInt32 indexes the int32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int32.
// It returns the randomly chosen value of type int32.
func ChoiceInt32(a []int32) int32 {
	return a[rand.Intn(len(a))]
}

// ChoiceInt64 indexes the int64 slice and pick a random choice from it.
// Its parameter 'a' must be of type []int64.
// It returns the randomly chosen value of type int64.
func ChoiceInt64(a []int64) int64 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint indexes the uint slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint.
// It returns the randomly chosen value of type uint.
func ChoiceUint(a []uint) uint {
	return a[rand.Intn(len(a))]
}

// ChoiceUint8 indexes the uint8 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint8.
// It returns the randomly chosen value of type uint8.
func ChoiceUint8(a []uint8) uint8 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint16 indexes the uint16 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint16.
// It returns the randomly chosen value of type uint16.
func ChoiceUint16(a []uint16) uint16 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint indexes the uint32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint32.
// It returns the randomly chosen value of type uint32.
func ChoiceUint32(a []uint32) uint32 {
	return a[rand.Intn(len(a))]
}

// ChoiceUint64 indexes the uint64 slice and pick a random choice from it.
// Its parameter 'a' must be of type []uint64.
// It returns the randomly chosen value of type uint64.
func ChoiceUint64(a []uint64) uint64 {
	return a[rand.Intn(len(a))]
}

// ChoiceFloat32 indexes the float32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []float32.
// It returns the randomly chosen value of type float32.
func ChoiceFloat32(a []float32) float32 {
	return a[rand.Intn(len(a))]
}

// ChoiceFloat64 indexes the float32 slice and pick a random choice from it.
// Its parameter 'a' must be of type []float64.
// It returns the randomly chosen value of type float64.
func ChoiceFloat64(a []float64) float64 {
	return a[rand.Intn(len(a))]
}

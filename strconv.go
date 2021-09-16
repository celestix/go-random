package random

import (
	"strconv"
)

// Atoi function allows you to get a random choice from a slice of type []string parsed into an integer with base 10.
// It returns the randomly chosen integer value from a slice containing elements (type string) and any write error encountered.
func Atoi(a []string) (int, error) {
	return strconv.Atoi(ChoiceString(a))
}

// Itoa function allows you to get a random choice from a slice of type []int parsed into a string.
// It returns the randomly chosen string from a slice containing elements (type int).
func Itoa(a []int) string {
	return strconv.Itoa(ChoiceInt(a))
}

// Quote function allows you to get a random choice from a slice of type []string as a double-quoted string.
// It returns the double-quoted randomly chosen string from a slice containing elements (type string).
func Quote(a []string) string {
	return strconv.Quote(ChoiceString(a))
}

// Unquote function allows you to get a unquoted random choice from a slice of type []string containing double-quoted strings.
// It returns the unquoted randomly chosen string from a slice containing double-quoted elements (type string) and any write error encountered.
func Unquote(a []string) (string, error) {
	const fn = "Unquote"
	s, err := strconv.Unquote(ChoiceString(a))
	if err != nil {
		err = &Error{fn, err}
	}
	return s, err
}

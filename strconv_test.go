package random_test

import (
	"fmt"

	random "github.com/celestix/go-random/v2"
)

func ExampleAtoi() {
	a := []string{"1", "332", "43"}
	r, err := random.Atoi(a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: any one string from the provided three, parsed into an integer with base 10.
}

func ExampleItoa() {
	a := []int{1, 332, 43}
	r := random.Itoa(a)
	fmt.Println(r)
	// Output: any one integer from these three, parsed into a string.
}

func ExampleQuote() {
	a := []string{"Hello", "Hi", "Nice"}
	r := random.Quote(a)
	fmt.Println(r)
	// Output: any one string from these three as a double-quoted string.
}

func ExampleUnquote() {
	a := []string{`"Hello"`, `"Hi"`, `"Nice"`}
	r, err := random.Unquote(a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: any one string from these three as an unquoted string.
}

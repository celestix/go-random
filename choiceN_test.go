package random_test

import (
	"fmt"

	random "github.com/anonyindian/go-random/v2"
)

func ExampleChoiceN() {
	r, err := random.ChoiceN(2, "a", "Hello", "false")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: Slice of length 2 containing any two from the provided args.
}

func ExampleChoiceStringN() {
	a := []string{"Hi", "Hello", "Why"}
	r, err := random.ChoiceStringN(2, a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: Slice of length 2 containing any two from the provided strings.
}

func ExampleChoiceIntN() {
	a := []int{41, 24, 33}
	r, err := random.ChoiceIntN(2, a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: Slice of length 2 containing any two from the provided integers.
}

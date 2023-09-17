package random_test

import (
	"fmt"

	random "github.com/celestix/go-random/v2"
)

func ExampleChoice() {
	r := random.ChoiceAny("a", 1, true)
	fmt.Println(r)
	// Output: "a" or 1 or true
}

func ExampleChoiceString() {
	a := []string{"a", "b", "c"}
	r := random.ChoiceString(a)
	fmt.Println(r)
	// Output: "a" or "b" or "c"
}

func ExampleChoiceInt() {
	a := []int{1, 2, 3}
	r := random.ChoiceInt(a)
	fmt.Println(r)
	// Output: 1 or 2 or 3
}

func ExampleChoiceFloat32() {
	a := []float32{1.223, 3.2, 4.2922}
	r := random.ChoiceFloat32(a)
	fmt.Println(r)
	// Output: 1.223 or 3.2 or 4.2922
}

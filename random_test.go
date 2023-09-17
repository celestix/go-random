package random_test

import (
	"fmt"

	random "github.com/celestix/go-random/v2"
)

func ExampleBool() {
	r := random.Bool()
	fmt.Println(r)
	// Output: either true or false
}

func ExampleInteger() {
	r, err := random.Integer(10, 20)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: any one integer from the range [10, 20]
}

func ExampleFloat32() {
	r, err := random.Float32(1.292, 1.388)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: any one float32 value from the range [1.292, 1.388]
}

func ExampelIntegerN() {
	r, err := random.IntegerN(10, 20, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: an array containing 2 integers from the range [10, 20]
}

func ExampleFloat32N() {
	r, err := random.Float32N(1.11, 2.22, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r)
	// Output: an array containing 2 float32 values from the range [1.11, 2.22]
}

package random_test

import (
	"fmt"

	random "github.com/anonyindian/go-random"
)

func ExampleShuffle() {
	r := random.Shuffle("Hello", "World", 4455, 2283, true)
	fmt.Println(r)
	// Output: shuffled slice containing the provided args.
}

func ExampleShuffleStrings() {
	sentenceSlice := []string{"Hello", "I", "am", "a", "human", "being"}
	r := random.ShuffleStrings(sentenceSlice)
	fmt.Println(r)
	// Output: shuffled slice containing elements of sentenceSlice.
}

func ExampleShuffleInt() {
	intSlice := []int{1833, 38382, 28282, 2929}
	r := random.ShuffleInt(intSlice)
	fmt.Println(r)
	// Output: shuffled slice containing elements of intSlice.
}

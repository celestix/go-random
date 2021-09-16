package main

import (
	"fmt"
	"strings"

	random "github.com/anonyindian/go-random"
)

func main() {
	fmt.Println(random.ChoiceN(2, 1, 2, 3))
	fmt.Println("Result of toss is", Toss())
	fmt.Println("I rolled a dice and the result was", RollDice())
	fmt.Println("Solve the jumbled sentence:", JumbleSentence("It's a sunny day"))
}

// This is an example function to explain the use of random.Bool()
func RandomBoolGeneration() string {
	b := random.Bool()
	if b {
		// will return the following value if b == true
		return "Yeah, you were right"
	}
	// will return the following value if b == false
	return "No, you were wrong"
}

// A function which returns any one value from heads and tails with the help random.Choice()
func Toss() string {
	r := random.Choice("heads", "tails")
	return fmt.Sprintf("%v", r)
}

// A function which return number from 1 to 6 like on a dice of 6 faces.
// random.Integer() was used for it.
func RollDice() int {
	i, err := random.Integer(1, 6)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i
}

// A function to return a string with random person with a random profession in a string.
// String format: "{person} will become {"profession"}
// random.ChoiceString() was used for it.
func ChooseRandomPersonWithProfession() string {
	// Our data
	people := []string{"John", "Maureen", "Will", "Heroki", "Vijay", "Bill", "Penni", "Dawn", "Judi"}
	professions := []string{"doctor", "engineer", "programmer", "teacher", "nothing", "pilot", "scientist"}

	return random.ChoiceString(people) + " will become " + random.ChoiceString(professions)
}

// A function to let you know how does Shuffle functions work.
// We've used random.ShuffleStrings() function to jumble the sentence of type string.
// All the Shuffle functions are meant to be used in a similar way.
func JumbleSentence(sentence string) string {
	return strings.Join(random.ShuffleStrings(strings.Fields(sentence)), " ")
}

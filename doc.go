/*
Package random implements pseudo-random number generators for various distributions.

Example Usage:

	// start with some source data to use
	people := []string{"John", "Maureen", "Will", "Heroki", "Vijay", "Bill", "Penni", "Dawn", "Judi"}
	professions := []string{"doctor", "engineer", "programmer", "teacher", "nothing", "pilot", "scientist"}

	fmt.Println(random.ChoiceString(people) + " will become " + random.ChoiceString(professions))

The MIT License Copyright (c) 2021 Veer (anonyindian)
*/
package random

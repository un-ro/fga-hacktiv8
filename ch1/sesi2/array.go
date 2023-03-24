package main

import "strings"

func main() {
	// Declare array
	var names [3]string

	// Assign value to array
	names = [3]string{"John", "Wick"}
	names[2] = "Doe"

	println(names[0])
	println(names[1])
	println(names[2])

	// Modify array
	var fruits = [4]string{"apel", "anggur", "pisang"}
	fruits[0] = "apple"
	fruits[1] = "grape"
	fruits[2] = "banana"

	println(fruits[0])

	// Looping array with index and value
	for i, fruit := range fruits {
		println(i, fruit)
	}

	// Repeat strings builtin
	println(strings.Repeat("=", 20))
}

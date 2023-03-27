package main

import "fmt"

func main() {
	arr := []string{"a", "k", "u", "c", "i", "n", "t", "a"}

	// Create a map to store the vowel count
	vowelCount := make(map[string]int)

	// Define a closure that takes a string as an argument and increments
	// the count for each vowel in the string
	addVowelCount := func(s string) {
		for _, r := range s {
			switch r {
			case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
				vowelCount[string(r)]++
			}
		}
	}

	// Iterate over each element of the array and call the closure
	for _, s := range arr {
		addVowelCount(s)
	}

	// Print the vowel count
	fmt.Println(vowelCount)
}

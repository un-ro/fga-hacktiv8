package main

import "fmt"

// Session 3
func main() {
	words := "selamat malam"
	wordMap := make(map[string]int)

	for _, word := range words {
		fmt.Println(string(word))

		if wordMap[string(word)] == 0 {
			wordMap[string(word)] = 1
		} else {
			wordMap[string(word)]++
		}
	}

	fmt.Println(wordMap)
}

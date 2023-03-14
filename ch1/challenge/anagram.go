package main

import "fmt"

func isAnagram(a, b string) bool {
	for i := 0; i < len(a); i++ {
		fmt.Println(string(a[i]))
		var isFound = false
		for j := 0; j < len(b); j++ {
			fmt.Println(string(b[j]))
			if a[i] == b[j] {
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}
	return true
}

func main() {
	word1 := "conifers"
	word2 := "firscone"

	fmt.Println(word1, " ", word2, " ==> ", isAnagram(word1, word2))
}

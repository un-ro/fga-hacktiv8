package main

import "fmt"

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		// Debug
		// fmt.Println(str[i], " ", str[len(str)-i-1])
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	word := "kapalapijak"

	fmt.Println(word, "==> ", isPalindrome(word))
}

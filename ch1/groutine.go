package main

import "fmt"

func main() {
	words := []string{"saya", "tinggal", "di", "jakarta"}
	ch := make(chan string)

	for _, word := range words {
		go func(w string) {
			ch <- w
		}(word)
	}

	var concatenated string
	for i := 0; i < len(words); i++ {
		concatenated += <-ch + " "
	}

	close(ch)

	fmt.Println(concatenated)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	// Defer
	useDefer()

	// Exit(code int)
	fmt.Println("Exit")
	os.Exit(1)
}

func useDefer() {
	defer fmt.Println("Hello")
	fmt.Println("World")
}

package main

import "fmt"

func main() {
	var firstNumber int = 10
	var pointerFirst *int = &firstNumber
	var pointerOfPointer *int = pointerFirst

	fmt.Println(firstNumber, pointerFirst, pointerOfPointer)

	*pointerFirst = 12

	fmt.Println(firstNumber, pointerFirst, pointerOfPointer)
}

func greet(name string, age int8) {
	fmt.Printf("My name is %s and I'm %d years old", name, age)
}

func add(a, b int) int {
	return a + b
}

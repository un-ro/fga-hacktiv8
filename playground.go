package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	role    string
}

type Teacher interface {
	Teach()
	Bladder() string
}

func (p Person) Teach() {
	fmt.Println(p.name, "is teaching")
}

func (p Person) Bladder() string {
	return "I'm full"
}

func main() {
	var firstName string

	_ = firstName

	fmt.Println("Hello World!")

	var numbers []int = []int{1, 2, 3, 4, 5}

	for _, number := range numbers {
		fmt.Print(number, " ")
	}

	unero := Person{name: "Unero", age: 20, address: "Jakarta", role: "Student"}
	unero.Teach()
	unero.Bladder()
}

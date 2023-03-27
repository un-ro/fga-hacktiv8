package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i, " ")
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("maxgood")
		case i%3 == 0:
			fmt.Print("max")
		case i%5 == 0:
			fmt.Print("good")
		}
		fmt.Println()
	}
}

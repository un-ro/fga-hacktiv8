package main

import "fmt"

// Sesion 1 - Challenge 1
func main() {
	// Variables
	i := 21
	j := true
	var k float64 = 123.456

	fmt.Printf("%v \n%T \n", i, i) // Int
	fmt.Printf("%% \n")            // String %
	fmt.Printf("%t \n", j)         // Bool true
	fmt.Printf("%c \n", 1071)      // Unicode Я (ya)
	fmt.Printf("%b \n", i)         // Base 2 (21)
	fmt.Printf("%o \n", 25)        // Base 8 (25)
	fmt.Printf("%x \n", 15)        // Base 16 (f)
	fmt.Printf("%X \n", 15)        // Base 16 (F)
	fmt.Printf("%U \n", 1071)      // Hex
	fmt.Printf("%.6f \n", k)       // 6 decimal places
	fmt.Printf("%e \n", k)         // Scientific notation
}

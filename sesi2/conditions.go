package main

import "fmt"

func main() {
	currentYear := 2024

	if age := currentYear - 1999; age < 17 {
		println("You are not eligible to vote")
	} else {
		println("You are eligible to vote")
	}

	examScore := 70

	switch examScore {
	case 100:
		fmt.Println("Perfect score")
	case 90:
		fmt.Println("Almost perfect score")
	case 80:
		fmt.Println("Good score")
	case 70:
		fmt.Print("Not bad. ")
		fallthrough
	default:
		fmt.Println("You need to study harder")
	}

	height := 170

	switch {
	case height >= 180:
		fmt.Println("You are tall")
	case height >= 160 && height < 180:
		fmt.Println("You are average height")
	default:
		fmt.Println("You are short")
	}
}

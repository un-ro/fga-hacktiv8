package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %d \n", i)
	}

	for j := 0; j <= 10; j++ {
		fmt.Printf("Nilai j = %d \n", j)

		if j == 5 {
			for index, value := range "САШАРВО" {
				fmt.Printf("character %U '%c' starts at byte position %d\n", value, value, index)
			}
		}
	}
}

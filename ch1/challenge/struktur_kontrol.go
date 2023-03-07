package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %d \n", i)

		if i == 4 {
			for j := 0; j <= 10; j++ {
				fmt.Printf("Nilai j = %d \n", j)
			}
		}
	}
}

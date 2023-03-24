package main

func main() {
	println("Demo for loop:")

	for i := 0; i < 5; i++ {
		println(i)
	}

	println("Demo for loop with continue and break:")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i > 8 {
			break
		}

		println(i)
	}

	println("Demo for loop with label:")

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}

			println(i, j)
		}
	}
}

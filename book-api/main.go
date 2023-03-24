package main

import "book-api/routers"

const PORT string = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}

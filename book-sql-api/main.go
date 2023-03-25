package main

import "book-api/routers"

const PORT string = ":1337"

func main() {
	routers.StartServer().Run(PORT)
}

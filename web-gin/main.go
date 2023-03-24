package main

import "web-gin/routers"

const PORT string = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}

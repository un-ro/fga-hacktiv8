package main

import (
	"web-swagger/database"
	"web-swagger/routers"
)

const PORT string = ":1337"

func main() {
	database.StartDB()
	routers.StartServer().Run(PORT)
}

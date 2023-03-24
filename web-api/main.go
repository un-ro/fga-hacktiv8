package main

import (
	"fmt"
	"net/http"
)

const PORT string = ":6969"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World! (from Go)")
	})

	fmt.Println("Listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

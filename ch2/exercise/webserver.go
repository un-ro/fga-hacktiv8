package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const PORT string = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "System Shutdown!")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	})

	http.ListenAndServe(PORT, nil)
}

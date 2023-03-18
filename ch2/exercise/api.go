package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "John", Age: 30, Division: "IT"},
	{ID: 2, Name: "Doe", Age: 25, Division: "HR"},
	{ID: 3, Name: "Jane", Age: 27, Division: "IT"},
}

const PORT string = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employees", createEmployee)

	fmt.Println("Server is running on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertedAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertedAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
}

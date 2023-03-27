package main

import (
	"go-auth/middleware"
	"go-auth/models"

	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("Server is running on port 9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(w, r) {
		return
	}
	if !middleware.AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, models.SelectStudent(id))
		return
	}

	OutputJSON(w, models.GetStudents())
}

func Auth(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

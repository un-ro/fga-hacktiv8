package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "fga"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	psqlInfo := fmt.Sprintf("postgres://postgres:%s@%s/%s?sslmode=disable",
		user, host, dbname)

	db, err = sql.Open("postgres", psqlInfo) // Validate the connection

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping() // Check if the connection is alive
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//CreateEmployee()
	GetEmployees()
	//UpdateEmployee()
	DeleteEmployee()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`

	err = db.QueryRow(sqlStatement, "Unero", "ramadhanbhg@gmail.com", 23, "IT").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data: %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees ORDER BY id ASC`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee Employee

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employees Data: ", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2
	WHERE id = $1
	`

	result, err := db.Exec(sqlStatement, 1, "Unero Ramadhan")

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Updated Employee Data: ", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE FROM employees
	WHERE id = $1
	`

	result, err := db.Exec(sqlStatement, 2)

	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted Employee Data: ", count)
}

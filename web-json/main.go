package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
	{
		"full_name": "John Wick",
		"email": "john@mail.com",
		"age": 27
	}
	`

	var jsonStringArray = `
	[
		{
			"full_name": "John Wick",
			"email": "john@gmail.com",
			"age": 27
		},
		{
			"full_name": "Ethan Hunt",
			"email": "ethan@gmail.com",
			"age": 32
		}
	]
	`

	fmt.Println("Decode JSON by Struct")
	decodeJSONByStruct(jsonString)

	fmt.Println("\nDecode JSON by Map")
	decodeJSONByMap(jsonString)

	fmt.Println("\nDecode JSON by Interface")
	decodeJSONByInterface(jsonString)

	fmt.Println("\nDecode Array JSON")
	decodeArrayJSON(jsonStringArray)

	fmt.Println("\nEncode JSON")
	encodeJSON()

	fmt.Println("\nURL Parsing")
	urlParsing()
}

// Best solution if you know the JSON structure
func decodeJSONByStruct(jsonString string) {
	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Full Name:", result.FullName)
	fmt.Println("Email:", result.Email)
	fmt.Println("Age:", result.Age)
}

// Also good solution if you don't know the JSON structure
func decodeJSONByMap(jsonString string) {
	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Full Name:", result["full_name"])
	fmt.Println("Email:", result["email"])
	fmt.Println("Age:", result["age"])
}

// Most flexible way to decode JSON, but need to cast the result
func decodeJSONByInterface(jsonString string) {
	var result interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var employee = result.(map[string]interface{})

	fmt.Println("Full Name:", employee["full_name"])
	fmt.Println("Email:", employee["email"])
	fmt.Println("Age:", employee["age"])
}

// Solution for array of JSON
func decodeArrayJSON(jsonString string) {
	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, each := range result {
		fmt.Println("Employee ID", i)
		fmt.Println("Full Name:", each.FullName)
		fmt.Println("Email:", each.Email)
		fmt.Println("Age:", each.Age)
	}
}

// Decode solution also works for encode (from struct to JSON) or vice versa
func encodeJSON() {
	var object = []Employee{
		{"John Wick", "john@gmail.com", 27},
		{"Ethan Hunt", "ethan@gmail.com", 32},
	}

	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}

func urlParsing() {
	var urlString = "http://developer.com:80/hello?name=John&age=27"
	var u, err = url.Parse(urlString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("scheme: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)
}

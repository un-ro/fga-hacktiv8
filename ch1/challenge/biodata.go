package main

import (
	"fmt"
	"os"
)

// Session 4
type Biodata struct {
	Name, Address, Job, Motive string
}

func showBiodata(biodata Biodata) {
	fmt.Println("Name: ", biodata.Name)
	fmt.Println("Address: ", biodata.Address)
	fmt.Println("Job: ", biodata.Job)
	fmt.Println("Motive: ", biodata.Motive)
}

func main() {
	var args = os.Args[1:][0]

	dict := map[string]Biodata{
		"1": {"Rizky", "Jakarta", "Programmer", "Learn Go"},
		"2": {"Rizky", "Jakarta", "Programmer", "Learn Go"},
		"3": {"Rizky", "Jakarta", "Programmer", "Learn Go"},
		"4": {"Rizky", "Jakarta", "Programmer", "Learn Go"},
	}

	showBiodata(dict[args])
}

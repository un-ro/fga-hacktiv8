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
	fmt.Println("Name\t: ", biodata.Name)
	fmt.Println("Address\t: ", biodata.Address)
	fmt.Println("Job\t: ", biodata.Job)
	fmt.Println("Motive\t: ", biodata.Motive)
}

func main() {
	var args = os.Args[1:][0]

	data := map[string]Biodata{
		"1": {"Rizky", "Jakarta", "Programmer", "Learn Go"},
		"2": {"Nero", "Malang", "Programmer", "Learn Go"},
		"3": {"Dimas", "Bandung", "Programmer", "Learn Go"},
		"4": {"Fadya", "Solo", "Programmer", "Learn Go"},
	}

	showBiodata(data[args])
}

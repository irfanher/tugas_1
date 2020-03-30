package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	job       string
	salary    string
}

func main() {
	p1 := Employee{
		firstName: "Budi",
		lastName:  "Pratama",
		job:       "Marketing",
		salary:    "2.000.000",
	}
	fmt.Println("Employee ", p1)
}

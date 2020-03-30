package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	job       string
	salary    string
}

func (s Employee) new() {
	fmt.Println("FirstName : ", s.firstName)
	fmt.Println("LastName : ", s.lastName)
	fmt.Println("Job : ", s.job)
	fmt.Println("Salary : ", s.salary)

}

func main() {
	var s1 = Employee{"Budi", "Pratama", "Marketing", "2.000.000"}
	s1.new()

}

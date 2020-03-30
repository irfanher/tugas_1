package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	job       string
	salary    string
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}
func (p *Person) SetLastName(lastName string) {
	p.lastName = lastName
}
func (p *Person) SetJob(job string) {
	p.job = job
}
func (p *Person) SetSalary(salary string) {
	p.salary = salary
}

func (p Person) GetFirstName() string {
	return p.firstName
}
func (p Person) GetlastName() string {
	return p.lastName
}
func (p Person) GetJob() string {
	return p.job
}
func (p Person) GetSalary() string {
	return p.salary
}

func main() {

	employee := Person{}

	employee.SetFirstName("Budi")
	employee.SetLastName("Pratama")
	employee.SetJob("Marketing")
	employee.SetSalary("Rp. 2.000.000")

	fmt.Println("FirstName : ", employee.GetFirstName())
	fmt.Println("LastName : ", employee.GetlastName())
	fmt.Println("Job : ", employee.GetJob())
	fmt.Println("Salary : ", employee.GetSalary())

}

package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	e1 := Employee{"Albert", "Einstein", 1}

	e2 := Employee{
		id:        2,
		firstName: "Benjamin",
		lastName:  "Franklin",
	}

	var e3 Employee
	e3.id = 3
	e3.firstName = "Charles"
	e3.lastName = "Darwin"

	var employees = []Employee{e1, e2, e3}
	fmt.Printf("Employees: %v\n", employees)
}

package main

import (
	"fmt"
	"time"
)

const (
	numElements = 10000000
	firstName   = "John"
	lastName    = "Doe"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Start time
	start := time.Now()

	var persons = []Person{} //better; make([]Person, 0, numElements)
	for i := 0; i < numElements; i++ {
		persons = append(persons, Person{firstName, lastName, 1})
	}

	fmt.Printf("Persons: %d; Time: %d\n", len(persons), time.Since(start).Milliseconds())
}

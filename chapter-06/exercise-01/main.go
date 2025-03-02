package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	// Create a person using MakePerson
	p1 := MakePerson("John", "Doe", 30)
	fmt.Printf("Person 1: %v\n", p1)

	// Create a person using MakePersonPointer
	p2 := MakePersonPointer("Jane", "Doe", 25)
	fmt.Printf("Person 2: %v\n", p2)
}

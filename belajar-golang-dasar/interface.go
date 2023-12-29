package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHi(value HasName) {
	fmt.Println("Hi", value.GetName())
}

type Person struct {
	Name string
	Age  int
}

func (person Person) GetName() string {
	return person.Name
}

func main() {

	person := Person{Name: "Rizqi", Age: 32}
	sayHi(person)
}

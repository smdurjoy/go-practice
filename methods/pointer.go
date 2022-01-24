package main

import "fmt"

type Employee struct {
	name string
	age int
}

func (e *Employee) changeName(newName string) {
	e.name = newName
}

func main() {
	e := Employee {
		name: "Evan You",
		age: 34,
	}

	fmt.Printf("Employee old name: %s \n", e.name)

	e.changeName("Tylor Otwell")

	fmt.Printf("Employee new name: %s", e.name)
}
package main

import "fmt"

type Student struct {
	name string
	roll int
}

func (s Student) studentInfo() {
	fmt.Printf("Student name is %s and roll is %d", s.name, s.roll)
}

func main() {
	s := Student {
		name: "durJoy",
		roll: 02,
	}

	s.studentInfo()
}
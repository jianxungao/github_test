package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type agent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func (a agent) speak() {
	fmt.Println("I am", a.first, a.last, " - the agent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func main() {
	agent007 := agent{
		person: person{
			"james",
			"bone",
		},
		ltk: true,
	}

	person1 := person{
		"Mr:",
		"tommorrow",
	}

	bar(agent007)
	bar(person1)

}

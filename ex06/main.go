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

//interface any type has speak method, it is human
type human interface {
	speak()
}

// polymorphism
func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

// method of agent
func (a agent) speak() {
	fmt.Println("I am", a.first, a.last, " - the agent speak")
}

// method of person
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

func main() {
	agent007 := agent{
		person: person{ //composite literal
			"james",
			"bone",
		},
		ltk: true,
	}

	person1 := person{ //composite literal
		"Mr:",
		"tommorrow",
	}

	bar(agent007)
	bar(person1)

}

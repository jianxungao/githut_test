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

//interface: any type has speak method, it's human
type human interface {
	speak()
}

// polymorphism
func bar(h human) {
	switch h.(type) { // asserting what type it is
	case agent:
		fmt.Println(" - agent - ", h.(agent).first)
	case person:
		fmt.Println(" - person -", h.(person).first)
	default:
		fmt.Println(" : - ")
	}
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

// function return a function
func rf() func() int {
	return func() int {
		return 45
	}
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

//callback - pass a func as argument
func even(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
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

	func(x int) {
		fmt.Println("anonymous func with para", x)
	}(42)

	f := func(x int) {
		fmt.Println("func expression", x)
	}
	f(42)

	fmt.Println(rf()()) //call the function which return anoter func and run it

	ii := []int{2, 5, 6, 8, 9}          //slice of int
	fmt.Println("the sum:", sum(ii...)) //unfold

	s := even(sum, ii...) // callback func accept a func as argument
	fmt.Println("the sum of even number:", s)

}

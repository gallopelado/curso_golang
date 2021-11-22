package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	phone     string
	document  string
}

func main() {
	// 1 forma
	p1 := person{
		firstName: "Juan Jose",
		lastName:  "Gonzalez Ramirez",
		age:       29,
		phone:     "123456",
		document:  "1111111",
	}
	fmt.Println(p1)

	// 2 forma
	var p2 person
	p2.firstName = "Sandra"
	p2.lastName = "Lopez"
	p2.age = 30
	p2.phone = "11223366"
	p2.document = "99643183"

	fmt.Println(p2)
}

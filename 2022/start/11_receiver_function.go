package main

import (
	"fmt"
)

var print = fmt.Println

type Employee struct {
	firstName, lastName string
}

/*
receiver function
e Employee : e is call by value not reference
*/
func (e Employee) fullName() string {
	fullName := e.firstName + " " + e.lastName
	return fullName
}

/*
function receiver by reference
e *Employee : e is call  reference
*/
func (e *Employee) fullName2() string {
	e.firstName = "Subhadip"
	e.lastName = "Pahari"
	fullName := e.firstName + " " + e.lastName
	return fullName
}

func main() {
	e := Employee{
		firstName: "subhadip",
		lastName:  "pahari",
	}
	print(e.fullName())
	print(e.fullName2())
	print(e.fullName())
}

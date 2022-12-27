package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	type Person struct {
		firstName string
		lastName  string
		age       int
		address   []string
	}

	p1 := Person{
		firstName: "subhadip",
		lastName:  "pahari",
		age:       30,
		address:   []string{"math", "physics", "cs"},
	}

	print(p1)

	// promotion in struct
	type Student struct {
		person Person
		rollNo int
	}

	s1 := Student{
		person: Person{
			firstName: "subhadip",
			lastName:  "pahari",
			age:       30,
			address:   []string{"fullerton", "california"},
		},
		rollNo: 17,
	}

	print(s1)
	print(s1.person.firstName)

}

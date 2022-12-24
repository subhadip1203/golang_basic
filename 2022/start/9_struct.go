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
		subjects  []string
	}

	p1 := Person{
		firstName: "subhadip",
		lastName:  "pahari",
		age:       30,
		subjects:  []string{"math", "physics", "cs"},
	}

	print(p1)
}

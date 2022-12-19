package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	print(a[0])
	print(a[1])

	// just initialization
	var b [5]int
	b[1] = 10
	print(b)

	// if we do not provide any length of the variable,
	// it will be slice

	c := make([]int, 3)
	c[0] = 20
	print(c)
	c = append(c, 10)
	print(c)
}

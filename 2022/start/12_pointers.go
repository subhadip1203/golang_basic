package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	var x int = 20
	var y int = 30

	z1 := addOnly(x, y)
	print(x, y, z1)

	z2 := addAndChange(&x, &y)
	print(x, y, z2)

}

func addOnly(num1 int, num2 int) int {
	num1 += 1
	num2 += 1
	return num1 + num2
}

func addAndChange(num1 *int, num2 *int) int {
	*num1 += 1
	*num2 += 1
	return *num1 + *num2
}

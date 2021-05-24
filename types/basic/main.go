package main

import "fmt"

func main() {

	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)

	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	a := 20.45
	b := 34.89

	// Subtraction of two
	// floating-point number
	c := b - a

	// Display the result
	fmt.Printf("Result is: %f \n", c)
	// Display the type of c variable
	fmt.Printf("The type of c is : %T \n", c)

	// variables
	str1 := "GeeksforGeeks"
	str2 := "geeksForgeeks"
	str3 := "GeeksforGeeks"
	result1 := str1 == str2
	result2 := str1 == str3

	// Display the result
	fmt.Println(result1)
	fmt.Println(result2)
}

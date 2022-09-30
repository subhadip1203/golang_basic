package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

/*
	A variable of an interface type can hold a value of the Type
	that implements the interface.
*/

func main() {
	var s Shape
	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)
}

/*
	From the above result, we can see that zero value and type of the interface is nil.
	This is because, at this moment, the interface has no idea who is implementing it.
*/

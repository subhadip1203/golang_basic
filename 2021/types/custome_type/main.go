package main

import "fmt"

type Num int
type Fl float32
type person struct {
	name string
	age  int
}

func main() {

	var n1 Num = 7
	n2 := Num(42)
	fmt.Println(n1, n2)

	var f1 Fl = 7.2
	f2 := Fl(42.6)
	fmt.Println(f1, f2)

	p := person{name: "Francesco", age: 41}
	fmt.Println(p)
}

package main

import "fmt"

// taking a struct
type Rect struct {
	len, wid int
}

func (re Rect) Area() int {
	return re.len * re.wid
}

func main() {
	r := Rect{10, 12}
	fmt.Println("Length and Width are:", r)
	fmt.Println("Area of Rectangle: ", r.Area())
}

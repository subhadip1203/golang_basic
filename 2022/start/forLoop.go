package main

import (
	"fmt"
)

var print = fmt.Println

func main() {

	// initialization , end , condition
	for x := 0; x < 5; x++ {
		print(x)
	}
}

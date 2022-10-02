package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var print = fmt.Println

func main() {
	x := 11.35
	y := int(x)
	print(y)

	a := "10000"
	b, err := strconv.Atoi(a) // ascii to integer
	if err != nil {
		print("error")
	} else {
		print(b, reflect.TypeOf(b))
	}

	c := 12345
	d := strconv.Itoa(c)
	print(d, reflect.TypeOf(d))

	e := "1234.456"
	f, err := strconv.ParseFloat(e, 64) // float 64bit
	if err != nil {
		print("error in float conversion")
	} else {
		print(f, reflect.TypeOf(f))
	}

	g := 134.234
	myStr := fmt.Sprintf("%f", g) // %f means float type
	print(myStr, reflect.TypeOf(myStr))
}

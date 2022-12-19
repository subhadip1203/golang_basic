package main

import (
	"fmt"
	"reflect"
)

var print = fmt.Println

func main() {
	vName := 2
	vName2 := 4.67
	print(vName, reflect.TypeOf(vName))
	print(vName2, reflect.TypeOf(vName2))
}

package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	//  map[key_type] value_type
	var sampleMap = map[string]int{
		"one": 1,
		"two": 2,
	}
	print(sampleMap)
	print(sampleMap["one"])

	var anotherMap = map[string]string{
		"name":   "subhadip pahari",
		"father": "Asit pahari",
	}
	print(anotherMap)

	for key, value := range anotherMap {
		print(key, "::", value)
	}
}

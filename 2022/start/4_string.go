package main

import (
	"fmt"
	"strings"
)

var print = fmt.Println

func main() {
	var vName string = "subhadip"
	var vName2 = "subhadip"
	vName3 := "subahdip"

	print(vName, vName2, vName3)

	sV1 := "a monster"
	sV2 := strings.NewReplacer("a", "another").Replace(sV1)
	print(sV2)
	print("Length of string :", len(sV2))

	print("contains another:", strings.Contains(sV2, "another"))

	print("Index of o :", strings.Index(sV2, "o"))

	print("Replace char o :", strings.Replace(sV2, "o", "O", 1))  // single time reolace , so 1
	print("Replace char o :", strings.Replace(sV2, "o", "O", 2))  // single time reolace , so 2
	print("Replace char s :", strings.Replace(sV2, "e", "E", -1)) // all  reolace , so -1

}

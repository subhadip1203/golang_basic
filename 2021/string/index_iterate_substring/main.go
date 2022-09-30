package main

import "fmt"

func main() {
	s := "Japan 日本"

	for i, ch := range s {
		fmt.Printf("%d:%q ", i, ch)
	}
	fmt.Printf("\n")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%q ", s[i])
	}
	fmt.Printf("\n")
}

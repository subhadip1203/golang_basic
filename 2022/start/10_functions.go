package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	// Anonymous functions do not have a name.
	func(){
		print("abc")
	}()
	// Defer functions 
	defer runLast()

	// call function
	print(oneResult())
	print(multiResult())
}

func oneResult() string {
	return "ok"
}

func multiResult() (string, string) {
	return "ok", "another"
}

func runLast() {
	print("run at last for => Defer")
}


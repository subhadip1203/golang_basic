package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	defer runLast()
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

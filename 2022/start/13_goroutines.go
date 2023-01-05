package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func printHi() {
	print("Hi 1")
}

func main() {
	// the key word go used for go routines
	go printHi()

	print("Hi 2")
	time.Sleep(time.Second)
}

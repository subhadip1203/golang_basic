package main

import (
	"fmt"
)

var print = fmt.Println

func main() {
	values := make(chan int) // no buffer channel
	defer close(values)

	//unbufferd channel is like "Queue" of 0 size
	// we have to read and write at the same time in different go-routines
	//  we should read and write the channel at the same time
	// or it will be blocked

	// deadlock , as we trying to read but no one is writing
	// firstValue := <-values

	// making a go rutine function to write the function
	go func() {
		// witing the channel
		values <- 10
	}()

	// reading the channel
	// if we put it before the go func , it will be a dead lock
	firstValue := <-values
	print(firstValue)

}

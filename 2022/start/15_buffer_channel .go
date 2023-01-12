package main

import (
	"fmt"
	"math/rand"
	"time"
)

var print = fmt.Println

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func foo(c chan string) {
	print("go routin started")
	time.Sleep(time.Second * 1)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c <- RandStringBytes(5)
	print("go routines ended")

}

func main() {
	values := make(chan string, 2) // buffer channel
	// values := make(chan string) // no buffer channel
	defer close(values)

	go foo(values)
	go foo(values)

	// this only get the first foo function go channe value
	firstValue := <-values
	print(firstValue)

	time.Sleep(time.Second * 2)
}

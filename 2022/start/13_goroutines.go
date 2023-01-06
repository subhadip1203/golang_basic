package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var print = fmt.Println

func printHi(s string) {
	for i := 0; i < 3; i++ {
		print(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	// the key word go used for go routines
	wg.Add(1)
	go printHi("hi")
	wg.Add(1)
	go printHi("hello")
	wg.Wait()
	wg.Add(1)
	go printHi("subhadip")
	wg.Wait()

}

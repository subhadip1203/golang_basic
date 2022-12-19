package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func main() {
	now := time.Now()
	print(now.Year(), now.Month(), now.Day())
	print(now.Hour(), now.Minute(), now.Second())
}

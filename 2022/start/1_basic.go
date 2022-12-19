package main

import (
	"bufio"
	"fmt"
	"os"
)

var print = fmt.Println

func main() {
	num := 10
	print(num)
	print("Enter some string : ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		print(err)
	}

	print(name)
}

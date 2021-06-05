package main

import (
	"fmt"
	"strings"
)

func main() {
	status1 := strings.Contains("Japan", "abc")
	status2 := strings.Contains("Japan", "apa")
	fmt.Println(status1)
	fmt.Println(status2)
}

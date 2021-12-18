package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Enter a string:")
	fmt.Scan(&s)
	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
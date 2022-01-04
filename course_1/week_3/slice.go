package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	sli := make([]int, 0, 3)

	var value string

	fmt.Println("Enter an integer: ")
	for {
		fmt.Scan(&value)
		strings.ToLower(value)

		if value == "x" {
			fmt.Println("Exit")
			break
		}

		parsedInt, _ := strconv.Atoi(value)
		sli = append(sli, parsedInt)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}

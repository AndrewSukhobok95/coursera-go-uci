package main

import (
	"fmt"
	"strconv"
)


func Swap(intSlice []int, index int) {
	tmp := intSlice[index]
	intSlice[index] = intSlice[index + 1]
	intSlice[index + 1] = tmp
}


func BubbleSort(intSlice []int) {
	for {
		swapped := false
		for i:=0; i<len(intSlice)-1; i++ {
			if intSlice[i] > intSlice[i+1] {
				Swap(intSlice, i)
				swapped = true
			}
		}
		if !swapped { break }
	}
}


func main() {
	var numIters int = 10
	sli := make([]int, numIters)

	var newNum string
	for i:=0; i<numIters; i++ {
		fmt.Printf("Enter a number (%d out of %d):\n", i+1, numIters)
		_, err := fmt.Scan(&newNum)
		if err != nil { fmt.Println(err) }
		sli[i], err = strconv.Atoi(newNum)
	}

	BubbleSort(sli)
	fmt.Println("Here is the sorted slice of entered numbers:")
	fmt.Println(sli)
}

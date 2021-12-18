package main

import (
	"fmt"
	"strconv"
)

func main() {
	var floatDigit float64
	fmt.Printf("Enter a floating point digit:")
	fmt.Scan(&floatDigit)
	intDigit := int64(floatDigit)
	fmt.Printf(strconv.FormatInt(intDigit, 10))
}
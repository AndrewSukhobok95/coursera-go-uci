package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return 0.5 * a * math.Pow(t, 2) + v0 * t + s0
	}
	return fn
}

func main() {
	var a float64
	var v0 float64
	var s0 float64
	var t float64

	fmt.Println("Enter acceleration value:")
	fmt.Scan(&a)
	fmt.Println("Enter initial velocity value:")
	fmt.Scan(&v0)
	fmt.Println("Enter initial displacement value:")
	fmt.Scan(&s0)
	fmt.Println("Enter time value:")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("The displacement for the given parameters: ", fn(t))
}

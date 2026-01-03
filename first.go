package main

import (
	"fmt"
	"math/cmplx"
)

// CONCEPT: Packaged Variables & Basic Types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// CONCEPT: Functions (Omitting type for consecutive args)
func add(x, y int) int {
	return x + y
}

// CONCEPT: Multiple Results
func swap(x, y string) (string, string) {
	return y, x
}

// CONCEPT: Named Return Values ("Naked" return)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(">> DAY 2: TOUR OF GO CONCEPTS")

	// CONCEPT: Short Variable Declarations (Inside function only)
	k := 3
	c, python, java := true, false, "no!"

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("Variables:", k, c, python, java)

	fmt.Println("Add Function:", add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println("Swap Function:", a, b)

	x, y := split(17)
	fmt.Println("Split Function (Named Returns):", x, y)
}
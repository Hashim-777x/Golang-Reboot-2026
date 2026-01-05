package main

import "fmt"

func main() {
	fmt.Println(">> DAY 5: MEMORY & FIXED STORAGE")

	// 1. POINTERS: The "Address" vs The "Value"
	i, j := 42, 2701

	p := &i                                    // point to i (p now holds the memory address of i)
	fmt.Println("Value of i via pointer:", *p) // read i through the pointer (*dereferencing)

	*p = 21 // set i through the pointer
	fmt.Println("New value of i:", i)

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println("New value of j:", j)

	// 2. ARRAYS: The Fixed-Length Foundation

	var a [2]string
	a[0] = "Hello"
	a[1] = "Go"
	fmt.Println("Array a:", a[0], a[1])
	fmt.Println("Full Array a:", a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // Array literal
	fmt.Println("Primes Array:", primes)
}

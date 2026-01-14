package main

import "fmt"

func main() {
	fmt.Println(">> DAY 4: THE POWER OF SLICES")

	// 1. SLICES ARE WINDOWS TO ARRAYS
	names := [4]string{"Hashim", "Abdullah", "Shaad", "Ahmed"} // An Array
	fmt.Println("Original Array:", names)

	a := names[0:2] // Slice: includes index 0, 1. Excludes 2.
	b := names[1:3] // Slice: includes index 1, 2. Excludes 3.
	fmt.Println("Slice A:", a)
	fmt.Println("Slice B:", b)

	// 2. CHANGING A SLICE CHANGES THE ARRAY
	// Slices are just headers pointing to the same memory.
	b[0] = "XXX"
	fmt.Println("After modifying Slice B, Original Array is:", names)

	// 3. LENGTH vs CAPACITY
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extending length.
	s = s[:4]
	printSlice(s)

	// 4. CREATING SLICES WITH make
	//creating dynamic array
	m := make([]int, 5) // len=5, cap=5
	fmt.Printf("Make Slice: len=%d cap=%d %v\n", len(m), cap(m), m)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

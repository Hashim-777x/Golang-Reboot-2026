package main

import "fmt"

func main() {
	fmt.Println(">> DAY 7: THE COMEBACK (MAPS & CLOSURES)")

	// 1. MAPS: Key-Value logic (The real workhorse)
	// Like JS object but strict types.
	m := make(map[string]string)
	m["Role"] = "Backend Engineer"
	m["Status"] = "Relentless"
	
	// Check if a key exists
	val, ok := m["Status"]
	fmt.Printf("Status: %v (Exists: %v)\n", val, ok)

	// 2. FUNCTION CLOSURES
	// Go functions can be values & remember their env.
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println("Positive:", pos(i), " | Negative:", neg(-2*i))
	}
}

// adder returns a function that 'closes over' the variable sum.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

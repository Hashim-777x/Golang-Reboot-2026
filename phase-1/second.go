package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(">> DAY 2: TOUR OF GO CONCEPTS:- FLOW CONTROL")

	// 1. THE ONLY LOOP: for
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum 0-9:", sum)

	// 2. THE "WHILE" LOOP 
	count := 1
	for count < 5 {
		count += count
	}
	fmt.Println("While-style result:", count)

	// 3. IF WITH A SHORT STATEMENT
	// we can execute a statement before the condition
	if v := math.Pow(3, 2); v < 10 {
		fmt.Println("3 squared is less than 10:", v)
	}

	// 4. SWITCH (No 'break' in golang)
	fmt.Print("Go runs on: ")
	os := "linux" // Hardcoded for example
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("Something else.")
	}
}



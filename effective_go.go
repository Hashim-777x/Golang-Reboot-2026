package main

import (
	"fmt"
)

// 1. ALLOCATION: new vs make
// 'new' allocates memory but doesn't initialize (returns a pointer to zero value).
// 'make' initializes internal data structures (slices, maps, channels only).

type DatabaseConfig struct {
	Host string
	Port int
}

func main() {
	fmt.Println(">> DAY 10: EFFECTIVE GO PATTERNS")

	// Pattern A: The Composite Literal (Preferred over new)
	// We use this to keep the code clean and set values immediately.
	db := &DatabaseConfig{
		Host: "localhost",
		Port: 5432,
	}
	fmt.Printf("Config: %+v\n", db)

	// Pattern B: The 'Make' Rule
	// If we don't use make for maps, our app crashes on assignment.
	// We must initialize the internal data structure first.
	scores := make(map[string]int) 
	scores["Hashim"] = 100

	// Pattern C: Comma-OK Idiom (The idiomatic Go error handling)
	// We use this to check for a key's existence without causing a panic.
	if val, ok := scores["Unknown"]; !ok {
		fmt.Println("Key not found - we are handling this gracefully.")
	} else {
		fmt.Println("Value found:", val)
	}

	// Pattern D: Defer (The Go way to clean up resources)
	// We use this to ensure a function call runs at the very end of the current function.
	// It's how we guarantee cleanup, like closing files or DB connections.
	defer fmt.Println(">> Process complete. We have cleared all resources.")
	
	doWork()
}

func doWork() {
	fmt.Println("Working on the scrap PC rig...")
}

package main

import (
	"fmt"
	"time"
)

// 1. THE INTERFACE (The Blueprint)
type Logger interface {
	Log(message string)
}

// 2. THE STRUCT (The Implementation)
type ConsoleLogger struct {
	Source string
}

// ConsoleLogger implicitly implements Logger because it has the Log method.
func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("[%v] %s: %s\n", time.Now().Format("15:04:05"), cl.Source, message)
}

// 3. ERROR HANDLING (The Go Way)
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work on the scrap PC",
	}
}

func main() {
	fmt.Println(">> DAY 7: INTERFACES & ERROR HANDLING")

	// Interface in action
	var l Logger = ConsoleLogger{Source: "System_Kernel"}
	l.Log("Initializing Go Architect sequence...")

	// Error handling in action
	if err := run(); err != nil {
		fmt.Println("Caught an error:", err)
	}
}

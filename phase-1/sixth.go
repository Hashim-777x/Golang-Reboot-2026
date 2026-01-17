package main

import "fmt"

// 1. THE DATA (Struct)
type Developer struct {
	Name string
	Role string
}

// 2. THE METHOD (Value Receiver)
// This attaches 'Greet' specifically to the Developer struct.
func (d Developer) Greet() {
	fmt.Printf("Hi, I'm %s, a %s working on a scrap PC.\n", d.Name, d.Role)
}

// 3. THE STRINGER INTERFACE
// fmt.Println looks for this specific method to decide how to print struct.
func (d Developer) String() string {
	return fmt.Sprintf("%v (%v) ", d.Name, d.Role)
}

func main() {
	fmt.Println(">> DAY 6: METHODS & THE STRINGER INTERFACE")

	me := Developer{Name: "Hashim", Role: "Golang Architect"}

	// Calling the method
	me.Greet()

	// Testing the Stringer interface implementation
	fmt.Println("Custom Print:", me)
}


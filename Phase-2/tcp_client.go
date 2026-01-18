package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// 1. Establish connection with a timeout. 
	address := "localhost:6379"
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Printf(" CRITICAL ERROR: Could not connect to %s: %v\n", address, err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type 'exit' to quit.")

	// 2. Initialize readers ONCE. Efficiency matters.
	terminalReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		// 3. Get user input
		fmt.Print("👉 Message: ")
		input, err := terminalReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		// Clean the input
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println(" Shutting down...")
			break
		}

		// 4. Set a Deadline for the WRITE and READ
		// 
		conn.SetDeadline(time.Now().Add(2 * time.Second))

		// 5. Send data
		_, err = fmt.Fprintf(conn, input+"\n")
		if err != nil {
			fmt.Println(" Failed to send data:", err)
			break
		}

		// 6. Read response
		message, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println(" Server timeout or connection lost:", err)
			break
		}

		fmt.Printf(" Server replied: %s", message)
	}
}

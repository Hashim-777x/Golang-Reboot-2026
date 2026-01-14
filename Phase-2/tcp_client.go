package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. Connect to the server we built yesterday
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// 2. Read input from the terminal (your scrap PC keyboard)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')

	// 3. Send the bytes across the wire
	fmt.Fprintf(conn, text + "\n")

	// 4. Wait for the ACK (Acknowledgment) from the server
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Server replied: " + message)
}

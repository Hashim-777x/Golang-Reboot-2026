package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

const (
	serverAddr      = "localhost:6379"
	connectTimeout  = 5 * time.Second
	operationTimeout = 3 * time.Second
	bufferSize      = 4096
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, " FATAL: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// 1. Establish connection with explicit timeout
	conn, err := net.DialTimeout("tcp", serverAddr, connectTimeout)
	if err != nil {
		return fmt.Errorf("connection failed to %s: %w", serverAddr, err)
	}
	defer conn.Close()

	fmt.Printf(" Connected to %s\n", serverAddr)
	fmt.Println(" Commands: Type your message or 'exit' to quit\n")

	// 2. Initialize buffered readers with custom buffer size for performance
	terminalReader := bufio.NewReaderSize(os.Stdin, bufferSize)
	serverReader := bufio.NewReaderSize(conn, bufferSize)

	for {
		// 3. Prompt and read user input
		fmt.Print("----------- ")
		input, err := terminalReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\n Connection closed")
				return nil
			}
			return fmt.Errorf("input read failed: %w", err)
		}

		// 4. Sanitize and validate input
		input = strings.TrimSpace(input)
		
		if input == "" {
			continue // Skip empty inputs
		}
		
		if strings.ToLower(input) == "exit" {
			fmt.Println(" Shutting down gracefully...")
			return nil
		}

		// 5. Send command with timeout protection
		if err := sendCommand(conn, input); err != nil {
			return fmt.Errorf("send failed: %w", err)
		}

		// 6. Receive response with timeout protection
		response, err := receiveResponse(conn, serverReader)
		if err != nil {
			return fmt.Errorf("receive failed: %w", err)
		}

		fmt.Printf("⬅️  Server: %s\n", response)
	}
}

// sendCommand writes data to connection with deadline enforcement
func sendCommand(conn net.Conn, command string) error {
	conn.SetWriteDeadline(time.Now().Add(operationTimeout))
	
	// Use buffered writer for efficiency
	writer := bufio.NewWriter(conn)
	_, err := writer.WriteString(command + "\n")
	if err != nil {
		return err
	}
	
	return writer.Flush()
}

// receiveResponse reads server reply with deadline enforcement
func receiveResponse(conn net.Conn, reader *bufio.Reader) (string, error) {
	conn.SetReadDeadline(time.Now().Add(operationTimeout))
	
	response, err := reader.ReadString('\n')
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return "", fmt.Errorf("server response timeout")
		}
		return "", err
	}
	
	return strings.TrimSpace(response), nil
}

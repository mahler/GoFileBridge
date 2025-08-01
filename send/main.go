package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Check for filename argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	targetURL := "http://localhost:8182/" + filePath // Adjust target URL to use filename

	// Read the file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create PUT request
	req, err := http.NewRequest(http.MethodPut, targetURL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	// Set headers (e.g., Content-Type as plain text)
	req.Header.Set("Content-Type", "text/plain")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending PUT request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print response status
	fmt.Printf("Response Status: %s\n", resp.Status)

	// Optionally, print response body
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response Body:\n%s\n", string(respBody))
}

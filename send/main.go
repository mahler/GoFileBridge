package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// File path and target URL
	filePath := "example.txt"                    // Replace with your actual file path
	targetURL := "http://localhost:8182/hello.txt"  // Replace with your target URL

	// Read the file content
	data, err := ioutil.ReadFile(filePath)
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
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response Body:\n%s\n", string(respBody))
}
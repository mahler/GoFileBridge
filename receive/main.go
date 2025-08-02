package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/", handlePut)

	port := "8182"
	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

	if r.Method != http.MethodPut {
		fmt.Printf("Rejected request: method not allowed (%s)\n", r.Method)
		http.Error(w, "Only PUT method is supported", http.StatusMethodNotAllowed)
		return
	}

	originalFileName := filepath.Base(r.URL.Path)
	if originalFileName == "" || originalFileName == "/" {
		fmt.Println("Rejected request: file name missing in URL path")
		http.Error(w, "File name missing in URL path", http.StatusBadRequest)
		return
	}

	now := time.Now()
	dateTimePrefix := now.Format("2006-01-02-15-01-")
	fileName := "./files/" + dateTimePrefix + originalFileName

	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file '%s': %v\n", fileName, err)
		http.Error(w, fmt.Sprintf("Error creating file: %v", err), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	written, err := io.Copy(outFile, r.Body)
	if err != nil {
		fmt.Printf("Error writing to file '%s': %v\n", fileName, err)
		http.Error(w, fmt.Sprintf("Error writing to file: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("File saved: %s (%d bytes)\n", fileName, written)
	fmt.Fprintf(w, "File '%s' saved (%d bytes)\n", fileName, written)
}

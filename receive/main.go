package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Optional: Use the URL path as the filename (e.g., PUT /file.txt)
	fileName := filepath.Base(r.URL.Path)
	if fileName == "" || fileName == "/" {
		http.Error(w, "File name missing in URL path", http.StatusBadRequest)
		return
	}

	// Open file for writing
	outFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating file: %v", err), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Copy request body to file
	written, err := io.Copy(outFile, r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing to file: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File '%s' saved (%d bytes)\n", fileName, written)
}
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 \n"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate_random_text.go <size_in_MB>")
		return
	}

	// Parse argument
	mb, err := strconv.Atoi(os.Args[1])
	if err != nil || mb <= 0 {
		fmt.Println("Invalid size. Please provide a positive integer.")
		return
	}

	sizeInBytes := mb * 1024 * 1024
	filename := fmt.Sprintf("random_text_%dMB.txt", mb)

	rand.Seed(time.Now().UnixNano())

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < sizeInBytes; i++ {
		randomChar := charset[rand.Intn(len(charset))]
		if _, err := file.Write([]byte{randomChar}); err != nil {
			panic(err)
		}
	}

	fmt.Println("File generated:", filename)
}

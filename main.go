package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	
	// Use "strings.TrimSpace" to remove any whitespace
	text = strings.TrimSpace(text)
	fmt.Printf("First Pass: %s\n", text)

	// Use "strings.Fields" to split into an array of words.
	words := strings.Fields(text)
	fmt.Printf("Second Pass: %v\n", words)

	return words
}

func main() {
	input := "         Hello,           world! This         is a test.    "
	cleaned := cleanInput(input)

	fmt.Print(cleaned)
}
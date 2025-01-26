package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan(){
			text := scanner.Text()
			cleanText := cleanInput(text)
			if len(cleanText) > 0 {
				fmt.Printf("Your command was: %s\n", cleanText[0])
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
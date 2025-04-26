package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *config) {
	userInput := os.Stdin
	pokeScanner := bufio.NewScanner(userInput)
	supportedCommands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	for {

		fmt.Printf("Pokedex > ")
		pokeScanner.Scan()
		userText := string(pokeScanner.Text())
		splitText := cleanInput(userText)
		if len(splitText) == 0 {
			continue
		}
		if command, exists := supportedCommands[splitText[0]]; !exists {
			fmt.Println("Unknown command")
		} else {
			args := splitText[1:]
			err := command.callback(c, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}
}

func cleanInput(text string) []string {
	var cleaned []string
	text = strings.ToLower(text)
	cleaned = strings.Fields(text)
	return cleaned
}

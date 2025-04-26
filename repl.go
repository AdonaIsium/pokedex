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

		if command, exists := supportedCommands[splitText[0]]; !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(c)
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

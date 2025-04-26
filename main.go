package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands = make(map[string]cliCommand)

func cleanInput(text string) []string {
	var cleaned []string
	text = strings.ToLower(text)
	cleaned = strings.Fields(text)
	return cleaned
}

func commandExit() error {
	// fmt.Println("Thank you for visiting and have a wonderful day!")
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Usage:\n")
	for k, v := range supportedCommands {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func main() {
	exitCommand := cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	helpCommand := cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	supportedCommands["help"] = helpCommand
	supportedCommands["exit"] = exitCommand

	userInput := os.Stdin
	pokeScanner := bufio.NewScanner(userInput)
	fmt.Println("Welcome to the Pokedex!")
	for {

		fmt.Printf("Pokedex > ")
		pokeScanner.Scan()
		userText := string(pokeScanner.Text())
		splitText := cleanInput(userText)

		if command, exists := supportedCommands[splitText[0]]; !exists {
			fmt.Println("Unknown command")
		} else {
			command.callback()
		}
	}
}

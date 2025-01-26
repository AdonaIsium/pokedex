package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func() error
}

var commands map[string]cliCommand


func init(){
	commands = map[string]cliCommand{
		"exit": {
			name: 	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: 	"help",
			description: "Displays a help menu",
			callback: commandHelp,
		},
	}
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan(){
			text := scanner.Text()
			cleanText := cleanInput(text)
			currentCommand := cleanText[0]
			if cmd, exists := commands[currentCommand]; exists{
				cmd.callback()
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	getCommands()
	return nil
}

func getCommands() error {
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}




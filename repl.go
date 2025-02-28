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
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		currentCommand := text[0]
		cmd, exists := getCommands()[currentCommand]
		if exists{
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name 		string
	description string
	callback 	func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: 	"help",
			description: "Displays a help menu",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Display 20 Pokemon map locations",
			callback: commandMap,
		},
		"exit": {
			name: 	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},

	}
}




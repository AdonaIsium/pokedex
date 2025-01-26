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
			currentCommand := cleanText[0]
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
		"exit": {
			name: 	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}




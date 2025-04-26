package main

import "fmt"

func commandHelp(c *config) error {
	supportedCommands := getCommands()
	fmt.Printf("Usage:\n")
	for k, v := range supportedCommands {
		fmt.Printf("  %s: %s\n", k, v.description)
	}
	return nil
}

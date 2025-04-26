package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AdonaIsium/pokedex/internal/pokecache"
)

func startRepl(c *config) {
	userInput := os.Stdin
	pokeScanner := bufio.NewScanner(userInput)
	supportedCommands := getCommands()
	cache := pokecache.NewCache(time.Minute)
	fmt.Println("Welcome to the Pokedex!")
	for {

		fmt.Printf("Pokedex > ")
		pokeScanner.Scan()
		userText := string(pokeScanner.Text())
		splitText := cleanInput(userText)

		if command, exists := supportedCommands[splitText[0]]; !exists {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(c, cache)
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

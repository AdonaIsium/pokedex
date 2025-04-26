package main

import (
	"fmt"

	"github.com/AdonaIsium/pokedex/internal/pokecache"
)

func commandHelp(c *config, cache *pokecache.Cache) error {
	supportedCommands := getCommands()
	fmt.Printf("Usage:\n")
	for k, v := range supportedCommands {
		fmt.Printf("  %s: %s\n", k, v.description)
	}
	return nil
}

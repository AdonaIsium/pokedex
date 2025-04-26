package main

import (
	"fmt"
	"os"

	"github.com/AdonaIsium/pokedex/internal/pokecache"
)

func commandExit(c *config, cache *pokecache.Cache) error {
	// fmt.Println("Thank you for visiting and have a wonderful day!")
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

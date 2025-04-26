package main

import (
	"time"

	"github.com/AdonaIsium/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	c := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(c)
}

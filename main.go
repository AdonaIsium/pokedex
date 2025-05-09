package main

import (
	"time"

	"github.com/AdonaIsium/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute)
	c := &config{
		caughtPokemon: map[string]pokeapi.PokemonResp{},
		pokeapiClient: pokeClient,
	}

	startRepl(c)
}

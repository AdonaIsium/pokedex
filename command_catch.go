package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {

	pokemonName := args[0]
	pokemonResp, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...", pokemonName)
	caught := attemptCatch(pokemonResp.BaseExperience)
	if caught {
		fmt.Printf("Caught a wild %s\n", pokemonResp.Name)
		fmt.Printf("It has some cool abilities:\n")
		for i, ability := range pokemonResp.Abilities {
			fmt.Printf("Ability %d: %s\n", i+1, ability.Ability.Name)
		}
	} else {
		fmt.Printf("Awww, the %s was unimpressed by your catch attempt\n", pokemonResp.Name)
	}

	c.caughtPokemon[pokemonResp.Name] = pokemonResp

	return nil
}

func attemptCatch(pokemonExp int) bool {
	return rand.Intn(100) > (pokemonExp / 2)
}

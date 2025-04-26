package main

import "fmt"

func commandExplore(c *config, args ...string) error {
	firstArg := args[0]
	exploreResp, err := c.pokeapiClient.GetLocation(firstArg)
	if err != nil {
		return err
	}

	for _, pokemonEncounter := range exploreResp.PokemonEncounters {
		fmt.Printf("%s\n", pokemonEncounter.Pokemon.Name)
	}
	return nil
}

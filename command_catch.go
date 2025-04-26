package main

import "fmt"

func commandCatch(c *config, args ...string) error {

	firstArg := args[0]
	pokemonResp, err := c.pokeapiClient.GetPokemon(firstArg)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemonResp.Name)
	for i, ability := range pokemonResp.Abilities {
		fmt.Printf("Ability %d: %s\n", i+1, ability.Ability.Name)
	}

	return nil
}

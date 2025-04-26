package main

import "fmt"

func commandInspect(c *config, args ...string) error {

	pokemonName := args[0]
	pokemon, exists := c.caughtPokemon[pokemonName]
	if !exists {
		return fmt.Errorf("%s not found in pokedex", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  %s\n", pokemonType.Type.Name)
	}

	return nil
}

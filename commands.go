package main

import (
	"github.com/AdonaIsium/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func getCommands() map[string]cliCommand {
	var supportedCommands = make(map[string]cliCommand)
	mapCommand := cliCommand{
		name:        "map",
		description: "Get information about a location",
		callback:    commandMapForward,
	}
	mapBackCommand := cliCommand{
		name:        "mapb",
		description: "Get information about a location",
		callback:    commandMapBack,
	}
	exploreCommand := cliCommand{
		name:        "explore",
		description: "Learn about an areas Pokemon",
		callback:    commandExplore,
	}
	catchCommand := cliCommand{
		name:        "catch",
		description: "catch a pokemon!",
		callback:    commandCatch,
	}
	exitCommand := cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	helpCommand := cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	supportedCommands["map"] = mapCommand
	supportedCommands["mapb"] = mapBackCommand
	supportedCommands["explore"] = exploreCommand
	supportedCommands["catch"] = catchCommand
	supportedCommands["help"] = helpCommand
	supportedCommands["exit"] = exitCommand
	return supportedCommands
}

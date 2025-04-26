package main

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}

type Config struct {
	Next     string
	Previous string
}

func getCommands() map[string]cliCommand {
	var supportedCommands = make(map[string]cliCommand)
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
	mapCommand := cliCommand{
		name:        "map",
		description: "Get information about a location",
		callback:    commandMap,
	}
	mapBackCommand := cliCommand{
		name:        "mapb",
		description: "Get information about a location",
		callback:    commandMapBack,
	}

	supportedCommands["map"] = mapCommand
	supportedCommands["mapb"] = mapBackCommand
	supportedCommands["help"] = helpCommand
	supportedCommands["exit"] = exitCommand
	return supportedCommands
}

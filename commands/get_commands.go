package commands

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a Pokemon",
			Callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			Callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a caught pokemon",
			Callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			Callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			Callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught pokemon",
			Callback:    commandPokedex,
		},
	}
}

package commands

import (
	"github.com/snowkittyselene/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a pokemon",
			Callback:    commandCatch,
		},
		"explore": {
			Name:        "explore",
			Description: "Shows the Pokemon encoutnered in an area",
			Callback:    commandExplore,
		},
		"map": {
			Name:        "map",
			Description: "Shows the next twenty map locations",
			Callback:    commandMapForward,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous twenty map locations",
			Callback:    commandMapBack,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}

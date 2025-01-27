package commands

import (
	"github.com/snowkittyselene/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			Name:        "map",
			Description: "Shows the next twenty map locations",
			Callback:    commandMap,
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

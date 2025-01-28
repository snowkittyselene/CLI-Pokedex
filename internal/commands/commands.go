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
	Pokedex  map[string]pokeapi.PokemonInfo
	Next     *string
	Previous *string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			Name:        "pokedex",
			Description: "Display all caught Pokemon",
			Callback:    commandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect the stats of a caught Pokemon",
			Callback:    commandInspect,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a Pokemon",
			Callback:    commandCatch,
		},
		"explore": {
			Name:        "explore",
			Description: "Shows the Pokemon encountered in an area",
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

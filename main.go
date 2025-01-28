package main

import (
	"time"

	"github.com/snowkittyselene/commands"
	"github.com/snowkittyselene/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &commands.Config{
		Client:  client,
		Pokedex: map[string]pokeapi.PokemonInfo{},
	}
	startRepl(config)
}

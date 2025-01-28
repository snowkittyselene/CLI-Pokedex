package commands

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch requires a pokemon")
	}
	pokemon := args[0]
	pokemonInfo, err := c.Client.CallPokemonInfo(pokemon)
	if err != nil {
		return err
	}
	speciesInfo, err := c.Client.CallPokemonSpeciesInfo(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at  %s...\n", pokemonInfo.Name)
	if rand.Intn(255) <= speciesInfo.CaptureRate {
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
	}
	return nil
}

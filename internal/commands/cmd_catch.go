package commands

import "fmt"

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch requires a pokemon")
	}
	pokemon := args[0]
	pokemonInfo, err := c.Client.CallPokemonInfo(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("You searched for: %s\n", pokemonInfo.Name)
	return nil
}

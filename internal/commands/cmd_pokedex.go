package commands

import "fmt"

func commandPokedex(c *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for pokemon, _ := range c.Pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}

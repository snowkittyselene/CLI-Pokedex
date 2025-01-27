package commands

import "fmt"

func commandExplore(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore requires a location")
	}
	location := args[0]

	fmt.Printf("Exploring %s...\n", location)

	res, err := c.Client.CallLocationDetails(location)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s:\n", location)
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

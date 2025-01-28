package commands

import "fmt"

func commandInspect(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect requires a pokemon")
	}
	pokemonName := args[0]
	info, ok := c.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon yet")
	}
	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Height: %0.1f m\n", float64(info.Height)/10.0)
	fmt.Printf("Weight: %0.1f kg\n", float64(info.Weight)/10.0)
	fmt.Println("Stats:")
	for _, item := range info.Stats {
		fmt.Printf(" - %s: %d\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Println("Types:")
	for _, item := range info.Types {
		fmt.Printf(" - %s\n", item.Type.Name)
	}
	return nil
}

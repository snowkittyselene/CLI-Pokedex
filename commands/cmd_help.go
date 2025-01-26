package commands

import "fmt"

func commandHelp() error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

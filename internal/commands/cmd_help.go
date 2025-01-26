package commands

import (
	"fmt"
	"slices"
)

func commandHelp(c *Config) error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	cmdNames := []string{}
	for _, command := range commands {
		cmdNames = append(cmdNames, command.Name)
	}
	slices.Sort(cmdNames)
	for _, cmd := range cmdNames {
		fmt.Printf("%s: %s\n", commands[cmd].Name, commands[cmd].Description)
	}
	return nil
}

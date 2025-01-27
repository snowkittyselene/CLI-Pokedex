package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/snowkittyselene/commands"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	res := strings.Fields(lower)
	return res
}

func startRepl(config *commands.Config) {

	reader := bufio.NewScanner(os.Stdin)
	commands := commands.GetCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())
		// skip empty
		if len(input) == 0 {
			continue
		}
		cmd := input[0]
		res, ok := commands[cmd]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := res.Callback(config)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

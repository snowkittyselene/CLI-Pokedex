package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	res := strings.Fields(lower)
	return res
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		res := cleanInput(reader.Text())
		if len(res) == 0 {
			continue
		}
		cmd := res[0]
		fmt.Printf("Your command was: %s\n", cmd)
	}
}

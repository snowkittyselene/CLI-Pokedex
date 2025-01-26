package commands

const baseApiUrl string = "https://pokeapi.co/api/v2/"

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			Name:        "map",
			Description: "Shows the next twenty map locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous twenty map locations",
			Callback:    commandMapBack,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}

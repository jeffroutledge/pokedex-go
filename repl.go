package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jeffroutledge/CliPokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type params struct {
	name string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		text := cleanInput(reader.Text())
		if command, commandExists := cliCommands()[text[0]]; commandExists {
			params := text[1:]
			err := command.callback(cfg, params)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			handleInvalidCmd(text[0])
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Displays the names of the last 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Diplays all the pokemon in a given area",
			callback:    commandExplore,
		},
	}
}

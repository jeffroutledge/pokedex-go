package main

import (
	"bufio"
	"fmt"
	"os"
)

var locationUrls []string
var locationUrlIndex int

// var loactions []PokeLocations

func main() {
	locationUrls = append(locationUrls, "https://pokeapi.co/api/v2/location/")
	locationUrlIndex = 0

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		text := cleanInput(reader.Text())
		if command, exists := cliCommands()[text[0]]; exists {
			err := command.callback()
			if err != nil {
				panic(err)
			}
		} else {
			handleInvalidCmd(text[0])
		}
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/jeffroutledge/CliPokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type cliConfig struct {
	nextUrl     string
	previousUrl string
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range cliCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMap() error {
	locationData := pokeapi.GetLocations(locationUrls[locationUrlIndex])
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	locationUrls = append(locationUrls, locationData.Next)
	locationUrlIndex++
	return nil
}

func commandMapBack() error {
	if locationUrlIndex <= 1 {
		fmt.Println("Can't go back beyond the start, try going forward")
		return nil
	}

	locationUrlIndex -= 2 //have to decrement by 2 to compensate for the additional step forward at the end of map
	locationData := pokeapi.GetLocations(locationUrls[locationUrlIndex])
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExit() error {
	defer os.Exit(3)
	return nil
}

func handleInvalidCmd(text string) {
	defer fmt.Printf("Unknown command: %s\n", text)
}

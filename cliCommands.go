package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params []string) error
}

func commandHelp(cfg *config, params []string) error {
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

func commandMap(cfg *config, params []string) error {
	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(cfg *config, params []string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("Can't go back beyond the start, try going forward")
	}

	locationsResp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *config, params []string) error {
	fmt.Printf("Exploring %s...\n", params[0])

	exploreResp, err := cfg.pokeapiClient.GetPokemonInArea(&params[0])
	if err != nil {
		return err
	}

	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func commandExit(cfg *config, params []string) error {
	defer os.Exit(3)
	return nil
}

func handleInvalidCmd(text string) {
	defer fmt.Printf("Unknown command: %s\n", text)
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
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

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
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
	if len(params) == 0 {
		return errors.New("No area name to explore")
	}

	fmt.Printf("Exploring %s...\n", params[0])

	exploreResp, err := cfg.pokeapiClient.GetLocation(&params[0])
	if err != nil {
		return err
	}

	for _, area := range exploreResp.PokemonEncounters {
		fmt.Println(area.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("No pokemon name to catch")
	}

	catchResp, err := cfg.pokeapiClient.GetPokemon(&params[0])
	if err != nil {
		return err
	}
	pokemon := catchResp

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	if rnd.Int31n(1000) > int32(pokemon.BaseExperience) {
		fmt.Printf("%s was caught\n", pokemon.Name)
		cfg.pokedex[params[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *config, params []string) error {
	if len(params) == 0 {
		return errors.New("No pokemon name to inspect")
	}

	pokemonName := params[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return errors.New("Can't inspect a pokemon you haven't caught")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, params []string) error {
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
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

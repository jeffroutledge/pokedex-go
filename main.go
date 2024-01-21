package main

import (
	"time"

	"github.com/jeffroutledge/CliPokedex/internal/pokeapi"
)

var locationUrls []string
var locationUrlIndex int

// var loactions []PokeLocations

func main() {
	locationUrls = append(locationUrls, "https://pokeapi.co/api/v2/location/")
	locationUrlIndex = 0

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}

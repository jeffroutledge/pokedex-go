package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PokeLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (PokeLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)
	if ok {
		pokeLocations := PokeLocations{}
		err := json.Unmarshal([]byte(val), &pokeLocations)
		if err != nil {
			log.Fatal(err)
		}
		return pokeLocations, nil
	}

	res, err := http.Get(fmt.Sprintf(url))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	c.cache.Add(url, body)

	pokeLocations := PokeLocations{}
	err = json.Unmarshal([]byte(body), &pokeLocations)
	if err != nil {
		log.Fatal(err)
	}

	return pokeLocations, nil
}

package main

import (
	"errors"
	"fmt"

	"github.com/AdonaIsium/pokedex/internal/pokecache"
)

type LocationArea struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMapForward(c *config, cache *pokecache.Cache) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.nextLocationsURL, cache)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(c *config, cache *pokecache.Cache) error {
	if c.prevLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	locationResp, err := c.pokeapiClient.ListLocations(c.prevLocationsURL, cache)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationResp.Next
	c.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

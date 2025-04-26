package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func getLocationArea(url string) (*LocationArea, error) {
	client := http.DefaultClient

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatalf("response status code: %d,\nbody: %s", resp.StatusCode, body)

	}

	var locationArea *LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return nil, err
	}

	return locationArea, nil
}

func commandMap(c *Config) error {

	url := "https://pokeapi.co/api/v2/location-area"

	var locationArea *LocationArea
	var err error

	if c.Next != "" {
		locationArea, err = getLocationArea(c.Next)
		if err != nil {
			return err
		}
	} else {
		locationArea, err = getLocationArea(url)
		if err != nil {
			return err
		}
	}

	for _, result := range locationArea.Results {
		fmt.Printf("%s\n", result.Name)
	}
	c.Next = locationArea.Next
	c.Previous = locationArea.Previous
	fmt.Printf("Next: %s, Previous: %s", c.Next, c.Previous)
	fmt.Println()

	return nil
}

func commandMapBack(c *Config) error {

	var locationArea *LocationArea
	var err error

	if c.Previous != "" {
		locationArea, err = getLocationArea(c.Previous)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("You're on the first page")
		return nil
	}

	for _, result := range locationArea.Results {
		fmt.Printf("%s\n", result.Name)
	}
	c.Next = locationArea.Next
	c.Previous = locationArea.Previous
	fmt.Printf("Next: %s, Previous: %s", c.Next, c.Previous)
	fmt.Println()

	return nil
}

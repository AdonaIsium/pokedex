package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(args ...string) (PokemonResp, error) {
	if len(args) != 1 {
		return PokemonResp{}, errors.New("incorrect number of args provided to get pokemon")
	}
	pokemon := args[0]

	url := baseURL + "/pokemon/" + pokemon

	if data, exists := c.cache.Get(url); exists {
		locationsResp := PokemonResp{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return PokemonResp{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}
	c.cache.Add(url, dat)

	exploreDetails := PokemonResp{}
	err = json.Unmarshal(dat, &exploreDetails)
	if err != nil {
		return PokemonResp{}, err
	}

	return exploreDetails, nil
}

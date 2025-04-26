package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(args ...string) (RespLocationDetails, error) {
	if len(args) != 1 {
		return RespLocationDetails{}, errors.New("incorrect number of args provided to get location")
	}
	location := args[0]

	url := baseURL + "/location-area/" + location

	if data, exists := c.cache.Get(url); exists {
		locationsResp := RespLocationDetails{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespLocationDetails{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RespLocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationDetails{}, err
	}
	c.cache.Add(url, dat)

	exploreDetails := RespLocationDetails{}
	err = json.Unmarshal(dat, &exploreDetails)
	if err != nil {
		return RespLocationDetails{}, err
	}

	return exploreDetails, nil
}

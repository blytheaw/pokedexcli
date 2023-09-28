package pokeapi

import (
	"encoding/json"
	"io"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(url *string) (Locations, error) {
	if url == nil {
		*url = baseURL + "/location-area?offset=0&limit=20"
	}
	fullUrl := *url

	if data, ok := c.cache.Get(fullUrl); ok {
		resp := Locations{}
		err := json.Unmarshal(data, &resp)

		return Locations{}, err
	}

	println("Making HTTP request to PokeAPI")

	resp, err := c.client.Get(fullUrl)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(fullUrl, data)
	return locations, nil
}

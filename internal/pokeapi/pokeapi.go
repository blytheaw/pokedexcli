package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/blytheaw/pokedexcli/internal/pokecache"
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

func GetLocations(url string, cache *pokecache.Cache) (Locations, error) {
	locations := Locations{}

	data, ok := cache.Get(url)

	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return locations, errors.New("Error requesting locations from PokeAPI")
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)

		cache.Add(url, data)
	}

	err := json.Unmarshal(data, &locations)
	if err != nil {
		return locations, err
	}

	return locations, nil
}

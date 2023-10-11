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

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationByID(name string) (Location, error) {
	fullUrl := baseURL + "/location-area/" + name

	if data, ok := c.cache.Get(fullUrl); ok {
		resp := Location{}
		err := json.Unmarshal(data, &resp)
		return resp, err
	}

	resp, err := c.client.Get(fullUrl)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	location := Location{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(fullUrl, data)
	return location, nil
}

func (c *Client) ListLocations(url *string) (Locations, error) {
	fullUrl := baseURL + "/location-area?offset=0&limit=20"
	if url != nil {
		fullUrl = *url
	}

	if data, ok := c.cache.Get(fullUrl); ok {
		resp := Locations{}
		err := json.Unmarshal(data, &resp)

		return resp, err
	}

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

package pokeapi

import "encoding/json"

const LOCATION_AREA_BASE_URL = "https://pokeapi.co/api/v2/location-area/"

type locationAreaResponseType struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationAreaDetailResponseType struct {
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

func (c *Client) GetLocationAreaDetail(areaName string) (locationAreaDetailResponseType, error) {
	return cachedGet[locationAreaDetailResponseType](c, LOCATION_AREA_BASE_URL+areaName)
}

func (c *Client) getLocationAreas(url string) (locationAreaResponseType, error) {
	if val, ok := c.cache.Get(url); ok {
		var result locationAreaResponseType
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}

		c.previousUrl = result.Previous
		c.nextUrl = result.Next
		return result, nil
	}

	cachedResponse, err := cachedGet[locationAreaResponseType](c, url)
	if err != nil {
		return locationAreaResponseType{}, err
	}
	c.previousUrl = cachedResponse.Previous
	c.nextUrl = cachedResponse.Next
	return cachedResponse, nil
}

func (c *Client) GetNextLocationAreas() (locationAreaResponseType, error) {
	if c.nextUrl == "" {
		return c.getLocationAreas(LOCATION_AREA_BASE_URL)
	}
	return c.getLocationAreas(c.nextUrl)
}

func (c *Client) GetPreviousLocationAreas() (locationAreaResponseType, error) {
	if c.previousUrl == "" {
		return c.getLocationAreas(LOCATION_AREA_BASE_URL)
	}
	return c.getLocationAreas(c.previousUrl)
}

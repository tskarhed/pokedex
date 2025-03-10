package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tskarhed/pokedex/internal/pokecache"
)

type Client struct {
	nextUrl     string
	previousUrl string
	cache       *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		nextUrl:     "",
		previousUrl: "",
		cache:       pokecache.NewCache(1 * time.Minute),
	}
}

// Handles caching of the response body for a given url
func cachedGet[T any](c *Client, url string) (T, error) {
	if val, ok := c.cache.Get(url); ok {
		var result T
		err := json.Unmarshal(val, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}
	response, err := http.Get(url)
	if err != nil {
		var zero T
		return zero, err
	}

	defer response.Body.Close()

	var decodedResponse T
	err = json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		return decodedResponse, err
	}

	// Encode the response body and write to cache
	jsonData, err := json.Marshal(decodedResponse)
	if err != nil {
		return decodedResponse, err
	}

	c.cache.Add(url, jsonData)

	return decodedResponse, nil

}

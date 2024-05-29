package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emmaahmads/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(5 * time.Minute)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}

// methods for Client
// List locations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	locations := RespShallowLocations{}
	// if there is a pageURL, use that
	if pageURL != nil {
		url = *pageURL
	}
	log.Println("test 1")
	val, ok := c.cache.Get(url)
	// check in cache
	if !ok {
		log.Println("test 2")
		//make http request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return locations, err
		}

		//send http request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return locations, err
		}

		//close the io interface to avoid resources leaks
		defer resp.Body.Close()

		//read the entire contents of an io.Reader into a byte slice
		val, err = io.ReadAll(resp.Body)
		if err != nil {
			return locations, err
		}

		c.cache.Add(url, val)
	}

	//unmarshal the byte slice into a struct
	err := json.Unmarshal(val, &locations)
	if err != nil {
		return locations, err
	}
	return locations, nil
}

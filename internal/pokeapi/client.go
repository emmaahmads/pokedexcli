package pokeapi

import (
	"encoding/json"
	"io"
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

// NewClient -
func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(10 * time.Second)
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

	val, ok := c.cache.Get(url)
	// check in cache
	if !ok {
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

// Explore locations
func (c *Client) ExploreLocations(name *string) (RespShallowExploreLocations, error) {
	url := baseURL + "/location-area"
	pokemon_list := RespShallowExploreLocations{}

	if name != nil {
		url = url + "/" + *name
	}

	val, ok := c.cache.Get(url)
	// check in cache
	if !ok {
		//make http request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return pokemon_list, err
		}

		//send http request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return pokemon_list, err
		}

		//close the io interface to avoid resources leaks
		defer resp.Body.Close()

		//read the entire contents of an io.Reader into a byte slice
		val, err = io.ReadAll(resp.Body) //ok = false
		if err != nil {
			return pokemon_list, err
		}

		c.cache.Add(url, val)
	}

	//unmarshal the byte slice into a struct
	err := json.Unmarshal(val, &pokemon_list)
	if err != nil {
		return pokemon_list, err
	}
	return pokemon_list, nil
}

func (c *Client) GetPokemon(name *string) (RespPokemon, error) {
	url := baseURL + "/pokemon"
	pokemon := RespPokemon{}

	if name != nil {
		url = url + "/" + *name
	}

	//make http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon, err
	}

	//send http request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, err
	}

	//close the io interface to avoid resources leaks
	defer resp.Body.Close()

	//read the entire contents of an io.Reader into a byte slice
	val, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemon, err
	}

	err = json.Unmarshal(val, &pokemon)
	if err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

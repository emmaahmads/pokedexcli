package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Client -
type Client struct {
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
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

// methods for Client
// List locations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	// if there is a pageURL, use that
	if pageURL != nil {
		url = *pageURL
	}

	//make http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	//send http request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	//close the io interface to avoid resources leaks
	defer resp.Body.Close()

	locations := RespShallowLocations{}
	//read the entire contents of an io.Reader into a byte slice
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	//unmarshal the byte slice into a struct
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locations, nil
}

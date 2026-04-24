package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type apiLocationResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c Client) GetLocations(pageURL *string) (apiLocationResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return apiLocationResp{}, err
	}

	var locations apiLocationResp
	if err = json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return apiLocationResp{}, err
	}
	return locations, nil
}

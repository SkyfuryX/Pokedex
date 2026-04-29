package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	httpCache  *Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		httpCache: NewCache(interval),
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

	var locations apiLocationResp
	elem, exists := c.httpCache.cache[url] //check cache before calling .Get()
	if exists {
		fmt.Printf("%v retrieved\n", url)
		if err := json.Unmarshal(elem.val, &locations); err != nil {
			return apiLocationResp{}, err
		}
		return locations, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return apiLocationResp{}, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return apiLocationResp{}, err
	}
	data, err := json.Marshal(locations)
	c.httpCache.Add(url, data)
	fmt.Printf("%v added\n", url)

	return locations, nil
}

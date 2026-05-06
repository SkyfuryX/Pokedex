package pokeapi

import (
	"encoding/json"
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

type apiAreaResp struct {
	Name       string `json:"name"`
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
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
	return locations, nil
}

func (c Client) GetArea(pageURL *string) (apiAreaResp, error) {
	var area apiAreaResp
	url := *pageURL

	elem, exists := c.httpCache.cache[url] //check cache before calling .Get()
	if exists {
		if err := json.Unmarshal(elem.val, &area); err != nil {
			return apiAreaResp{}, err
		}
		return area, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return apiAreaResp{}, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&area); err != nil {
		return apiAreaResp{}, err
	}
	data, err := json.Marshal(area)
	c.httpCache.Add(url, data)

	return area, nil
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c Client) GetPokemon(pageURL *string) (Pokemon, error) {
	var pokemon Pokemon
	url := *pageURL

	elem, exists := c.httpCache.cache[url] //check cache before calling .Get()
	if exists {
		if err := json.Unmarshal(elem.val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return Pokemon{}, err
	}
	data, err := json.Marshal(pokemon)
	c.httpCache.Add(url, data)

	return pokemon, nil
}

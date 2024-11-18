package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/interyx/pokedexcli/pokecache"
	"io"
	"log"
	"net/http"
)

const baseURL string = "https://pokeapi.co/api/v2"

var cache = pokecache.NewCache("5s")

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
}

type LocationDetail struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
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
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func ReadBody(url string) ([]byte, error) {
	var body []byte
	data, cached := cache.Get(url)
	if !cached {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return []byte{}, fmt.Errorf("Request failed with status code %v", res.StatusCode)
		}
		if err != nil {
			return []byte{}, err
		}
		cache.Add(url, body)
	} else {
		body = data
	}
	return body, nil
}

func GetLocations(url string) ([]Location, string, string, error) {
	body, err := ReadBody(url)
	if err != nil {
		return []Location{}, "", "", err
	}
	var res LocationResponse
	if err = json.Unmarshal(body, &res); err != nil {
		return []Location{}, "", "", err
	}
	return res.Results, res.Previous, res.Next, nil
}

func GetNextLocation(url *string) ([]Location, string, string, error) {
	var target string
	if url == nil || *url == "" {
		target = baseURL + "/location-area"
	} else {
		target = *url
	}
	return GetLocations(target)
}

func GetPokemonAtLocation(url string) ([]string, error) {
	body := ReadBody(url)
	var res LocationDetail
	if err := json.Unmarshal(body, &res); err != nil {
		return []string{}, err
	}
	results := []string{}
	for _, encounter := range res.PokemonEncounters {
		results = append(results, encounter.Pokemon.Name)
	}
	return results, nil
}

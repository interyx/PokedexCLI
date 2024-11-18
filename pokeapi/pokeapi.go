package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const baseURL string = "https://pokeapi.co/api/v2/"

type Response struct {
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

func ReadBody(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func GetLocations(url string) ([]Location, string, string) {
	body := ReadBody(url)
	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatal(err)
	}
	return res.Results, res.Previous, res.Next
}

func GetNextLocation(url *string) ([]Location, string, string) {
	var target string
	if url == nil || *url == "" {
		target = baseURL + "/location-area"
	} else {
		target = *url
	}
	return GetLocations(target)
}

package poke_api

import (
	"encoding/json"
	"net/http"
)

type PokeApiResponse struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func CallAPI(url string) (PokeApiResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiResponse{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return PokeApiResponse{}, err
	}
	defer res.Body.Close()

	var apiResponse PokeApiResponse
	if err := json.NewDecoder(res.Body).Decode(&apiResponse); err != nil {
		return PokeApiResponse{}, err
	}
	return apiResponse, nil
}

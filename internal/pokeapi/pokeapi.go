package pokeapi

import (
	"encoding/json"
	"net/http"
)

type APILocationArea struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func callLocationArea(url string) (APILocationArea, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return APILocationArea{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return APILocationArea{}, err
	}
	defer res.Body.Close()

	var apiResponse APILocationArea
	if err := json.NewDecoder(res.Body).Decode(&apiResponse); err != nil {
		return APILocationArea{}, err
	}
	return apiResponse, nil
}

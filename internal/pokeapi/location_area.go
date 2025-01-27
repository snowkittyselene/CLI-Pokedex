package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type APILocationArea struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

const locationAreaUrl = baseApiUrl + "/location-area/"

func (client *Client) CallLocationArea(url *string) (APILocationArea, error) {
	pageUrl := locationAreaUrl
	if url != nil {
		pageUrl = *url
	}

	if val, ok := client.cache.Get(pageUrl); ok {
		res := APILocationArea{}
		if err := json.Unmarshal(val, &res); err != nil {
			return APILocationArea{}, err
		}
		return res, nil
	}

	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return APILocationArea{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.httpClient.Do(req)
	if err != nil {
		return APILocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return APILocationArea{}, nil
	}

	var apiResponse APILocationArea
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return APILocationArea{}, err
	}
	return apiResponse, nil
}

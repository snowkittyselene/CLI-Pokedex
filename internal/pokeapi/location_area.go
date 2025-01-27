package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const locationAreaUrl = baseApiUrl + "location-area/"

func (client *Client) CallLocationArea(url *string) (LocationArea, error) {
	pageUrl := locationAreaUrl
	if url != nil {
		pageUrl = *url
	}

	if val, ok := client.cache.Get(pageUrl); ok {
		res := LocationArea{}
		if err := json.Unmarshal(val, &res); err != nil {
			return LocationArea{}, err
		}
		return res, nil
	}

	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	client.cache.Add(pageUrl, data)

	var apiResponse LocationArea
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return LocationArea{}, err
	}
	return apiResponse, nil
}

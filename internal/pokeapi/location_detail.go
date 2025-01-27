package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationDetails struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounter     `json:"pokemon_encounters"`
}
type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod  `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}
type EncounterVersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}
type PokemonEncounter struct {
	Pokemon        Pokemon          `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

func (client *Client) CallLocationDetails(url string) (LocationDetails, error) {
	pageUrl := locationAreaUrl + url

	if val, ok := client.cache.Get(pageUrl); ok {
		res := LocationDetails{}
		if err := json.Unmarshal(val, &res); err != nil {
			return LocationDetails{}, err
		}
		return res, nil
	}

	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}
	client.cache.Add(pageUrl, data)

	var apiResponse LocationDetails
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return LocationDetails{}, err
	}
	return apiResponse, nil
}

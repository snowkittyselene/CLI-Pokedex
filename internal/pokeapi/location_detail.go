package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationDetails struct {
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
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (client *Client) CallLocationDetails(location string) (LocationDetails, error) {
	pageUrl := locationAreaUrl + location

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

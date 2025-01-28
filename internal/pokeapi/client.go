package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/snowkittyselene/pokecache"
)

const baseApiUrl string = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

type APIResponse interface {
	LocationArea | LocationDetails | PokemonInfo | PokemonSpeciesInfo
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func CallAPI[T APIResponse](client *Client, url string) (T, error) {
	var response, empty T
	if val, ok := client.cache.Get(url); ok {
		if err := json.Unmarshal(val, &response); err != nil {
			return empty, err
		}
		return response, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return empty, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return empty, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return empty, err
	}
	client.cache.Add(url, data)

	if err := json.Unmarshal(data, &response); err != nil {
		return empty, err
	}
	return response, nil
}

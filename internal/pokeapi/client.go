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
	LocationArea | LocationDetails | PokemonInfo
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func CallAPI[T APIResponse](client *Client, url string, typeEmpty T) (T, error) {
	response := typeEmpty
	if val, ok := client.cache.Get(url); ok {
		if err := json.Unmarshal(val, &response); err != nil {
			return typeEmpty, err
		}
		return response, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return typeEmpty, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return typeEmpty, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return typeEmpty, err
	}
	client.cache.Add(url, data)

	if err := json.Unmarshal(data, &response); err != nil {
		return typeEmpty, err
	}
	return response, nil
}

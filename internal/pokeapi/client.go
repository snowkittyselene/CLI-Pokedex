package pokeapi

import (
	"net/http"
	"time"

	"github.com/snowkittyselene/pokecache"
)

const baseApiUrl string = "https://pokeapi.co/api/v2/"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

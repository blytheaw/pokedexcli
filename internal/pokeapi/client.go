package pokeapi

import (
	"net/http"
	"time"

	"github.com/blytheaw/pokedexcli/internal/pokecache"
)

type Client struct {
	cache  pokecache.Cache
	client http.Client
}

func NewClient(httpTimeout, cacheInterval time.Duration) Client {
	client := Client{
		cache: pokecache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: httpTimeout,
		},
	}

	return client
}

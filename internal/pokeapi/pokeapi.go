package pokeapi

import (
	"net/http"
	"time"

	"github.com/Evilcmd/pokedexclifollow/internal/pokecache"
)

const BaseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		pokecache.NewCache(time.Second * 10),
		http.Client{
			Timeout: time.Minute,
		},
	}
}

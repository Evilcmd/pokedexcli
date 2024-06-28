package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(alturl *string) (LocationAreaDefn, error) {
	endpoint := "/location-area"
	fullurl := BaseUrl + endpoint

	if alturl != nil {
		fullurl = *alturl
	}

	body, ok := c.cache.Get(fullurl)
	if ok {
		fmt.Println("Cache Hit")
		LocationAreaRes := LocationAreaDefn{}
		err := json.Unmarshal(body, &LocationAreaRes)

		if err != nil {
			return LocationAreaDefn{}, err
		}

		return LocationAreaRes, nil
	}

	fmt.Println("Cache Miss")

	// creates a request but not sent yet
	req, err := http.NewRequest("GET", fullurl, nil)

	if err != nil {
		return LocationAreaDefn{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaDefn{}, nil
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreaDefn{}, fmt.Errorf("error: returned with staus code %v", res.StatusCode)
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDefn{}, err
	}

	LocationAreaRes := LocationAreaDefn{}
	err = json.Unmarshal(body, &LocationAreaRes)

	if err != nil {
		return LocationAreaDefn{}, err
	}

	c.cache.Add(fullurl, body)

	return LocationAreaRes, nil

}

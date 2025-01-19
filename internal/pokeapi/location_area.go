package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	url := c.baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	locationResp := LocationAreaResponse{}
	if data, found := c.cache.Get(url); found {
		fmt.Println(data, found)
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			fmt.Println(err)
			return LocationAreaResponse{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(url, dat)
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationResp, nil

}

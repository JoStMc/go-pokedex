package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	body, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Location{}, err
		} 
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Location{}, err
		} 
		if resp.StatusCode > 299 {
			return Location{}, fmt.Errorf("unable to complete request. Status code: %d", resp.StatusCode)
		}
		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return Location{}, err
		} 
		c.cache.Add(url, body)
	} 

	var data Location
	if err := json.Unmarshal(body, &data); err != nil {
		return Location{}, err
	} 
	return data, nil
}


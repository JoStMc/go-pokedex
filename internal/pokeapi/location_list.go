package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
	    url = *pageURL
	} 

	body, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Locations{}, err
		} 
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Locations{}, err
		} 
		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return Locations{}, err
		} 
		c.cache.Add(url, body)
	} 

	var data Locations
	if err := json.Unmarshal(body, &data); err != nil {
		return Locations{}, err
	} 
	return data, nil
}

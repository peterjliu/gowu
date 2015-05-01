// A library for making Weather Underground calls for weather forecast.
// Requires a free API key.
// Free call limits are: 500/day, 10/minute
package gowu

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const urlfmt = "http://api.wunderground.com/api/%s/conditions/q/%s/%s.json"

type WuClient struct {
	apikey string
}

func NewClient(key string) *WuClient {
	return &WuClient{apikey: key}
}

// Get current conditions for a city, state
// city should have underscores between words if multiple words
func (w *WuClient) GetConditions(city, state string) (*CurConditions, error) {
	c := &CurConditions{}
	req_url := fmt.Sprintf(urlfmt, w.apikey, state, city)
	log.Println(req_url)
	resp, err := http.Get(req_url)
	if err != nil {
		return nil, err
	}
	// Note: Weather Underground responds with 200 even if API key is wrong
	// TODO detect using wrong api key

	err = json.NewDecoder(resp.Body).Decode(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

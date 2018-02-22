// A library for making Weather Underground calls for weather forecast.
// Requires a free API key.
// Free call limits are: 500/day, 10/minute
package gowu

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	urlfmt     = "http://api.wunderground.com/api/%s/%s/q/%s/%s.json"
	forecast   = "forecast"
	conditions = "conditions"
)

type WuClient struct {
	apikey string
}

func NewClient(key string) *WuClient {
	return &WuClient{apikey: key}
}

func APICall(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Issue the request to the API
func (w *WuClient) getForEntity(e interface{}, section, state, city string) error {
	req_url := fmt.Sprintf(urlfmt, w.apikey, section, state, city)
	log.Println(req_url)

	body, err := APICall(req_url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, e)
	if err != nil {
		return err
	}

	return nil
}

// Get current conditions for a city, state
// city should have underscores between words if multiple words
func (w *WuClient) GetConditions(city, state string) (*CurConditions, error) {
	var c CurConditions
	err := w.getForEntity(&c, conditions, state, city)
	if err != nil {
		log.Error(err)
	}
	return &c, nil
}

func (w *WuClient) GetForecast(city, state string) (*Forecast, error) {
	var f ForecastResponse
	err := w.getForEntity(&f, forecast, state, city)
	if err != nil {
		log.Error(err)
	}
	return &f.Forecast, nil
}

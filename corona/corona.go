package corona

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Corona struct {
	endpoint string
}

type Country struct {
	Name        string `json:"country"`
	Cases       int    `json:"cases"`
	TodayCases  int    `json:"todayCases"`
	Deaths      int    `json:"deaths"`
	TodayDeaths int    `json:"todayDeaths"`
}

const (
	userAgent   = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"
	endpoint    = "https://disease.sh"
	jpPath      = "/v3/covid-19/countries/JP"
	rankingPath = "/v3/covid-19/countries?sort=cases"
)

func New() *Corona {
	return &Corona{endpoint}
}

func (c *Corona) JP() (*Country, error) {
	resp, err := get(c.endpoint + jpPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var country Country
	if err := json.NewDecoder(resp.Body).Decode(&country); err != nil {
		return nil, err
	}

	return &country, nil
}

func (c *Corona) WorldRanking() ([]*Country, error) {
	resp, err := get(c.endpoint + rankingPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var countries []*Country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		return nil, err
	}

	return countries, nil
}

func get(endpoint string) (*http.Response, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", endpointURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

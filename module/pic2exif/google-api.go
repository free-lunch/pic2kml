package pic2exif

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"errors"
	"encoding/json"
)

type AddressResult struct {
	Address string `json:"formatted_address"`
}

type AddressResponse struct {
	Results []AddressResult `json:"results"`
	Status  string   `json:"status"`
}

func GetAddress(lat, lon float64, apiKey string) (string, error) {
	if apiKey == "" {
		return "", errors.New("Not exist Google API key")
	}

	base_url := "https://maps.googleapis.com/maps/api/geocode/json"
	query := fmt.Sprintf("?latlng=%f,%f&key=%s", lat, lon, apiKey)
	resp, err := http.Get(base_url + query)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s, err := GetResults([]byte(body))
	if err != nil {
		return "", err
	}
	return s.Results[0].Address, nil
}

func GetResults(body []byte) (*AddressResponse, error) {
	var s = new(AddressResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

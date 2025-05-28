package handlers_and_api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-all/structs"
	"io"
	"net/http"
	"net/url"
)

// takes a location and turns it into coordinates
func geocodeLocation(location string, apiKey string) (float64, float64, error) {
	encoded := url.QueryEscape(location)
	apiURL := fmt.Sprintf("https://api.opencagedata.com/geocode/v1/json?q=%s&key=%s", encoded, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result structs.GeocodeResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, 0, err
	}

	if len(result.Results) == 0 {
		return 0, 0, fmt.Errorf("no results for location: %s", location)
	}

	return result.Results[0].Geometry.Lat, result.Results[0].Geometry.Lng, nil
}

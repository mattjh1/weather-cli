package api

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type GeoLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func GetWeatherData(city string) []byte {
	geoLocation, err := getGeoLocation(city)
	if err != nil {
		fmt.Println("Error getting coordinates:", err)
		return nil
	}

	client := resty.New()
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m&hourly=temperature_2m&daily=weather_code,temperature_2m_max,temperature_2m_min,sunrise,sunset,daylight_duration,sunshine_duration,rain_sum,showers_sum,snowfall_sum&timezone=auto&forecast_days=3", geoLocation.Lat, geoLocation.Lon)

	response, err := client.R().Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return response.Body()
}

func getGeoLocation(city string) (GeoLocation, error) {
	client := resty.New()
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", city)

	response, err := client.R().Get(url)
	if err != nil {
		return GeoLocation{}, err
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(response.Body(), &apiResponse); err != nil {
		fmt.Println("Error unmarshaling response:", err)
		fmt.Println("Response:", string(response.Body()))
		return GeoLocation{}, err
	}

	results, ok := apiResponse["results"].([]interface{})
	if !ok || len(results) == 0 {
		return GeoLocation{}, fmt.Errorf("no coordinates found for %s", city)
	}

	firstResult, ok := results[0].(map[string]interface{})
	if !ok {
		return GeoLocation{}, fmt.Errorf("unexpected result format for %s", city)
	}

	latitude, latOk := firstResult["latitude"].(float64)
	longitude, lonOk := firstResult["longitude"].(float64)

	if !latOk || !lonOk {
		return GeoLocation{}, fmt.Errorf("latitude or longitude not found for %s", city)
	}

	return GeoLocation{Lat: latitude, Lon: longitude}, nil
}

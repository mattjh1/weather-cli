package display

import (
	"encoding/json"
	"fmt"
)

func DisplayWeatherInfo(responseBody []byte) {
	var data map[string]interface{}
	if err := json.Unmarshal(responseBody, &data); err != nil {
		fmt.Println("Error parsing response:", err)
		fmt.Println("Response:", string(responseBody))
		return
	}

	// Display current temperature
	currentTemperature, ok := data["current"].(map[string]interface{})["temperature_2m"].(float64)
	if !ok {
		fmt.Println("Error: Current temperature information not found.")
		return
	}
	fmt.Printf("Current Temperature in %s: %.1f°C\n", data["timezone_abbreviation"], currentTemperature)

	// Display three-day forecast
	dailyForecast, ok := data["daily"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Daily forecast information not found.")
		return
	}

	days := dailyForecast["time"].([]interface{})
	maxTemps := dailyForecast["temperature_2m_max"].([]interface{})
	minTemps := dailyForecast["temperature_2m_min"].([]interface{})

	fmt.Println("Three-Day Forecast:")
	for i, day := range days {
		date := day.(string)
		maxTemp := maxTemps[i].(float64)
		minTemp := minTemps[i].(float64)

		sunrise := dailyForecast["sunrise"].([]interface{})[i].(string)
		sunset := dailyForecast["sunset"].([]interface{})[i].(string)
		daylightDuration := dailyForecast["daylight_duration"].([]interface{})[i].(float64)
		sunshineDuration := dailyForecast["sunshine_duration"].([]interface{})[i].(float64)

		fmt.Printf("Date: %s\n", date)
		fmt.Printf("  Max Temperature: %.1f°C\n", maxTemp)
		fmt.Printf("  Min Temperature: %.1f°C\n", minTemp)
		fmt.Printf("  Sunrise: %s\n", sunrise)
		fmt.Printf("  Sunset: %s\n", sunset)
		fmt.Printf("  Daylight Duration: %.2f hours\n", daylightDuration/3600)
		fmt.Printf("  Sunshine Duration: %.2f hours\n", sunshineDuration/3600)
	}
}

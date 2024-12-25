package display

import (
	"encoding/json"
	"fmt"
)

// Helper function to format time to display as "HH:MM"
func formatTime(time string) string {
	return time[:5] // Extract "HH:MM" format from "HH:MM:SS"
}

// DisplayWeatherInfo - Displays weather information with better formatting
func DisplayWeatherInfo(responseBody []byte) {
	var data map[string]interface{}
	if err := json.Unmarshal(responseBody, &data); err != nil {
		fmt.Println("Error parsing response:", err)
		fmt.Println("Response:", string(responseBody))
		return
	}

	// Get timezone abbreviation
	timezone, ok := data["timezone_abbreviation"].(string)
	if !ok {
		fmt.Println("Error: Timezone information not found.")
		return
	}

	// Display current temperature
	currentTemperature, ok := data["current"].(map[string]interface{})["temperature_2m"].(float64)
	if !ok {
		fmt.Println("Error: Current temperature information not found.")
		return
	}
	fmt.Printf("\n🌡  Current Temperature in %s: %.1f°C\n", timezone, currentTemperature)

	// Display weather conditions
	weatherDescription, ok := data["current"].(map[string]interface{})["weathercode"].(float64)
	if !ok {
		weatherDescription = 0
	}

	// Display weather icon and description based on code
	var weatherIcon, weatherText string
	switch weatherDescription {
	case 0:
		weatherIcon, weatherText = "☀️", "Clear sky"
	case 1:
		weatherIcon, weatherText = "🌤", "Mainly clear"
	case 2:
		weatherIcon, weatherText = "🌥", "Partly cloudy"
	case 3:
		weatherIcon, weatherText = "☁️", "Cloudy"
	case 45, 48:
		weatherIcon, weatherText = "🌫️", "Fog"
	case 51, 53, 55:
		weatherIcon, weatherText = "🌧️", "Light rain"
	case 56, 57:
		weatherIcon, weatherText = "❄️", "Freezing rain"
	case 61, 63, 65:
		weatherIcon, weatherText = "🌦️", "Moderate rain"
	default:
		weatherIcon, weatherText = "🌈", "Weather info unavailable"
	}

	fmt.Printf("  Weather: %s %s\n", weatherIcon, weatherText)

	// Display three-day forecast
	dailyForecast, ok := data["daily"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Daily forecast information not found.")
		return
	}

	days := dailyForecast["time"].([]interface{})
	maxTemps := dailyForecast["temperature_2m_max"].([]interface{})
	minTemps := dailyForecast["temperature_2m_min"].([]interface{})
	sunriseTimes := dailyForecast["sunrise"].([]interface{})
	sunsetTimes := dailyForecast["sunset"].([]interface{})
	daylightDuration := dailyForecast["daylight_duration"].([]interface{})
	sunshineDuration := dailyForecast["sunshine_duration"].([]interface{})

	fmt.Printf("\n📅 Three-Day Forecast:\n")
	for i, day := range days {
		date := day.(string)
		maxTemp := maxTemps[i].(float64)
		minTemp := minTemps[i].(float64)

		// Display formatted times and durations
		sunrise := formatTime(sunriseTimes[i].(string))
		sunset := formatTime(sunsetTimes[i].(string))
		daylight := daylightDuration[i].(float64) / 3600 // Convert to hours
		sunshine := sunshineDuration[i].(float64) / 3600 // Convert to hours

		// Print the forecast for the day
		fmt.Printf("\n📅 Date: %s\n", date)
		fmt.Printf("  🌡 Max Temp: %.1f°C  Min Temp: %.1f°C\n", maxTemp, minTemp)
		fmt.Printf("  🌅 Sunrise: %s  🌇 Sunset: %s\n", sunrise, sunset)
		fmt.Printf("  🌞 Daylight Duration: %.2f hrs  ☀️ Sunshine Duration: %.2f hrs\n", daylight, sunshine)
	}

	fmt.Println("\n=============================")
}

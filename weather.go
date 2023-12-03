package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type GeoLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get the weather forecast for a city",
	Run:   runCmd,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringP("city", "c", "", "City for which to get the weather forecast")
	rootCmd.Flags().BoolP("help", "h", false, "Help for the weather command")

	viper.BindPFlag("city", rootCmd.Flags().Lookup("city"))
	viper.BindPFlag("help", rootCmd.Flags().Lookup("help"))
}

func initConfig() {
	viper.SetConfigFile(".weather")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func runCmd(cmd *cobra.Command, args []string) {
	if viper.GetBool("help") {
		cmd.Help()
		return
	}

	city := viper.GetString("city")

	if city == "" {
		fmt.Println("Please provide a city using the -c flag.")
		return
	}

	// Step 1: Get coordinates for the city using Open Meteo Geocoding API
	geoLocation, err := getGeoLocation(city)
	if err != nil {
		fmt.Println("Error getting coordinates:", err)
		return
	}

	// Step 2: Use coordinates to get weather information
	client := resty.New()
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m&hourly=temperature_2m&daily=weather_code,temperature_2m_max,temperature_2m_min,sunrise,sunset,daylight_duration,sunshine_duration,rain_sum,showers_sum,snowfall_sum&timezone=auto&forecast_days=3", geoLocation.Lat, geoLocation.Lon)


	response, err := client.R().Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	displayWeatherInfo(response.Body())
}

func displayWeatherInfo(responseBody []byte) {
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

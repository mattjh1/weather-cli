package api

import (
	"fmt"
	"os"
)

// gptWeatherReport generates a poetic weather report using AI
func gptWeatherReport(weatherData map[string]interface{}) {
	// Implement your GPT logic here
	// Use weatherData as context and generate a creative weather report
}

// validateGPTApiKey checks if the OpenAI API key is available.
func validateGPTApiKey() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY not found. --gpt flag requires an API key.")
		os.Exit(1)
	}
}

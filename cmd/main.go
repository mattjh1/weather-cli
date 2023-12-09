package main

import (
	"fmt"
	"os"

	"github.com/mattjh1/weather-cli/api"
	config "github.com/mattjh1/weather-cli/config"
	"github.com/mattjh1/weather-cli/display"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get the weather forecast for a city",
	Run:   runCmd,
}

func init() {
	config.InitConfig(rootCmd)
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

	if viper.GetBool("gpt") {
		// generatePoeticWeatherReport(city)
		return
	}

	weatherData := api.GetWeatherData(city)
	display.DisplayWeatherInfo(weatherData)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

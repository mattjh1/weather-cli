package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get the weather forecast for a city",
	Run:   runCmd,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringP("city", "c", "", "City for which to get the weather forecast")
	rootCmd.Flags().BoolP("help", "h", false, "Help for the weather command")
	rootCmd.Flags().BoolP("gpt", "g", false, "Generate a poetic weather report using GPT AI")

	viper.BindPFlag("city", rootCmd.Flags().Lookup("city"))
	viper.BindPFlag("help", rootCmd.Flags().Lookup("help"))
	viper.BindPFlag("gpt", rootCmd.Flags().Lookup("gpt"))
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

	if viper.GetBool("gpt") {
		// generatePoeticWeatherReport(city)
		return
	}

	weatherData := GetWeatherData(city)
	DisplayWeatherInfo(weatherData)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

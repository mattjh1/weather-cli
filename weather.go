package main

import (
	"fmt"
	"os"

	"github.com/mattjh1/weather-cli/api"
	config "github.com/mattjh1/weather-cli/config"
	"github.com/mattjh1/weather-cli/display"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/mbndr/figlet4go"
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
	fig := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorRed,
		figlet4go.ColorGreen,
		figlet4go.ColorYellow,
		figlet4go.ColorBlue,
		figlet4go.ColorMagenta,
		figlet4go.ColorCyan,
	}

	renderStr, err := fig.RenderOpts("Weather CLI", options)
	if err != nil {
		fmt.Println("Error generating Figlet banner:", err)
		return
	}
	fmt.Print(renderStr)

	if viper.GetBool("help") {
		cmd.Help()
		return
	}

	city := viper.GetString("city")

	if city == "" {
		fmt.Println("Please provide a city using the -c flag.")
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

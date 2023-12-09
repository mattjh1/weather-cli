// config/config.go
package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitConfig(cmd *cobra.Command) {
	cobra.OnInitialize(initViper)

	cmd.Flags().StringP("city", "c", "", "City for which to get the weather forecast")
	cmd.Flags().BoolP("help", "h", false, "Help for the weather command")
	cmd.Flags().BoolP("gpt", "g", false, "Generate a poetic weather report using GPT AI")

	viper.BindPFlag("city", cmd.Flags().Lookup("city"))
	viper.BindPFlag("help", cmd.Flags().Lookup("help"))
	viper.BindPFlag("gpt", cmd.Flags().Lookup("gpt"))
}

func initViper() {
	viper.SetConfigFile(".weather")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
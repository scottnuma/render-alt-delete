package render

import (
	"fmt"
	"os"

	"github.com/scottnuma/render-alt-delete/internal/log"
	"github.com/spf13/viper"
)

func GetConfig() (token string, endpoint string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/render-alt-delete")
	viper.SetDefault("render_api_endpoint", "api.render.com")
	viper.SetEnvPrefix("RAD")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignoring
		} else {
			// Config file was found but another error was produced
			log.Logger.Error("failed to read config file", "err", err)
			panic(err)
		}
	}

	profile := viper.GetString("profile")
	if profile == "" {
		token = viper.GetString("render_api_token")
		if token == "" {
			log.Logger.Error("RAD_RENDER_API_TOKEN is not set")
			os.Exit(1)
		}

		endpoint = viper.GetString("render_api_endpoint")
		return

	}

	profileConfig := viper.Sub(fmt.Sprintf("profiles.%s", profile))
	if profileConfig == nil {
		log.Logger.Error("profile not found", "profile", profile)
		os.Exit(1)
	}
	profileConfig.SetDefault("render_api_endpoint", "api.render.com")

	token = profileConfig.GetString("render_api_token")
	if token == "" {
		log.Logger.Error("render_api_token is not set for profile", "profile", profile)
		os.Exit(1)
	}

	endpoint = profileConfig.GetString("render_api_endpoint")
	return
}

package render

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() (token string, endpoint string, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/render-alt-delete")
	viper.SetDefault("render_api_endpoint", "api.render.com")
	viper.SetEnvPrefix("RAD")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignoring
		} else {
			// Config file was found but another error was produced
			return "", "", fmt.Errorf("failed to read config file: %w", err)
		}
	}

	profile := viper.GetString("profile")
	if profile == "" {
		token = viper.GetString("render_api_token")
		if token == "" {
			return "", "", fmt.Errorf("RAD_RENDER_API_TOKEN  is not set")
		}

		endpoint = viper.GetString("render_api_endpoint")
		return

	}

	profileConfig := viper.Sub(fmt.Sprintf("profiles.%s", profile))
	if profileConfig == nil {
		return "", "", fmt.Errorf("%s profile not found", profile)
	}
	profileConfig.SetDefault("render_api_endpoint", "api.render.com")

	token = profileConfig.GetString("render_api_token")
	if token == "" {
		return "", "", fmt.Errorf("render_api_token is not set for profile %s", profile)
	}

	endpoint = profileConfig.GetString("render_api_endpoint")
	return
}

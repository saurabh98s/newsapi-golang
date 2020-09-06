package config

import (
	"newsbackend/logger"

	"github.com/spf13/viper"
)

// api holds the config for the API
type api struct {
	apiKey string
}

// load returns the config for the API
func (apiConfig *api) load() {
	logger.Log.Info("Reading API config...")
	// viper.SetEnvPrefix("api")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	apiConfig.apiKey = viper.GetString("api_key")
}

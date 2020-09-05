package config

type config struct {
	apiConfig api
}

var configuration = &config{}

// Load loads the config into the configuration object
func Load() {
	configuration.apiConfig.load()
}


// APIKey returns the api port
func APIKey() string {
	return configuration.apiConfig.apiKey
}
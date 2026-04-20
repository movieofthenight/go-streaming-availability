package streaming

import (
	"strings"
)

// Deprecated: Use NewAPIClientFromApiKey instead
func NewAPIClientFromRapidAPIKey(rapidApiKey string, config *Configuration) *APIClient {
	return NewAPIClientFromApiKey(rapidApiKey, config)
}

func newAPIClientFromRapidAPIKey(rapidApiKey string, config *Configuration) *APIClient {
	if config == nil {
		config = NewConfiguration()
	}
	config.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	config.Servers = ServerConfigurations{
		{
			URL: "https://streaming-availability.p.rapidapi.com",
		},
	}
	return NewAPIClient(config)
}

func newAPIClientFromDevelopersPlatformAPIKey(developersPlatformApiKey string, config *Configuration) *APIClient {
	if config == nil {
		config = NewConfiguration()
	}
	config.AddDefaultHeader("X-API-Key", developersPlatformApiKey)
	config.Servers = ServerConfigurations{
		{
			URL: "https://api.movieofthenight.com/v4",
		},
	}
	return NewAPIClient(config)
}

func NewAPIClientFromApiKey(apiKey string, config *Configuration) *APIClient {
	if strings.HasPrefix(apiKey, "motn-key-") {
		return newAPIClientFromDevelopersPlatformAPIKey(apiKey, config)
	}
	return newAPIClientFromRapidAPIKey(apiKey, config)
}

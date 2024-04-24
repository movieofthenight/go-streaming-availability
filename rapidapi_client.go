package streaming

// NewAPIClientFromRapidAPIKey creates an *APIClient with the given RapidAPI key.
// A nil config will be replaced with the default configuration.
func NewAPIClientFromRapidAPIKey(rapidApiKey string, config *Configuration) *APIClient {
	if config == nil {
		config = NewConfiguration()
	}
	config.AddDefaultHeader("X-RapidAPI-Key", rapidApiKey)
	return NewAPIClient(config)
}

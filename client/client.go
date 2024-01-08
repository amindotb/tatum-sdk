package client

import (
	"tatumsdk/constants"

	"github.com/go-resty/resty/v2"
)



type ApiClient struct {
	baseUrl string
	apiKey  string
	Chain   string
	Type    string
	Client  *resty.Client
}

// ApiClientOption is a shorthand for the initializing function
type ApiClientOption func(*ApiClient)

func NewApiClient(options ...ApiClientOption) *ApiClient {
	apiClient := ApiClient{
		baseUrl: constants.DEFAULT_API_BASE_URL,
	}

	for _, option := range options {
		option(&apiClient)
	}

	if apiClient.apiKey == "" {
		panic(constants.ERROR_CLIENT_NO_API_KEY)
	}

	apiClient.Chain = ""
	apiClient.Type = constants.DEFAULT_TYPE
	apiClient.Client = resty.New().
		SetBaseURL(apiClient.baseUrl).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"type": apiClient.Type,
		}).
		SetHeader("x-api-key", apiClient.apiKey)

	return &apiClient
}

func WithAPIKey(apiKey string) ApiClientOption {
	return func(apiClient *ApiClient) {
		apiClient.apiKey = apiKey
	}
}

func WithAPIKeyAndChain(apiKey string, chain string) ApiClientOption {
	return func(apiClient *ApiClient) {
		apiClient.apiKey = apiKey
		apiClient.Chain = chain
	}
}

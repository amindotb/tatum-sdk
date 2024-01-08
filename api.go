package main

import (
	"tatumsdk/client"
	"tatumsdk/domains/address"
	"tatumsdk/domains/subscription"
)

type Api struct {
	apiClient    *client.ApiClient
	subscription *subscription.Subscription
	address      *address.Address
}

func NewAPI(options ...client.ApiClientOption) *Api {
	return &Api{
		apiClient:    client.NewApiClient(options...),
		subscription: subscription.NewSubscription(options...),
		address:      address.NewAddress(options...),
	}
}

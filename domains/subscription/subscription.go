package subscription

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tatumsdk/client"
)

// Types of subscription
const (
	// Incoming Transaction Notifications
	TYPE_INCOMING_NATIVE_TX     = "INCOMING_NATIVE_TX"     // See: https://docs.tatum.io/docs/notifications/notification-types/incoming-native-transactions
	TYPE_INCOMING_INTERNAL_TX   = "INCOMING_INTERNAL_TX"   // See: https://docs.tatum.io/docs/notifications/notification-types/incoming-internal-transactions
	TYPE_INCOMING_FUNGIBLE_TX   = "INCOMING_FUNGIBLE_TX"   // See: https://docs.tatum.io/docs/notifications/notification-types/incoming-tokens
	TYPE_INCOMING_NFT_TX        = "INCOMING_NFT_TX"        // See: https://docs.tatum.io/docs/notifications/notification-types/incoming-nfts
	TYPE_INCOMING_MULTITOKEN_TX = "INCOMING_MULTITOKEN_TX" // See: https://docs.tatum.io/docs/notifications/notification-types/incoming-multitokens

	// Outgoing Transaction Notifications
	TYPE_OUTGOING_NATIVE_TX     = "OUTGOING_NATIVE_TX"     // See: https://docs.tatum.io/docs/notifications/notification-types/outgoing-native-transactions
	TYPE_OUTGOING_INTERNAL_TX   = "OUTGOING_INTERNAL_TX"   // See: https://docs.tatum.io/docs/notifications/notification-types/outgoing-internal-transactions
	TYPE_OUTGOING_FUNGIBLE_TX   = "OUTGOING_FUNGIBLE_TX"   // See: https://docs.tatum.io/docs/notifications/notification-types/outgoing-tokens
	TYPE_OUTGOING_NFT_TX        = "OUTGOING_NFT_TX"        // See: https://docs.tatum.io/docs/notifications/notification-types/outgoing-nfts
	TYPE_OUTGOING_MULTITOKEN_TX = "OUTGOING_MULTITOKEN_TX" // See: https://docs.tatum.io/docs/notifications/notification-types/outgoing-multitokens
	TYPE_OUTGOING_FAILED_TX     = "OUTGOING_FAILED_TX"     // See: https://docs.tatum.io/docs/notifications/notification-types/outgoing-failed-transactions

	// Special Abstraction Notifications
	TYPE_ADDRESS_EVENT              = "ADDRESS_EVENT"              // See: https://docs.tatum.io/docs/notifications/notification-types/address-event
	TYPE_PAID_FEE                   = "PAID_FEE"                   // See: https://docs.tatum.io/docs/notifications/notification-types/paid-fee
	TYPE_FAILED_TX_PER_BLOCK        = "FAILED_TX_PER_BLOCK"        // See: https://docs.tatum.io/docs/notifications/notification-types/failed-transactions-in-a-block
	TYPE_CONTRACT_ADDRESS_LOG_EVENT = "CONTRACT_ADDRESS_LOG_EVENT" // See: https://docs.tatum.io/docs/notifications/notification-types/contract-address-log-event
)

// Orders of fetching notifications
const (
	DIRECTION_ASC  = "asc"
	DIRECTION_DESC = "desc"
)

type SubscriptionAttribute struct {
	Address string `json:"address"`
	Chain   string `json:"chain"`
	Url     string `json:"url"`
}

type SubscriptionInput struct {
	Type      string                `json:"type"`
	Attribute SubscriptionAttribute `json:"attr"`
}

type GetAllResponse struct {
	Id   string                `json:"id"`
	Type string                `json:"type"`
	Attr SubscriptionAttribute `json:"attr"`
}

type SubscribeResponse struct {
	Id string `json:"id"`
}

type UnsubscribeResponse struct {
	Id string `json:"id"`
}

type GetAllExecutedWebhooksResponse struct {
	Id             string                `json:"id"`
	SubscriptionId string                `json:"subscriptionId"`
	Network        string                `json:"network"`
	Url            string                `json:"url"`
	Type           string                `json:"type"`
	Data           NotificationsData     `json:"data"`
	Response       NotificationsResponse `json:"response"`
	Timestamp      uint64                `json:"timestamp"`
	Failed         bool                  `json:"failed"`
}

type NotificationsData struct {
	Address          string `json:"address"`
	Amount           string `json:"amount"`
	Asset            string `json:"asset"`
	BlockNumber      uint64 `json:"blockNumber"`
	CounterAddress   string `json:"counterAddress"`
	TxId             string `json:"txId"`
	Type             string `json:"type"`
	Chain            string `json:"chain"`
	SubscriptionType string `json:"subscriptionType"`
}

type NotificationsResponse struct {
	Code         uint   `json:"code"`
	Data         string `json:"data"`
	NetworkError bool   `json:"networkError"`
}

type Subscription struct {
	apiClient *client.ApiClient
}

func NewSubscription(options ...client.ApiClientOption) *Subscription {
	return &Subscription{
		apiClient: client.NewApiClient(options...),
	}
}

func (subscription *Subscription) GetAll(pagination client.Pagination, address string) ([]GetAllResponse, error) {
	var successResponse []GetAllResponse

	pagination = client.FormatPagination(pagination)
	resp, err := subscription.apiClient.Client.R().
		SetQueryParams(map[string]string{
			"pageSize": fmt.Sprintf("%d", pagination.PageSize),
			"offset":   fmt.Sprintf("%d", pagination.Offset),
			"address":  address,
		}).
		Get("/subscription")
	err = client.ParseResponse(resp, &successResponse)

	if err != nil {
		return successResponse, err
	}

	return successResponse, nil
}

func (subscription *Subscription) Subscribe(subscriptionInput SubscriptionInput) (SubscribeResponse, error) {
	var successResponse SubscribeResponse
	jsonBody, err := json.Marshal(subscriptionInput)
	if err != nil {
		return successResponse, err
	}

	resp, err := subscription.apiClient.Client.R().
		SetBody(jsonBody).
		Post("/subscription")
	err = client.ParseResponse(resp, &successResponse)

	if err != nil {
		return successResponse, err
	}

	return successResponse, nil
}

func (subscription *Subscription) Unsubscribe(subscriptionId string) (UnsubscribeResponse, error) {
	var successResponse UnsubscribeResponse

	resp, err := subscription.apiClient.Client.R().
		SetPathParams(map[string]string{"subscriptionId": subscriptionId}).
		Delete("/subscription/{subscriptionId}")
	err = client.ParseResponse(resp, &successResponse)

	if err != nil {
		return successResponse, err
	}

	return successResponse, nil
}

func (subscription *Subscription) GetAllExecutedWebhooks(pagination client.Pagination, direction string, filterFailed bool) ([]GetAllExecutedWebhooksResponse, error) {
	var successResponse []GetAllExecutedWebhooksResponse

	pagination = client.FormatPagination(pagination)

	resp, err := subscription.apiClient.Client.R().
		SetQueryParams(map[string]string{
			"pageSize":     fmt.Sprintf("%d", pagination.PageSize),
			"offset":       fmt.Sprintf("%d", pagination.Offset),
			"direction":    direction,
			"filterFailed": strconv.FormatBool(filterFailed),
		}).
		Get("/subscription/webhook")
	err = client.ParseResponse(resp, &successResponse)

	if err != nil {
		return successResponse, err
	}

	return successResponse, nil
}

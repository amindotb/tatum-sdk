package address

import (
	"fmt"
	"strings"
	"tatumsdk/client"
)

type GetBalanceResultResponse struct {
	Result []GetBalanceResponse `json:"result"`
}

type GetTransactionResultResponse struct {
	Result []GetTransactionResponse `json:"result"`
}

type GetBalanceResponse struct {
	Asset            string `json:"asset"`
	Address          string `json:"address"`
	Decimals         uint   `json:"decimals"`
	Balance          string `json:"balance"`
	Type             string `json:"type"`
	TokenAddress     string `json:"tokenAddress"`
	TokenId          string `json:"tokenId"`
	LastUpdatedBlock uint64 `json:"lastUpdatedBlock"`
}

type GetTransactionResponse struct {
	Chain              string `json:"chain"`
	Hash               string `json:"hash"`
	Address            string `json:"address"`
	BlockNumber        uint64 `json:"blockNumber"`
	TransactionIndex   uint64 `json:"transactionIndex"`
	TransactionType    string `json:"transactionType"`
	TransactionSubtype string `json:"transactionSubtype"`
	Amount             string `json:"amount"`
	CounterAddress     string `json:"counterAddress"`
	Timestamp          uint64 `json:"timestamp"`
}

type Address struct {
	apiClient *client.ApiClient
}

func NewAddress(options ...client.ApiClientOption) *Address {
	return &Address{
		apiClient: client.NewApiClient(options...),
	}
}

func (address *Address) GetBalance(pagination client.Pagination, addresses []string, chain string) ([]GetBalanceResponse, error) {
	var successResponse GetBalanceResultResponse

	pagination = client.FormatPagination(pagination)
	resp, err := address.apiClient.Client.R().
		SetQueryParams(map[string]string{
			"pageSize":  fmt.Sprintf("%d", pagination.PageSize),
			"offset":    fmt.Sprintf("%d", pagination.Offset),
			"addresses": strings.Join(addresses, ","),
			"chain":     chain,
		}).
		Get("/data/balances")
	err = client.ParseResponse(resp, &successResponse)

	if err != nil {
		return successResponse.Result, err
	}

	return successResponse.Result, nil
}

func (address *Address) GetTransactions(pagination client.Pagination, addresses []string, chain string) ([]GetTransactionResponse, error) {
	var successResponse GetTransactionResultResponse

	pagination = client.FormatPagination(pagination)
	resp, err := address.apiClient.Client.R().
		SetQueryParams(map[string]string{
			"pageSize":  fmt.Sprintf("%d", pagination.PageSize),
			"offset":    fmt.Sprintf("%d", pagination.Offset),
			"addresses": strings.Join(addresses, ","),
			"chain":     chain,
		}).
		Get("/data/transactions")
	err = client.ParseResponse(resp, &successResponse)

	if err != nil {
		return successResponse.Result, err
	}

	return successResponse.Result, nil
}

package client

import (
	"encoding/json"
	"errors"
	"tatumsdk/constants"

	"github.com/go-resty/resty/v2"
)

type ErrorResponse struct {
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       string `json:"data"`
}

type Pagination struct {
	PageSize uint `json:"pageSize"`
	Offset   uint `json:"offset"`
}

func ParseResponse(resp *resty.Response, successResponse interface{}) error {
	var errorResponse ErrorResponse
	if resp.StatusCode() != 200 {
		json.Unmarshal(resp.Body(), &errorResponse)
		return errors.New(errorResponse.ErrorCode)
	}
	json.Unmarshal(resp.Body(), &successResponse)
	return nil
}

func FormatPagination(pagination Pagination) Pagination {
	if pagination.PageSize == 0 {
		pagination.PageSize = constants.DEFAULT_PAGE_SIZE
	}

	if pagination.Offset == 0 {
		pagination.Offset = constants.DEFAULT_OFFSET
	}

	return pagination
}

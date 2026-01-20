// Package errors provides error types for Backpack Exchange SDK.
package errors

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Common errors
var (
	ErrNoCredentials = errors.New("backpack: API credentials not provided")
	ErrInvalidKey    = errors.New("backpack: invalid API key format")
)

// APIError represents an error returned by the Backpack Exchange API.
type APIError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	RawBody    string `json:"-"`
}

func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("backpack API error [%d]: %s - %s", e.StatusCode, e.Code, e.Message)
	}
	if e.Message != "" {
		return fmt.Sprintf("backpack API error [%d]: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("backpack API error [%d]: %s", e.StatusCode, e.RawBody)
}

// HasCode checks if the error matches a specific error code.
func (e *APIError) HasCode(code ErrorCode) bool {
	return e.Code == string(code)
}

// RequestError represents an error that occurred while making a request.
type RequestError struct {
	Method  string
	URL     string
	Message string
	Err     error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("backpack request error [%s %s]: %s", e.Method, e.URL, e.Message)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}

// ParseAPIError parses an API error response.
func ParseAPIError(statusCode int, body []byte) *APIError {
	apiErr := &APIError{
		StatusCode: statusCode,
		RawBody:    string(body),
	}

	// Try to parse as JSON
	var errResp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.Unmarshal(body, &errResp); err == nil {
		apiErr.Code = errResp.Code
		if errResp.Message != "" {
			apiErr.Message = errResp.Message
		} else if errResp.Error != "" {
			apiErr.Message = errResp.Error
		}
	}

	return apiErr
}

// IsAPIError checks if an error is an APIError and returns it.
func IsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

// IsRequestError checks if an error is a RequestError and returns it.
func IsRequestError(err error) (*RequestError, bool) {
	var reqErr *RequestError
	if errors.As(err, &reqErr) {
		return reqErr, true
	}
	return nil, false
}

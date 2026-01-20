// Package http provides HTTP client functionality for Backpack Exchange API.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/errors"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/internal/auth"
)

const (
	DefaultBaseURL = "https://api.backpack.exchange"
	DefaultTimeout = 30 * time.Second
	DefaultWindow  = int64(5000)
)

// Client is the HTTP client for making API requests.
type Client struct {
	baseURL    string
	httpClient *http.Client
	signer     *auth.Signer
	window     int64
	debug      bool
}

// Config holds configuration for the HTTP client.
type Config struct {
	BaseURL    string
	HTTPClient *http.Client
	Signer     *auth.Signer
	Window     int64
	Debug      bool
}

// NewClient creates a new HTTP client with the given configuration.
func NewClient(cfg Config) *Client {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: DefaultTimeout,
		}
	}

	window := cfg.Window
	if window == 0 {
		window = DefaultWindow
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
		signer:     cfg.Signer,
		window:     window,
		debug:      cfg.Debug,
	}
}

// SetSigner sets the signer for authenticated requests.
func (c *Client) SetSigner(signer *auth.Signer) {
	c.signer = signer
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, path string, params map[string]string, result any) error {
	return c.doRequest(ctx, http.MethodGet, path, params, nil, "", result)
}

// GetAuthenticated performs an authenticated GET request.
func (c *Client) GetAuthenticated(ctx context.Context, path string, params map[string]string, instruction string, result any) error {
	return c.doAuthenticatedRequest(ctx, http.MethodGet, path, params, nil, instruction, result)
}

// Post performs a POST request.
func (c *Client) Post(ctx context.Context, path string, body any, result any) error {
	return c.doRequest(ctx, http.MethodPost, path, nil, body, "", result)
}

// PostAuthenticated performs an authenticated POST request.
func (c *Client) PostAuthenticated(ctx context.Context, path string, body any, instruction string, result any) error {
	return c.doAuthenticatedRequest(ctx, http.MethodPost, path, nil, body, instruction, result)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, path string, body any, result any) error {
	return c.doRequest(ctx, http.MethodDelete, path, nil, body, "", result)
}

// DeleteAuthenticated performs an authenticated DELETE request.
func (c *Client) DeleteAuthenticated(ctx context.Context, path string, body any, instruction string, result any) error {
	return c.doAuthenticatedRequest(ctx, http.MethodDelete, path, nil, body, instruction, result)
}

// Patch performs a PATCH request.
func (c *Client) Patch(ctx context.Context, path string, body any, result any) error {
	return c.doRequest(ctx, http.MethodPatch, path, nil, body, "", result)
}

// PatchAuthenticated performs an authenticated PATCH request.
func (c *Client) PatchAuthenticated(ctx context.Context, path string, body any, instruction string, result any) error {
	return c.doAuthenticatedRequest(ctx, http.MethodPatch, path, nil, body, instruction, result)
}

// PostBatchOrders performs an authenticated batch order request.
func (c *Client) PostBatchOrders(ctx context.Context, path string, orders []map[string]any, result any) error {
	if c.signer == nil {
		return errors.ErrNoCredentials
	}

	reqURL := c.baseURL + "/" + path

	bodyBytes, err := json.Marshal(orders)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	headers := c.signer.GenerateBatchHeaders(orders, c.window)
	for k, v := range headers.ToMap() {
		req.Header.Set(k, v)
	}

	return c.executeRequest(req, result)
}

func (c *Client) doRequest(ctx context.Context, method, path string, params map[string]string, body any, _ string, result any) error {
	reqURL := c.buildURL(path, params)

	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return c.executeRequest(req, result)
}

func (c *Client) doAuthenticatedRequest(ctx context.Context, method, path string, params map[string]string, body any, instruction string, result any) error {
	if c.signer == nil {
		return errors.ErrNoCredentials
	}

	reqURL := c.buildURL(path, params)

	var bodyReader io.Reader
	var signParams map[string]any

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)

		// Convert body to map for signing
		if err := json.Unmarshal(bodyBytes, &signParams); err != nil {
			return fmt.Errorf("failed to unmarshal body for signing: %w", err)
		}
	} else if len(params) > 0 {
		signParams = make(map[string]any, len(params))
		for k, v := range params {
			signParams[k] = v
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	headers := c.signer.GenerateHeaders(instruction, signParams, c.window)
	for k, v := range headers.ToMap() {
		req.Header.Set(k, v)
	}

	return c.executeRequest(req, result)
}

func (c *Client) buildURL(path string, params map[string]string) string {
	reqURL := c.baseURL + "/" + path

	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Set(k, v)
		}
		reqURL += "?" + values.Encode()
	}

	return reqURL
}

func (c *Client) executeRequest(req *http.Request, result any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &errors.RequestError{
			Method:  req.Method,
			URL:     req.URL.String(),
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.ParseAPIError(resp.StatusCode, bodyBytes)
	}

	if result != nil && len(bodyBytes) > 0 {
		if err := json.Unmarshal(bodyBytes, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

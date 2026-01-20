package backpack

import (
	"net/http"
	"time"
)

// Option is a functional option for configuring the Client.
type Option func(*options)

type options struct {
	publicKey  string
	secretKey  string
	baseURL    string
	timeout    time.Duration
	window     int64
	debug      bool
	httpClient *http.Client
}

func defaultOptions() *options {
	return &options{
		baseURL: "https://api.backpack.exchange",
		timeout: 30 * time.Second,
		window:  5000,
	}
}

// WithCredentials sets the API credentials.
func WithCredentials(publicKey, secretKey string) Option {
	return func(o *options) {
		o.publicKey = publicKey
		o.secretKey = secretKey
	}
}

// WithBaseURL sets a custom base URL for the API.
func WithBaseURL(url string) Option {
	return func(o *options) {
		o.baseURL = url
	}
}

// WithTimeout sets the HTTP timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

// WithWindow sets the signature validity window in milliseconds.
func WithWindow(ms int64) Option {
	return func(o *options) {
		o.window = ms
	}
}

// WithDebug enables debug logging.
func WithDebug(debug bool) Option {
	return func(o *options) {
		o.debug = debug
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(o *options) {
		o.httpClient = client
	}
}

// Package backpack provides a Go SDK for Backpack Exchange API.
package backpack

import (
	"net/http"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/internal/auth"
	internalhttp "github.com/solomeowl/backpack-exchange-sdk-go/backpack/internal/http"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/services"
)

// Client is the main entry point for the Backpack Exchange SDK.
type Client struct {
	httpClient *internalhttp.Client

	// Public APIs
	System            *services.SystemService
	Assets            *services.AssetsService
	Markets           *services.MarketsService
	Trades            *services.TradesService
	BorrowLendMarkets *services.BorrowLendMarketsService
	Prediction        *services.PredictionService

	// Authenticated APIs
	Account    *services.AccountService
	Capital    *services.CapitalService
	Orders     *services.OrdersService
	Positions  *services.PositionsService
	BorrowLend *services.BorrowLendService
	RFQ        *services.RFQService
	Strategy   *services.StrategyService
	History    *services.HistoryService
}

// NewClient creates a new Backpack Exchange client.
func NewClient(opts ...Option) (*Client, error) {
	cfg := defaultOptions()
	for _, opt := range opts {
		opt(cfg)
	}

	// Create HTTP client
	httpClient := cfg.httpClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: cfg.timeout,
		}
	}

	// Create signer if credentials are provided
	var signer *auth.Signer
	if cfg.publicKey != "" && cfg.secretKey != "" {
		var err error
		signer, err = auth.NewSigner(cfg.publicKey, cfg.secretKey)
		if err != nil {
			return nil, err
		}
	}

	// Create internal HTTP client
	internalClient := internalhttp.NewClient(internalhttp.Config{
		BaseURL:    cfg.baseURL,
		HTTPClient: httpClient,
		Signer:     signer,
		Window:     cfg.window,
		Debug:      cfg.debug,
	})

	c := &Client{
		httpClient: internalClient,
	}

	// Initialize public services
	c.System = services.NewSystemService(internalClient)
	c.Assets = services.NewAssetsService(internalClient)
	c.Markets = services.NewMarketsService(internalClient)
	c.Trades = services.NewTradesService(internalClient)
	c.BorrowLendMarkets = services.NewBorrowLendMarketsService(internalClient)
	c.Prediction = services.NewPredictionService(internalClient)

	// Initialize authenticated services
	c.Account = services.NewAccountService(internalClient)
	c.Capital = services.NewCapitalService(internalClient)
	c.Orders = services.NewOrdersService(internalClient)
	c.Positions = services.NewPositionsService(internalClient)
	c.BorrowLend = services.NewBorrowLendService(internalClient)
	c.RFQ = services.NewRFQService(internalClient)
	c.Strategy = services.NewStrategyService(internalClient)
	c.History = services.NewHistoryService(internalClient)

	return c, nil
}

// SetCredentials sets or updates the API credentials.
func (c *Client) SetCredentials(publicKey, secretKey string) error {
	signer, err := auth.NewSigner(publicKey, secretKey)
	if err != nil {
		return err
	}
	c.httpClient.SetSigner(signer)
	return nil
}

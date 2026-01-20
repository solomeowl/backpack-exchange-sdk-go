package services

import (
	"context"
	"strconv"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// MarketsService provides market-related operations.
type MarketsService struct {
	client HTTPClient
}

// NewMarketsService creates a new MarketsService.
func NewMarketsService(client HTTPClient) *MarketsService {
	return &MarketsService{client: client}
}

// GetMarketsParams represents parameters for getting markets.
type GetMarketsParams struct {
	MarketType []enums.MarketType
}

// GetMarkets retrieves all available markets.
func (s *MarketsService) GetMarkets(ctx context.Context, params *GetMarketsParams) ([]types.Market, error) {
	var result []types.Market
	queryParams := make(map[string]string)
	if params != nil && len(params.MarketType) > 0 {
		// Note: API accepts multiple marketType values
		for i, mt := range params.MarketType {
			if i == 0 {
				queryParams["marketType"] = string(mt)
			}
		}
	}
	if err := s.client.Get(ctx, "api/v1/markets", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarket retrieves a specific market by symbol.
func (s *MarketsService) GetMarket(ctx context.Context, symbol string) (*types.Market, error) {
	var result types.Market
	params := map[string]string{"symbol": symbol}
	if err := s.client.Get(ctx, "api/v1/market", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTickerParams represents parameters for getting a ticker.
type GetTickerParams struct {
	Symbol   string
	Interval enums.TickerInterval
}

// GetTicker retrieves ticker data for a specific market.
func (s *MarketsService) GetTicker(ctx context.Context, params GetTickerParams) (*types.Ticker, error) {
	var result types.Ticker
	queryParams := map[string]string{"symbol": params.Symbol}
	if params.Interval != "" {
		queryParams["interval"] = string(params.Interval)
	}
	if err := s.client.Get(ctx, "api/v1/ticker", queryParams, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTickersParams represents parameters for getting multiple tickers.
type GetTickersParams struct {
	Interval enums.TickerInterval
}

// GetTickers retrieves ticker data for all markets.
func (s *MarketsService) GetTickers(ctx context.Context, params *GetTickersParams) ([]types.Ticker, error) {
	var result []types.Ticker
	queryParams := make(map[string]string)
	if params != nil && params.Interval != "" {
		queryParams["interval"] = string(params.Interval)
	}
	if err := s.client.Get(ctx, "api/v1/tickers", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDepthParams represents parameters for getting order book depth.
type GetDepthParams struct {
	Symbol string
	Limit  enums.DepthLimit
}

// GetDepth retrieves order book depth for a market.
func (s *MarketsService) GetDepth(ctx context.Context, params GetDepthParams) (*types.OrderBook, error) {
	var result types.OrderBook
	queryParams := map[string]string{"symbol": params.Symbol}
	if params.Limit != "" {
		queryParams["limit"] = string(params.Limit)
	}
	if err := s.client.Get(ctx, "api/v1/depth", queryParams, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetKlinesParams represents parameters for getting klines.
type GetKlinesParams struct {
	Symbol    string               // Required
	Interval  enums.KlineInterval  // Required
	StartTime int64                // Required: UTC timestamp in seconds
	EndTime   int64                // Optional: UTC timestamp in seconds
	PriceType enums.KlinePriceType // Optional: defaults to Last
}

// GetKlines retrieves kline/candlestick data for a market.
func (s *MarketsService) GetKlines(ctx context.Context, params GetKlinesParams) ([]types.Kline, error) {
	var result []types.Kline
	queryParams := map[string]string{
		"symbol":    params.Symbol,
		"interval":  string(params.Interval),
		"startTime": strconv.FormatInt(params.StartTime, 10),
	}
	if params.EndTime > 0 {
		queryParams["endTime"] = strconv.FormatInt(params.EndTime, 10)
	}
	if params.PriceType != "" {
		queryParams["priceType"] = string(params.PriceType)
	}
	if err := s.client.Get(ctx, "api/v1/klines", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarkPricesParams represents parameters for getting mark prices.
type GetMarkPricesParams struct {
	Symbol     string           // Optional
	MarketType enums.MarketType // Optional: defaults to PERP
}

// GetMarkPrices retrieves mark prices for perpetual markets.
func (s *MarketsService) GetMarkPrices(ctx context.Context, params *GetMarkPricesParams) ([]types.MarkPrice, error) {
	var result []types.MarkPrice
	queryParams := make(map[string]string)
	if params != nil {
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.MarketType != "" {
			queryParams["marketType"] = string(params.MarketType)
		}
	}
	if err := s.client.Get(ctx, "api/v1/markPrices", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpenInterestParams represents parameters for getting open interest.
type GetOpenInterestParams struct {
	Symbol string // Optional
}

// GetOpenInterest retrieves open interest for perpetual markets.
func (s *MarketsService) GetOpenInterest(ctx context.Context, params *GetOpenInterestParams) ([]types.OpenInterest, error) {
	var result []types.OpenInterest
	queryParams := make(map[string]string)
	if params != nil && params.Symbol != "" {
		queryParams["symbol"] = params.Symbol
	}
	if err := s.client.Get(ctx, "api/v1/openInterest", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFundingRatesParams represents parameters for getting funding rates.
type GetFundingRatesParams struct {
	Symbol string // Required
	Limit  int    // Optional: default 100, max 10000
	Offset int    // Optional: default 0
}

// GetFundingRates retrieves funding rate history.
func (s *MarketsService) GetFundingRates(ctx context.Context, params GetFundingRatesParams) ([]types.FundingRate, error) {
	var result []types.FundingRate
	queryParams := map[string]string{"symbol": params.Symbol}
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(params.Limit)
	}
	if params.Offset > 0 {
		queryParams["offset"] = strconv.Itoa(params.Offset)
	}
	if err := s.client.Get(ctx, "api/v1/fundingRates", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

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

// GetMarkets retrieves all available markets.
func (s *MarketsService) GetMarkets(ctx context.Context) ([]types.Market, error) {
	var result []types.Market
	if err := s.client.Get(ctx, "api/v1/markets", nil, &result); err != nil {
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

// GetTicker retrieves ticker data for a specific market.
func (s *MarketsService) GetTicker(ctx context.Context, symbol string) (*types.Ticker, error) {
	var result types.Ticker
	params := map[string]string{"symbol": symbol}
	if err := s.client.Get(ctx, "api/v1/ticker", params, &result); err != nil {
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
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(int(params.Limit))
	}
	if err := s.client.Get(ctx, "api/v1/depth", queryParams, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetKlinesParams represents parameters for getting klines.
type GetKlinesParams struct {
	Symbol    string
	Interval  enums.KlineInterval
	StartTime int64
	EndTime   int64
	Limit     int
}

// GetKlines retrieves kline/candlestick data for a market.
func (s *MarketsService) GetKlines(ctx context.Context, params GetKlinesParams) ([]types.Kline, error) {
	var result []types.Kline
	queryParams := map[string]string{
		"symbol":   params.Symbol,
		"interval": string(params.Interval),
	}
	if params.StartTime > 0 {
		queryParams["startTime"] = strconv.FormatInt(params.StartTime, 10)
	}
	if params.EndTime > 0 {
		queryParams["endTime"] = strconv.FormatInt(params.EndTime, 10)
	}
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(params.Limit)
	}
	if err := s.client.Get(ctx, "api/v1/klines", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarkPrice retrieves the mark price for a perpetual market.
func (s *MarketsService) GetMarkPrice(ctx context.Context, symbol string) (*types.MarkPrice, error) {
	var result types.MarkPrice
	params := map[string]string{"symbol": symbol}
	if err := s.client.Get(ctx, "api/v1/markPrice", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMarkPrices retrieves mark prices for all perpetual markets.
func (s *MarketsService) GetMarkPrices(ctx context.Context) ([]types.MarkPrice, error) {
	var result []types.MarkPrice
	if err := s.client.Get(ctx, "api/v1/markPrices", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpenInterest retrieves open interest for a perpetual market.
func (s *MarketsService) GetOpenInterest(ctx context.Context, symbol string) (*types.OpenInterest, error) {
	var result types.OpenInterest
	params := map[string]string{"symbol": symbol}
	if err := s.client.Get(ctx, "api/v1/openInterest", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFundingRates retrieves funding rate information.
func (s *MarketsService) GetFundingRates(ctx context.Context, symbol string) ([]types.FundingRate, error) {
	var result []types.FundingRate
	params := map[string]string{"symbol": symbol}
	if err := s.client.Get(ctx, "api/v1/fundingRates", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

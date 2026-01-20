package services

import (
	"context"
	"strconv"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// TradesService provides public trade operations.
type TradesService struct {
	client HTTPClient
}

// NewTradesService creates a new TradesService.
func NewTradesService(client HTTPClient) *TradesService {
	return &TradesService{client: client}
}

// GetRecentTradesParams represents parameters for getting recent trades.
type GetRecentTradesParams struct {
	Symbol string
	Limit  int
}

// GetRecentTrades retrieves recent public trades for a market.
func (s *TradesService) GetRecentTrades(ctx context.Context, params GetRecentTradesParams) ([]types.Trade, error) {
	var result []types.Trade
	queryParams := map[string]string{"symbol": params.Symbol}
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(params.Limit)
	}
	if err := s.client.Get(ctx, "api/v1/trades", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetHistoricalTradesParams represents parameters for getting historical trades.
type GetHistoricalTradesParams struct {
	Symbol string
	Limit  int
	Offset int
}

// GetHistoricalTrades retrieves historical public trades for a market.
func (s *TradesService) GetHistoricalTrades(ctx context.Context, params GetHistoricalTradesParams) ([]types.Trade, error) {
	var result []types.Trade
	queryParams := map[string]string{"symbol": params.Symbol}
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(params.Limit)
	}
	if params.Offset > 0 {
		queryParams["offset"] = strconv.Itoa(params.Offset)
	}
	if err := s.client.Get(ctx, "api/v1/trades/history", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

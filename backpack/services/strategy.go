package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// StrategyService provides strategy operations.
type StrategyService struct {
	client HTTPClient
}

// NewStrategyService creates a new StrategyService.
func NewStrategyService(client HTTPClient) *StrategyService {
	return &StrategyService{client: client}
}

// CreateStrategy creates a new trading strategy.
func (s *StrategyService) CreateStrategy(ctx context.Context, params types.CreateStrategyParams) (*types.Strategy, error) {
	var result types.Strategy
	body := map[string]any{
		"symbol":        params.Symbol,
		"side":          string(params.Side),
		"strategyType":  string(params.StrategyType),
		"totalQuantity": params.TotalQuantity,
	}
	if params.Duration > 0 {
		body["duration"] = params.Duration
	}
	if params.PriceLimit != "" {
		body["priceLimit"] = params.PriceLimit
	}
	if err := s.client.PostAuthenticated(ctx, "api/v1/strategy", body, "strategyCreate", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetStrategy retrieves a specific strategy.
func (s *StrategyService) GetStrategy(ctx context.Context, strategyID string) (*types.Strategy, error) {
	var result types.Strategy
	params := map[string]string{"strategyId": strategyID}
	if err := s.client.GetAuthenticated(ctx, "api/v1/strategy", params, "strategyQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelStrategy cancels a strategy.
func (s *StrategyService) CancelStrategy(ctx context.Context, strategyID string) (*types.Strategy, error) {
	var result types.Strategy
	body := map[string]any{"strategyId": strategyID}
	if err := s.client.DeleteAuthenticated(ctx, "api/v1/strategy", body, "strategyCancel", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOpenStrategies retrieves all open strategies.
func (s *StrategyService) GetOpenStrategies(ctx context.Context, symbol string) ([]types.Strategy, error) {
	var result []types.Strategy
	params := make(map[string]string)
	if symbol != "" {
		params["symbol"] = symbol
	}
	if err := s.client.GetAuthenticated(ctx, "api/v1/strategies", params, "strategyQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CancelAllStrategies cancels all open strategies.
func (s *StrategyService) CancelAllStrategies(ctx context.Context, symbol string) ([]types.Strategy, error) {
	var result []types.Strategy
	body := make(map[string]any)
	if symbol != "" {
		body["symbol"] = symbol
	}
	if err := s.client.DeleteAuthenticated(ctx, "api/v1/strategies", body, "strategyCancelAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

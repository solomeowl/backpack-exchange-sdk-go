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
		"symbol":       params.Symbol,
		"side":         string(params.Side),
		"strategyType": string(params.StrategyType),
	}
	if params.Quantity != "" {
		body["quantity"] = params.Quantity
	}
	if params.Price != "" {
		body["price"] = params.Price
	}
	if params.Duration > 0 {
		body["duration"] = params.Duration
	}
	if params.Interval > 0 {
		body["interval"] = params.Interval
	}
	if params.RandomizedIntervalQuantity != nil {
		body["randomizedIntervalQuantity"] = *params.RandomizedIntervalQuantity
	}
	if params.TimeInForce != "" {
		body["timeInForce"] = string(params.TimeInForce)
	}
	if params.PostOnly != nil {
		body["postOnly"] = *params.PostOnly
	}
	if params.ReduceOnly != nil {
		body["reduceOnly"] = *params.ReduceOnly
	}
	if params.SelfTradePrevention != "" {
		body["selfTradePrevention"] = string(params.SelfTradePrevention)
	}
	if params.ClientStrategyID != nil {
		body["clientStrategyId"] = *params.ClientStrategyID
	}
	if params.BrokerID != nil {
		body["brokerId"] = *params.BrokerID
	}
	if params.AutoBorrow != nil {
		body["autoBorrow"] = *params.AutoBorrow
	}
	if params.AutoBorrowRepay != nil {
		body["autoBorrowRepay"] = *params.AutoBorrowRepay
	}
	if params.AutoLend != nil {
		body["autoLend"] = *params.AutoLend
	}
	if params.AutoLendRedeem != nil {
		body["autoLendRedeem"] = *params.AutoLendRedeem
	}
	if params.SlippageTolerance != "" {
		body["slippageTolerance"] = params.SlippageTolerance
	}
	if params.SlippageToleranceType != "" {
		body["slippageToleranceType"] = string(params.SlippageToleranceType)
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

package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// BorrowLendService provides borrow/lend operations.
type BorrowLendService struct {
	client HTTPClient
}

// NewBorrowLendService creates a new BorrowLendService.
func NewBorrowLendService(client HTTPClient) *BorrowLendService {
	return &BorrowLendService{client: client}
}

// GetPositions retrieves all borrow/lend positions.
func (s *BorrowLendService) GetPositions(ctx context.Context) ([]types.BorrowLendPosition, error) {
	var result []types.BorrowLendPosition
	if err := s.client.GetAuthenticated(ctx, "api/v1/borrowLend/positions", nil, "borrowLendPositionQuery", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Execute executes a borrow or lend operation.
func (s *BorrowLendService) Execute(ctx context.Context, params types.BorrowLendExecuteParams) (*types.BorrowLendPosition, error) {
	var result types.BorrowLendPosition
	body := map[string]any{
		"symbol":   params.Symbol,
		"side":     string(params.Side),
		"quantity": params.Quantity,
	}
	if err := s.client.PostAuthenticated(ctx, "api/v1/borrowLend", body, "borrowLendExecute", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLiquidationPrice retrieves the estimated liquidation price.
func (s *BorrowLendService) GetLiquidationPrice(ctx context.Context, symbol string) (*types.EstimatedLiquidationPrice, error) {
	var result types.EstimatedLiquidationPrice
	params := map[string]string{"symbol": symbol}
	if err := s.client.GetAuthenticated(ctx, "api/v1/borrowLend/position/liquidationPrice", params, "borrowLendLiquidationPriceQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BorrowLendMarketsService provides borrow/lend market operations (public).
type BorrowLendMarketsService struct {
	client HTTPClient
}

// NewBorrowLendMarketsService creates a new BorrowLendMarketsService.
func NewBorrowLendMarketsService(client HTTPClient) *BorrowLendMarketsService {
	return &BorrowLendMarketsService{client: client}
}

// GetMarkets retrieves all borrow/lend markets.
func (s *BorrowLendMarketsService) GetMarkets(ctx context.Context) ([]types.BorrowLendMarket, error) {
	var result []types.BorrowLendMarket
	if err := s.client.Get(ctx, "api/v1/borrowLend/markets", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarketHistoryParams represents parameters for getting market history.
type GetMarketHistoryParams struct {
	Interval enums.BorrowLendMarketHistoryInterval // Required
	Symbol   string                                // Optional
}

// GetMarketHistory retrieves historical borrow/lend market data.
func (s *BorrowLendMarketsService) GetMarketHistory(ctx context.Context, params GetMarketHistoryParams) ([]types.BorrowLendMarketHistory, error) {
	var result []types.BorrowLendMarketHistory
	queryParams := map[string]string{
		"interval": string(params.Interval),
	}
	if params.Symbol != "" {
		queryParams["symbol"] = params.Symbol
	}
	if err := s.client.Get(ctx, "api/v1/borrowLend/markets/history", queryParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

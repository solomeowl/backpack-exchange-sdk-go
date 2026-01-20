package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// PositionsService provides position operations.
type PositionsService struct {
	client HTTPClient
}

// NewPositionsService creates a new PositionsService.
func NewPositionsService(client HTTPClient) *PositionsService {
	return &PositionsService{client: client}
}

// GetPositionsParams represents parameters for getting positions.
type GetPositionsParams struct {
	Symbol     string           // Optional: filter for a single position by symbol
	MarketType enums.MarketType // Optional: SPOT or PERP
}

// GetPositions retrieves futures positions.
func (s *PositionsService) GetPositions(ctx context.Context, params *GetPositionsParams) ([]types.Position, error) {
	var result []types.Position
	queryParams := make(map[string]string)
	if params != nil {
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.MarketType != "" {
			queryParams["marketType"] = string(params.MarketType)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "api/v1/position", queryParams, "positionQuery", &result); err != nil {
		return nil, err
	}
	return result, nil
}

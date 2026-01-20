package services

import (
	"context"

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

// GetPositions retrieves all futures positions.
func (s *PositionsService) GetPositions(ctx context.Context) ([]types.Position, error) {
	var result []types.Position
	if err := s.client.GetAuthenticated(ctx, "api/v1/position", nil, "positionQuery", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPosition retrieves a specific position by symbol.
func (s *PositionsService) GetPosition(ctx context.Context, symbol string) (*types.Position, error) {
	var result types.Position
	params := map[string]string{"symbol": symbol}
	if err := s.client.GetAuthenticated(ctx, "api/v1/position", params, "positionQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

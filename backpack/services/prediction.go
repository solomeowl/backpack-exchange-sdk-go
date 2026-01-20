package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// PredictionService provides prediction market operations (public).
type PredictionService struct {
	client HTTPClient
}

// NewPredictionService creates a new PredictionService.
func NewPredictionService(client HTTPClient) *PredictionService {
	return &PredictionService{client: client}
}

// GetMarkets retrieves all prediction markets.
func (s *PredictionService) GetMarkets(ctx context.Context) ([]types.PredictionMarket, error) {
	var result []types.PredictionMarket
	if err := s.client.Get(ctx, "api/v1/prediction", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTags retrieves all prediction market tags.
func (s *PredictionService) GetTags(ctx context.Context) ([]types.PredictionTag, error) {
	var result []types.PredictionTag
	if err := s.client.Get(ctx, "api/v1/prediction/tags", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

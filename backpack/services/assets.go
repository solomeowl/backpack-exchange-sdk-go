package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// AssetsService provides asset-related operations.
type AssetsService struct {
	client HTTPClient
}

// NewAssetsService creates a new AssetsService.
func NewAssetsService(client HTTPClient) *AssetsService {
	return &AssetsService{client: client}
}

// GetAssets retrieves all available assets.
func (s *AssetsService) GetAssets(ctx context.Context) ([]types.Asset, error) {
	var result []types.Asset
	if err := s.client.Get(ctx, "api/v1/assets", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCollateral retrieves collateral information for all assets.
func (s *AssetsService) GetCollateral(ctx context.Context) ([]types.CollateralInfo, error) {
	var result []types.CollateralInfo
	if err := s.client.Get(ctx, "api/v1/collateral", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

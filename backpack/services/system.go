package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// SystemService provides system-related operations.
type SystemService struct {
	client HTTPClient
}

// NewSystemService creates a new SystemService.
func NewSystemService(client HTTPClient) *SystemService {
	return &SystemService{client: client}
}

// GetStatus retrieves the system status.
func (s *SystemService) GetStatus(ctx context.Context) (*types.SystemStatus, error) {
	var result types.SystemStatus
	if err := s.client.Get(ctx, "api/v1/status", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Ping pings the server to check connectivity.
func (s *SystemService) Ping(ctx context.Context) error {
	return s.client.Get(ctx, "api/v1/ping", nil, nil)
}

// GetTime retrieves the server time.
func (s *SystemService) GetTime(ctx context.Context) (*types.ServerTime, error) {
	var result types.ServerTime
	if err := s.client.Get(ctx, "api/v1/time", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWallets retrieves available wallet addresses for deposits.
func (s *SystemService) GetWallets(ctx context.Context) ([]types.Wallet, error) {
	var result []types.Wallet
	if err := s.client.Get(ctx, "api/v1/wallets", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

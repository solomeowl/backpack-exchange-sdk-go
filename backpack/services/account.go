package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// AccountService provides account-related operations.
type AccountService struct {
	client HTTPClient
}

// NewAccountService creates a new AccountService.
func NewAccountService(client HTTPClient) *AccountService {
	return &AccountService{client: client}
}

// GetAccount retrieves account information.
func (s *AccountService) GetAccount(ctx context.Context) (*types.Account, error) {
	var result types.Account
	if err := s.client.GetAuthenticated(ctx, "api/v1/account", nil, "accountQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateAccount updates account settings using PATCH method.
func (s *AccountService) UpdateAccount(ctx context.Context, settings types.AccountSettings) error {
	return s.client.PatchAuthenticated(ctx, "api/v1/account", settings, "accountUpdate", nil)
}

// GetMaxBorrowQuantity retrieves the maximum borrow quantity for an asset.
func (s *AccountService) GetMaxBorrowQuantity(ctx context.Context, symbol string) (*types.MaxQuantityResponse, error) {
	var result types.MaxQuantityResponse
	params := map[string]string{"symbol": symbol}
	if err := s.client.GetAuthenticated(ctx, "api/v1/account/limits/borrow", params, "maxBorrowQuantity", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMaxOrderQuantityParams represents parameters for max order quantity.
type GetMaxOrderQuantityParams struct {
	Symbol         string
	Side           enums.Side
	Price          string
	ReduceOnly     *bool
	AutoBorrow     *bool
	AutoBorrowRepay *bool
	AutoLendRedeem *bool
}

// GetMaxOrderQuantity retrieves the maximum order quantity.
func (s *AccountService) GetMaxOrderQuantity(ctx context.Context, params GetMaxOrderQuantityParams) (*types.MaxQuantityResponse, error) {
	var result types.MaxQuantityResponse
	queryParams := map[string]string{
		"symbol": params.Symbol,
		"side":   string(params.Side),
	}
	if params.Price != "" {
		queryParams["price"] = params.Price
	}
	if params.ReduceOnly != nil {
		queryParams["reduceOnly"] = boolToString(*params.ReduceOnly)
	}
	if params.AutoBorrow != nil {
		queryParams["autoBorrow"] = boolToString(*params.AutoBorrow)
	}
	if params.AutoBorrowRepay != nil {
		queryParams["autoBorrowRepay"] = boolToString(*params.AutoBorrowRepay)
	}
	if params.AutoLendRedeem != nil {
		queryParams["autoLendRedeem"] = boolToString(*params.AutoLendRedeem)
	}
	if err := s.client.GetAuthenticated(ctx, "api/v1/account/limits/order", queryParams, "maxOrderQuantity", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMaxWithdrawalQuantityParams represents parameters for max withdrawal quantity.
type GetMaxWithdrawalQuantityParams struct {
	Symbol         string
	AutoBorrow     *bool
	AutoLendRedeem *bool
}

// GetMaxWithdrawalQuantity retrieves the maximum withdrawal quantity for an asset.
func (s *AccountService) GetMaxWithdrawalQuantity(ctx context.Context, params GetMaxWithdrawalQuantityParams) (*types.MaxQuantityResponse, error) {
	var result types.MaxQuantityResponse
	queryParams := map[string]string{"symbol": params.Symbol}
	if params.AutoBorrow != nil {
		queryParams["autoBorrow"] = boolToString(*params.AutoBorrow)
	}
	if params.AutoLendRedeem != nil {
		queryParams["autoLendRedeem"] = boolToString(*params.AutoLendRedeem)
	}
	if err := s.client.GetAuthenticated(ctx, "api/v1/account/limits/withdrawal", queryParams, "maxWithdrawalQuantity", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConvertDust converts a dust balance to USDC.
func (s *AccountService) ConvertDust(ctx context.Context, symbol string) (*types.DustConversion, error) {
	var result types.DustConversion
	body := map[string]any{"symbol": symbol}
	if err := s.client.PostAuthenticated(ctx, "api/v1/account/convertDust", body, "convertDust", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

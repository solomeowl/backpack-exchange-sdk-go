package services

import (
	"context"
	"strconv"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// CapitalService provides capital operations (deposits, withdrawals).
type CapitalService struct {
	client HTTPClient
}

// NewCapitalService creates a new CapitalService.
func NewCapitalService(client HTTPClient) *CapitalService {
	return &CapitalService{client: client}
}

// GetBalances retrieves account balances.
// Returns a map of symbol to balance (e.g., {"SOL": {Available: "1.5", ...}, "USDC": {...}})
func (s *CapitalService) GetBalances(ctx context.Context) (types.Balances, error) {
	var result types.Balances
	if err := s.client.GetAuthenticated(ctx, "api/v1/capital", nil, "balanceQuery", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetCollateralParams represents parameters for getting collateral.
type GetCollateralParams struct {
	SubaccountID *uint16
}

// GetCollateral retrieves account collateral information.
func (s *CapitalService) GetCollateral(ctx context.Context, params *GetCollateralParams) (*types.MarginAccountSummary, error) {
	var result types.MarginAccountSummary
	queryParams := make(map[string]string)
	if params != nil && params.SubaccountID != nil {
		queryParams["subaccountId"] = strconv.FormatUint(uint64(*params.SubaccountID), 10)
	}
	if err := s.client.GetAuthenticated(ctx, "api/v1/capital/collateral", queryParams, "collateralQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDepositsParams represents parameters for getting deposits.
type GetDepositsParams struct {
	From            int64
	To              int64
	Limit           int
	Offset          int
	ExcludePlatform *bool
}

// GetDeposits retrieves deposit history.
func (s *CapitalService) GetDeposits(ctx context.Context, params *GetDepositsParams) ([]types.Deposit, error) {
	var result []types.Deposit
	queryParams := make(map[string]string)
	if params != nil {
		if params.From > 0 {
			queryParams["from"] = strconv.FormatInt(params.From, 10)
		}
		if params.To > 0 {
			queryParams["to"] = strconv.FormatInt(params.To, 10)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.ExcludePlatform != nil {
			queryParams["excludePlatform"] = strconv.FormatBool(*params.ExcludePlatform)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/capital/deposits", queryParams, "depositQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDepositAddress retrieves a deposit address for a specific blockchain.
func (s *CapitalService) GetDepositAddress(ctx context.Context, blockchain enums.Blockchain) (*types.DepositAddress, error) {
	var result types.DepositAddress
	params := map[string]string{"blockchain": string(blockchain)}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/capital/deposit/address", params, "depositAddressQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWithdrawalsParams represents parameters for getting withdrawals.
type GetWithdrawalsParams struct {
	From   int64
	To     int64
	Limit  int
	Offset int
}

// GetWithdrawals retrieves withdrawal history.
func (s *CapitalService) GetWithdrawals(ctx context.Context, params *GetWithdrawalsParams) ([]types.Withdrawal, error) {
	var result []types.Withdrawal
	queryParams := make(map[string]string)
	if params != nil {
		if params.From > 0 {
			queryParams["from"] = strconv.FormatInt(params.From, 10)
		}
		if params.To > 0 {
			queryParams["to"] = strconv.FormatInt(params.To, 10)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/capital/withdrawals", queryParams, "withdrawalQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// RequestWithdrawal requests a withdrawal.
func (s *CapitalService) RequestWithdrawal(ctx context.Context, req types.WithdrawalRequest) (*types.Withdrawal, error) {
	var result types.Withdrawal
	if err := s.client.PostAuthenticated(ctx, "wapi/v1/capital/withdrawals", req, "withdraw", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWithdrawalDelay retrieves withdrawal delay settings.
func (s *CapitalService) GetWithdrawalDelay(ctx context.Context) (*types.WithdrawalDelay, error) {
	var result types.WithdrawalDelay
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/capital/withdrawals/delay", nil, "withdrawalDelayQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateWithdrawalDelayParams represents parameters for creating withdrawal delay.
type CreateWithdrawalDelayParams struct {
	WithdrawalDelayHours int    `json:"withdrawalDelayHours"`
	TwoFactorToken       string `json:"twoFactorToken"`
}

// CreateWithdrawalDelay creates or enables withdrawal delay.
func (s *CapitalService) CreateWithdrawalDelay(ctx context.Context, params CreateWithdrawalDelayParams) (*types.WithdrawalDelay, error) {
	var result types.WithdrawalDelay
	if err := s.client.PostAuthenticated(ctx, "wapi/v1/capital/withdrawals/delay", params, "withdrawalDelayCreate", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateWithdrawalDelayParams represents parameters for updating withdrawal delay.
type UpdateWithdrawalDelayParams struct {
	WithdrawalDelayHours int    `json:"withdrawalDelayHours"`
	TwoFactorToken       string `json:"twoFactorToken"`
}

// UpdateWithdrawalDelay updates withdrawal delay settings using PATCH.
func (s *CapitalService) UpdateWithdrawalDelay(ctx context.Context, params UpdateWithdrawalDelayParams) (*types.WithdrawalDelay, error) {
	var result types.WithdrawalDelay
	if err := s.client.PatchAuthenticated(ctx, "wapi/v1/capital/withdrawals/delay", params, "withdrawalDelayUpdate", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

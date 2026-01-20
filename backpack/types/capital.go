package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Balance represents an account balance for an asset.
type Balance struct {
	Available string `json:"available"`
	Locked    string `json:"locked"`
	Staked    string `json:"staked,omitempty"`
}

// Balances represents all account balances as a map of symbol to balance.
type Balances map[string]Balance

// MarginAccountSummary represents the collateral API response from GET /api/v1/capital/collateral.
type MarginAccountSummary struct {
	AssetsValue        string           `json:"assetsValue"`
	BorrowLiability    string           `json:"borrowLiability"`
	Collateral         []CollateralItem `json:"collateral"`
	IMF                string           `json:"imf"`
	UnsettledEquity    string           `json:"unsettledEquity"`
	LiabilitiesValue   string           `json:"liabilitiesValue"`
	MarginFraction     string           `json:"marginFraction,omitempty"`
	MMF                string           `json:"mmf"`
	NetEquity          string           `json:"netEquity"`
	NetEquityAvailable string           `json:"netEquityAvailable"`
	NetEquityLocked    string           `json:"netEquityLocked"`
	NetExposureFutures string           `json:"netExposureFutures"`
	PnlUnrealized      string           `json:"pnlUnrealized"`
}

// CollateralItem represents a single collateral item.
type CollateralItem struct {
	Symbol            enums.CustodyAsset `json:"symbol"`
	AssetMarkPrice    string `json:"assetMarkPrice"`
	TotalQuantity     string `json:"totalQuantity"`
	BalanceNotional   string `json:"balanceNotional"`
	CollateralWeight  string `json:"collateralWeight"`
	CollateralValue   string `json:"collateralValue"`
	OpenOrderQuantity string `json:"openOrderQuantity"`
	LendQuantity      string `json:"lendQuantity"`
	AvailableQuantity string `json:"availableQuantity"`
}

// Deposit represents a deposit record from GET /wapi/v1/capital/deposits.
type Deposit struct {
	ID              int32                `json:"id"`
	ToAddress       string               `json:"toAddress,omitempty"`
	FromAddress     string               `json:"fromAddress,omitempty"`
	Source          enums.DepositSource  `json:"source"`
	Status          enums.DepositStatus  `json:"status"`
	TransactionHash string               `json:"transactionHash,omitempty"`
	Symbol          enums.CustodyAsset   `json:"symbol"`
	Quantity        string               `json:"quantity"`
	CreatedAt       string               `json:"createdAt"`
	FiatAmount      float64              `json:"fiatAmount,omitempty"`
	FiatCurrency    enums.FiatAsset      `json:"fiatCurrency,omitempty"`
	InstitutionBic  string               `json:"institutionBic,omitempty"`
	PlatformMemo    string               `json:"platformMemo,omitempty"`
}

// DepositAddress represents a deposit address from GET /wapi/v1/capital/deposit/address.
type DepositAddress struct {
	Address string `json:"address"`
}

// Withdrawal represents a withdrawal record from GET /wapi/v1/capital/withdrawals.
type Withdrawal struct {
	ID                int32                           `json:"id"`
	Blockchain        enums.Blockchain                `json:"blockchain"`
	ClientID          string                          `json:"clientId,omitempty"`
	Identifier        string                          `json:"identifier,omitempty"`
	Quantity          string                          `json:"quantity"`
	Fee               string                          `json:"fee"`
	FiatFee           string                          `json:"fiatFee,omitempty"`
	FiatState         enums.EqualsMoneyWithdrawalState `json:"fiatState,omitempty"`
	FiatSymbol        enums.FiatAsset                 `json:"fiatSymbol,omitempty"`
	ProviderID        string                          `json:"providerId,omitempty"`
	Symbol            enums.CustodyAsset              `json:"symbol"`
	Status            enums.WithdrawalStatus          `json:"status"`
	SubaccountID      uint16                          `json:"subaccountId,omitempty"`
	ToAddress         string                          `json:"toAddress"`
	TransactionHash   string                          `json:"transactionHash,omitempty"`
	CreatedAt         string                          `json:"createdAt"`
	IsInternal        bool                            `json:"isInternal"`
	BankName          string                          `json:"bankName,omitempty"`
	BankIdentifier    string                          `json:"bankIdentifier,omitempty"`
	AccountIdentifier string                          `json:"accountIdentifier,omitempty"`
	TriggerAt         string                          `json:"triggerAt,omitempty"`
}

// WithdrawalRequest represents a withdrawal request for POST /wapi/v1/capital/withdrawals.
type WithdrawalRequest struct {
	Address              string                          `json:"address"`
	Blockchain           enums.Blockchain               `json:"blockchain"`
	ClientID             string                          `json:"clientId,omitempty"`
	Quantity             string                          `json:"quantity"`
	Symbol               enums.CustodyAsset             `json:"symbol"`
	TwoFactorToken       string                          `json:"twoFactorToken,omitempty"`
	AutoBorrow           *bool                           `json:"autoBorrow,omitempty"`
	AutoLendRedeem       *bool                           `json:"autoLendRedeem,omitempty"`
	RecipientInformation *WithdrawalRecipientInformation `json:"recipientInformation,omitempty"`
}

// WithdrawalRecipientInformation represents recipient details for fiat withdrawals.
type WithdrawalRecipientInformation struct {
	WithdrawalAddressID     int32  `json:"withdrawal_address_id"`
	WithdrawalPurpose       string `json:"withdrawal_purpose,omitempty"`
	SanctionsRepresentation *bool  `json:"sanctions_representation,omitempty"`
}

// WithdrawalDelay represents withdrawal delay settings from /wapi/v1/capital/withdrawals/delay.
type WithdrawalDelay struct {
	CurrentWithdrawalDelayHours         int32  `json:"currentWithdrawalDelayHours,omitempty"`
	PendingWithdrawalDelayHours         int32  `json:"pendingWithdrawalDelayHours,omitempty"`
	PendingWithdrawalDelayHoursEnabledAt string `json:"pendingWithdrawalDelayHoursEnabledAt,omitempty"`
}

// DustConversion represents dust conversion result from history API.
type DustConversion struct {
	ID           uint64 `json:"id"`
	Quantity     string `json:"quantity"`
	Symbol       string `json:"symbol"`
	USDCReceived string `json:"usdcReceived"`
	Timestamp    string `json:"timestamp"`
}

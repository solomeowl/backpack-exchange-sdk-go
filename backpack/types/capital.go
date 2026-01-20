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

// CollateralResponse represents the collateral API response.
type CollateralResponse struct {
	AssetsValue      string           `json:"assetsValue"`
	BorrowLiability  string           `json:"borrowLiability"`
	Collateral       []CollateralItem `json:"collateral"`
}

// CollateralItem represents a single collateral item.
type CollateralItem struct {
	Symbol            string `json:"symbol"`
	AssetMarkPrice    string `json:"assetMarkPrice"`
	AvailableQuantity string `json:"availableQuantity"`
	BalanceNotional   string `json:"balanceNotional"`
	CollateralValue   string `json:"collateralValue"`
	CollateralWeight  string `json:"collateralWeight"`
	LendQuantity      string `json:"lendQuantity"`
	OpenOrderQuantity string `json:"openOrderQuantity"`
	TotalQuantity     string `json:"totalQuantity"`
}

// Deposit represents a deposit record.
type Deposit struct {
	ID              int64               `json:"id"`
	Symbol          string              `json:"symbol"`
	Quantity        string              `json:"quantity"`
	Status          enums.DepositStatus `json:"status"`
	Blockchain      enums.Blockchain    `json:"blockchain"`
	TransactionHash string              `json:"transactionHash,omitempty"`
	FromAddress     string              `json:"fromAddress,omitempty"`
	Confirmations   int                 `json:"confirmations,omitempty"`
	CreatedAt       string              `json:"createdAt,omitempty"`
}

// DepositAddress represents a deposit address.
type DepositAddress struct {
	Address    string           `json:"address"`
	Blockchain enums.Blockchain `json:"blockchain"`
	Tag        string           `json:"tag,omitempty"`
}

// Withdrawal represents a withdrawal record.
type Withdrawal struct {
	ID              int64                  `json:"id"`
	Symbol          string                 `json:"symbol"`
	Quantity        string                 `json:"quantity"`
	Fee             string                 `json:"fee,omitempty"`
	Status          enums.WithdrawalStatus `json:"status"`
	Blockchain      enums.Blockchain       `json:"blockchain"`
	ToAddress       string                 `json:"toAddress"`
	TransactionHash string                 `json:"transactionHash,omitempty"`
	CreatedAt       string                 `json:"createdAt,omitempty"`
}

// WithdrawalRequest represents a withdrawal request.
type WithdrawalRequest struct {
	Symbol        string           `json:"symbol"`
	Blockchain    enums.Blockchain `json:"blockchain"`
	Address       string           `json:"address"`
	Quantity      string           `json:"quantity"`
	TwoFactorCode string           `json:"twoFactorCode,omitempty"`
	ClientID      string           `json:"clientId,omitempty"`
}

// WithdrawalDelay represents withdrawal delay settings.
type WithdrawalDelay struct {
	Enabled         bool   `json:"enabled"`
	DelayHours      int    `json:"delayHours,omitempty"`
	AddressWhitelist []string `json:"addressWhitelist,omitempty"`
}

// DustConversion represents dust conversion result.
type DustConversion struct {
	ConvertedSymbol string `json:"convertedSymbol"`
	ConvertedAmount string `json:"convertedAmount"`
	ReceivedSymbol  string `json:"receivedSymbol"`
	ReceivedAmount  string `json:"receivedAmount"`
}

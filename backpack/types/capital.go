package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Balance represents an account balance.
type Balance struct {
	Symbol    string `json:"symbol"`
	Available string `json:"available"`
	Locked    string `json:"locked"`
	Staked    string `json:"staked,omitempty"`
	Total     string `json:"total,omitempty"`
}

// Collateral represents collateral information.
type Collateral struct {
	Symbol      string `json:"symbol"`
	Quantity    string `json:"quantity"`
	Value       string `json:"value,omitempty"`
	Weight      string `json:"weight,omitempty"`
}

// Deposit represents a deposit record.
type Deposit struct {
	ID              string              `json:"id"`
	Symbol          string              `json:"symbol"`
	Quantity        string              `json:"quantity"`
	Status          enums.DepositStatus `json:"status"`
	Blockchain      enums.Blockchain    `json:"blockchain"`
	TransactionHash string              `json:"transactionHash,omitempty"`
	FromAddress     string              `json:"fromAddress,omitempty"`
	Confirmations   int                 `json:"confirmations,omitempty"`
	CreatedAt       int64               `json:"createdAt,omitempty"`
}

// DepositAddress represents a deposit address.
type DepositAddress struct {
	Address    string           `json:"address"`
	Blockchain enums.Blockchain `json:"blockchain"`
	Tag        string           `json:"tag,omitempty"`
}

// Withdrawal represents a withdrawal record.
type Withdrawal struct {
	ID              string                 `json:"id"`
	Symbol          string                 `json:"symbol"`
	Quantity        string                 `json:"quantity"`
	Fee             string                 `json:"fee,omitempty"`
	Status          enums.WithdrawalStatus `json:"status"`
	Blockchain      enums.Blockchain       `json:"blockchain"`
	ToAddress       string                 `json:"toAddress"`
	TransactionHash string                 `json:"transactionHash,omitempty"`
	CreatedAt       int64                  `json:"createdAt,omitempty"`
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

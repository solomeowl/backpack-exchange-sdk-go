package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// SystemStatus represents the system status response.
type SystemStatus struct {
	Status  enums.SystemStatus `json:"status"`
	Message string             `json:"message,omitempty"`
}

// ServerTime represents the server time response.
type ServerTime struct {
	ServerTime int64 `json:"serverTime"`
}

// Wallet represents a wallet address for deposits.
type Wallet struct {
	Address    string           `json:"address"`
	Blockchain enums.Blockchain `json:"blockchain"`
}

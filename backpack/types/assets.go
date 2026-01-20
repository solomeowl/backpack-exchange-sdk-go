package types

// Asset represents an asset/token.
type Asset struct {
	Symbol      string        `json:"symbol"`
	Name        string        `json:"name,omitempty"`
	Decimals    int           `json:"decimals"`
	TokenMint   string        `json:"tokenMint,omitempty"`
	Blockchains []AssetChain  `json:"blockchains,omitempty"`
}

// AssetChain represents blockchain-specific asset information.
type AssetChain struct {
	Blockchain        string `json:"blockchain"`
	DepositEnabled    bool   `json:"depositEnabled"`
	WithdrawEnabled   bool   `json:"withdrawEnabled"`
	MinWithdraw       string `json:"minWithdraw,omitempty"`
	WithdrawFee       string `json:"withdrawFee,omitempty"`
	ContractAddress   string `json:"contractAddress,omitempty"`
	Confirmations     int    `json:"confirmations,omitempty"`
}

// CollateralInfo represents collateral information for an asset.
type CollateralInfo struct {
	Symbol           string `json:"symbol"`
	InitialWeight    string `json:"initialWeight"`
	MaintenanceWeight string `json:"maintenanceWeight"`
	MaxQuantity      string `json:"maxQuantity,omitempty"`
}

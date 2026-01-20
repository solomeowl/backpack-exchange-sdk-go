package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Asset represents a market asset from GET /api/v1/assets.
type Asset struct {
	Symbol      enums.CustodyAsset `json:"symbol"`
	DisplayName string  `json:"displayName"`
	CoingeckoID string  `json:"coingeckoId,omitempty"`
	Tokens      []Token `json:"tokens"`
}

// Token represents blockchain-specific token information.
type Token struct {
	DisplayName       string           `json:"displayName"`
	Blockchain        enums.Blockchain `json:"blockchain"`
	ContractAddress   string           `json:"contractAddress,omitempty"`
	DepositEnabled    bool             `json:"depositEnabled"`
	MinimumDeposit    string           `json:"minimumDeposit"`
	WithdrawEnabled   bool             `json:"withdrawEnabled"`
	MinimumWithdrawal string           `json:"minimumWithdrawal"`
	MaximumWithdrawal string           `json:"maximumWithdrawal,omitempty"`
	WithdrawalFee     string           `json:"withdrawalFee"`
}

// CollateralInfo represents collateral parameters from GET /api/v1/collateral.
type CollateralInfo struct {
	Symbol          enums.CustodyAsset  `json:"symbol"`
	ImfFunction     PositionImfFunction `json:"imfFunction"`
	MmfFunction     PositionImfFunction `json:"mmfFunction"`
	HaircutFunction CollateralFunction  `json:"haircutFunction"`
}

// PositionImfFunction represents the IMF/MMF function configuration.
type PositionImfFunction struct {
	Type   string `json:"type"`
	A      string `json:"a,omitempty"`
	B      string `json:"b,omitempty"`
	C      string `json:"c,omitempty"`
	Floor  string `json:"floor,omitempty"`
}

// CollateralFunction represents the collateral haircut function.
type CollateralFunction struct {
	Weight string                 `json:"weight"`
	Kind   CollateralFunctionKind `json:"kind"`
}

// CollateralFunctionKind represents the kind of collateral function.
type CollateralFunctionKind struct {
	Type   string `json:"type"`
	A      string `json:"a,omitempty"`
	B      string `json:"b,omitempty"`
	C      string `json:"c,omitempty"`
	Floor  string `json:"floor,omitempty"`
}

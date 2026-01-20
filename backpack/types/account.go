package types

// Account represents account information.
type Account struct {
	AccountType          string `json:"accountType,omitempty"`
	TotalEquity          string `json:"totalEquity,omitempty"`
	TotalCollateral      string `json:"totalCollateral,omitempty"`
	AvailableCollateral  string `json:"availableCollateral,omitempty"`
	InitialMargin        string `json:"initialMargin,omitempty"`
	MaintenanceMargin    string `json:"maintenanceMargin,omitempty"`
	UnrealizedPnl        string `json:"unrealizedPnl,omitempty"`
	RealizedPnl          string `json:"realizedPnl,omitempty"`
	Leverage             string `json:"leverage,omitempty"`
	AutoBorrow           bool   `json:"autoBorrow,omitempty"`
	AutoLend             bool   `json:"autoLend,omitempty"`
	AutoRepay            bool   `json:"autoRepay,omitempty"`
	AutoRedeem           bool   `json:"autoRedeem,omitempty"`
}

// AccountSettings represents account settings for update.
type AccountSettings struct {
	AutoBorrow bool `json:"autoBorrow,omitempty"`
	AutoLend   bool `json:"autoLend,omitempty"`
	AutoRepay  bool `json:"autoRepay,omitempty"`
	AutoRedeem bool `json:"autoRedeem,omitempty"`
}

// MaxQuantityResponse represents max quantity query response.
type MaxQuantityResponse struct {
	MaxQuantity string `json:"maxQuantity"`
}

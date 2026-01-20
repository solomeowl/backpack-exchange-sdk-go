package types

// Account represents account summary from GET /api/v1/account.
type Account struct {
	AutoBorrowSettlements bool   `json:"autoBorrowSettlements"`
	AutoLend              bool   `json:"autoLend"`
	AutoRealizePnl        bool   `json:"autoRealizePnl"`
	AutoRepayBorrows      bool   `json:"autoRepayBorrows"`
	BorrowLimit           string `json:"borrowLimit"`
	FuturesMakerFee       string `json:"futuresMakerFee"`
	FuturesTakerFee       string `json:"futuresTakerFee"`
	LeverageLimit         string `json:"leverageLimit"`
	LimitOrders           int64  `json:"limitOrders"`
	Liquidating           bool   `json:"liquidating"`
	PositionLimit         string `json:"positionLimit"`
	SpotMakerFee          string `json:"spotMakerFee"`
	SpotTakerFee          string `json:"spotTakerFee"`
	TriggerOrders         int64  `json:"triggerOrders"`
}

// AccountSettings represents account settings for PATCH /api/v1/account.
type AccountSettings struct {
	AutoBorrowSettlements *bool `json:"autoBorrowSettlements,omitempty"`
	AutoLend              *bool `json:"autoLend,omitempty"`
	AutoRealizePnl        *bool `json:"autoRealizePnl,omitempty"`
	AutoRepayBorrows      *bool `json:"autoRepayBorrows,omitempty"`
}

// MaxQuantityResponse represents max quantity query response.
type MaxQuantityResponse struct {
	MaxQuantity string `json:"maxQuantity"`
}

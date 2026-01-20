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
	AutoRepayBorrows      *bool `json:"autoRepayBorrows,omitempty"`
	LeverageLimit         string `json:"leverageLimit,omitempty"`
}

// MaxBorrowQuantity represents max borrow quantity query response.
type MaxBorrowQuantity struct {
	MaxBorrowQuantity string `json:"maxBorrowQuantity"`
	Symbol            string `json:"symbol"`
}

// MaxOrderQuantity represents max order quantity query response.
type MaxOrderQuantity struct {
	AutoBorrow      *bool  `json:"autoBorrow,omitempty"`
	AutoBorrowRepay *bool  `json:"autoBorrowRepay,omitempty"`
	AutoLendRedeem  *bool  `json:"autoLendRedeem,omitempty"`
	MaxOrderQuantity string `json:"maxOrderQuantity"`
	Price           string `json:"price,omitempty"`
	Side            string `json:"side"`
	Symbol          string `json:"symbol"`
	ReduceOnly      *bool  `json:"reduceOnly,omitempty"`
}

// MaxWithdrawalQuantity represents max withdrawal quantity query response.
type MaxWithdrawalQuantity struct {
	AutoBorrow          *bool  `json:"autoBorrow,omitempty"`
	AutoLendRedeem      *bool  `json:"autoLendRedeem,omitempty"`
	MaxWithdrawalQuantity string `json:"maxWithdrawalQuantity"`
	Symbol              string `json:"symbol"`
}

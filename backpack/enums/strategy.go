package enums

// StrategyType represents the type of strategy.
type StrategyType string

const (
	StrategyTypeScheduled StrategyType = "Scheduled"
)

// StrategyStatus represents the status of a strategy.
type StrategyStatus string

const (
	StrategyStatusRunning    StrategyStatus = "Running"
	StrategyStatusCompleted  StrategyStatus = "Completed"
	StrategyStatusCancelled  StrategyStatus = "Cancelled"
	StrategyStatusTerminated StrategyStatus = "Terminated"
)

// StrategyCancelReason represents the reason a strategy was cancelled.
type StrategyCancelReason string

const (
	StrategyCancelReasonExpired                       StrategyCancelReason = "Expired"
	StrategyCancelReasonFillOrKill                    StrategyCancelReason = "FillOrKill"
	StrategyCancelReasonInsufficientBorrowableQuantity StrategyCancelReason = "InsufficientBorrowableQuantity"
	StrategyCancelReasonInsufficientFunds             StrategyCancelReason = "InsufficientFunds"
	StrategyCancelReasonInsufficientLiquidity         StrategyCancelReason = "InsufficientLiquidity"
	StrategyCancelReasonInvalidPrice                  StrategyCancelReason = "InvalidPrice"
	StrategyCancelReasonInvalidQuantity               StrategyCancelReason = "InvalidQuantity"
	StrategyCancelReasonInsufficientMargin            StrategyCancelReason = "InsufficientMargin"
	StrategyCancelReasonLiquidation                   StrategyCancelReason = "Liquidation"
	StrategyCancelReasonPriceOutOfBounds              StrategyCancelReason = "PriceOutOfBounds"
	StrategyCancelReasonReduceOnlyNotReduced          StrategyCancelReason = "ReduceOnlyNotReduced"
	StrategyCancelReasonSelfTradePrevention           StrategyCancelReason = "SelfTradePrevention"
	StrategyCancelReasonUnknown                       StrategyCancelReason = "Unknown"
	StrategyCancelReasonUserPermissions               StrategyCancelReason = "UserPermissions"
)

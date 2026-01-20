package enums

// StrategyType represents the type of strategy.
type StrategyType string

const (
	StrategyTypeScheduled StrategyType = "Scheduled"
)

// StrategyTypeEnum matches the OpenAPI enum name.
type StrategyTypeEnum = StrategyType

// StrategyStatus represents the status of a strategy.
type StrategyStatus string

const (
	StrategyStatusRunning    StrategyStatus = "Running"
	StrategyStatusCompleted  StrategyStatus = "Completed"
	StrategyStatusCancelled  StrategyStatus = "Cancelled"
	StrategyStatusTerminated StrategyStatus = "Terminated"
)

// StrategyCrankCancelReason represents crank cancellation reasons for strategies.
type StrategyCrankCancelReason string

const (
	StrategyCrankCancelReasonExpired                       StrategyCrankCancelReason = "Expired"
	StrategyCrankCancelReasonFillOrKill                    StrategyCrankCancelReason = "FillOrKill"
	StrategyCrankCancelReasonInsufficientBorrowableQuantity StrategyCrankCancelReason = "InsufficientBorrowableQuantity"
	StrategyCrankCancelReasonInsufficientFunds             StrategyCrankCancelReason = "InsufficientFunds"
	StrategyCrankCancelReasonInsufficientLiquidity         StrategyCrankCancelReason = "InsufficientLiquidity"
	StrategyCrankCancelReasonInvalidPrice                  StrategyCrankCancelReason = "InvalidPrice"
	StrategyCrankCancelReasonInvalidQuantity               StrategyCrankCancelReason = "InvalidQuantity"
	StrategyCrankCancelReasonInsufficientMargin            StrategyCrankCancelReason = "InsufficientMargin"
	StrategyCrankCancelReasonLiquidation                   StrategyCrankCancelReason = "Liquidation"
	StrategyCrankCancelReasonPriceOutOfBounds              StrategyCrankCancelReason = "PriceOutOfBounds"
	StrategyCrankCancelReasonReduceOnlyNotReduced          StrategyCrankCancelReason = "ReduceOnlyNotReduced"
	StrategyCrankCancelReasonSelfTradePrevention           StrategyCrankCancelReason = "SelfTradePrevention"
	StrategyCrankCancelReasonUnknown                       StrategyCrankCancelReason = "Unknown"
	StrategyCrankCancelReasonUserPermissions               StrategyCrankCancelReason = "UserPermissions"
)

// StrategyCancelReason is kept for backwards compatibility with previous SDK versions.
type StrategyCancelReason = StrategyCrankCancelReason

package enums

// StrategyType represents the type of strategy.
type StrategyType string

const (
	StrategyTypeTWAP StrategyType = "TWAP"
)

// StrategyStatus represents the status of a strategy.
type StrategyStatus string

const (
	StrategyStatusActive    StrategyStatus = "Active"
	StrategyStatusCompleted StrategyStatus = "Completed"
	StrategyStatusCancelled StrategyStatus = "Cancelled"
	StrategyStatusFailed    StrategyStatus = "Failed"
)

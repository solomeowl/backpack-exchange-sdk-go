package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Strategy represents a trading strategy.
type Strategy struct {
	ID               string               `json:"id"`
	Symbol           string               `json:"symbol"`
	Side             enums.Side           `json:"side"`
	StrategyType     enums.StrategyType   `json:"strategyType"`
	Status           enums.StrategyStatus `json:"status"`
	TotalQuantity    string               `json:"totalQuantity"`
	ExecutedQuantity string               `json:"executedQuantity,omitempty"`
	StartTime        int64                `json:"startTime,omitempty"`
	EndTime          int64                `json:"endTime,omitempty"`
	CreatedAt        int64                `json:"createdAt,omitempty"`
}

// CreateStrategyParams represents parameters for creating a strategy.
type CreateStrategyParams struct {
	Symbol        string             `json:"symbol"`
	Side          enums.Side         `json:"side"`
	StrategyType  enums.StrategyType `json:"strategyType"`
	TotalQuantity string             `json:"totalQuantity"`
	Duration      int64              `json:"duration,omitempty"` // Duration in seconds
	PriceLimit    string             `json:"priceLimit,omitempty"`
}

// StrategyHistoryItem represents a strategy history record.
type StrategyHistoryItem struct {
	Strategy
	Fills []Fill `json:"fills,omitempty"`
}

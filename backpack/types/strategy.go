package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Strategy represents a trading strategy from GET /api/v1/strategy.
type Strategy struct {
	ID                         string                            `json:"id"`
	CreatedAt                  string                            `json:"createdAt"`
	StrategyType               enums.StrategyType                `json:"strategyType"`
	SelfTradePrevention        enums.SelfTradePrevention         `json:"selfTradePrevention"`
	Status                     enums.StrategyStatus              `json:"status"`
	Side                       enums.Side                        `json:"side"`
	Symbol                     string                            `json:"symbol"`
	TimeInForce                enums.TimeInForce                 `json:"timeInForce"`
	Duration                   uint64                            `json:"duration"`
	Interval                   uint64                            `json:"interval"`
	RandomizedIntervalQuantity bool                              `json:"randomizedIntervalQuantity"`
	ExecutedQuantity           string                            `json:"executedQuantity,omitempty"`
	ExecutedQuoteQuantity      string                            `json:"executedQuoteQuantity,omitempty"`
	CancelReason               enums.StrategyCancelReason        `json:"cancelReason,omitempty"`
	Quantity                   string                            `json:"quantity,omitempty"`
	ClientStrategyID           uint32                            `json:"clientStrategyId,omitempty"`
	SlippageTolerance          string                            `json:"slippageTolerance,omitempty"`
	SlippageToleranceType      enums.SlippageToleranceType       `json:"slippageToleranceType,omitempty"`
}

// CreateStrategyParams represents parameters for creating a strategy via POST /api/v1/strategy.
type CreateStrategyParams struct {
	Symbol                     string                       `json:"symbol"`
	Side                       enums.Side                   `json:"side"`
	StrategyType               enums.StrategyType           `json:"strategyType"`
	Quantity                   string                       `json:"quantity,omitempty"`
	Price                      string                       `json:"price,omitempty"`
	Duration                   uint64                       `json:"duration,omitempty"`
	Interval                   uint64                       `json:"interval,omitempty"`
	RandomizedIntervalQuantity *bool                        `json:"randomizedIntervalQuantity,omitempty"`
	TimeInForce                enums.TimeInForce            `json:"timeInForce,omitempty"`
	PostOnly                   *bool                        `json:"postOnly,omitempty"`
	ReduceOnly                 *bool                        `json:"reduceOnly,omitempty"`
	SelfTradePrevention        enums.SelfTradePrevention    `json:"selfTradePrevention,omitempty"`
	ClientStrategyID           *uint32                      `json:"clientStrategyId,omitempty"`
	BrokerID                   *uint16                      `json:"brokerId,omitempty"`
	AutoBorrow                 *bool                        `json:"autoBorrow,omitempty"`
	AutoBorrowRepay            *bool                        `json:"autoBorrowRepay,omitempty"`
	AutoLend                   *bool                        `json:"autoLend,omitempty"`
	AutoLendRedeem             *bool                        `json:"autoLendRedeem,omitempty"`
	SlippageTolerance          string                       `json:"slippageTolerance,omitempty"`
	SlippageToleranceType      enums.SlippageToleranceType  `json:"slippageToleranceType,omitempty"`
}

// StrategyHistoryItem represents a strategy history record.
type StrategyHistoryItem struct {
	Strategy
	Fills []Fill `json:"fills,omitempty"`
}

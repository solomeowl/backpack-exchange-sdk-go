package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Position represents a futures position.
type Position struct {
	Symbol            string             `json:"symbol"`
	Side              enums.PositionSide `json:"side,omitempty"`
	Quantity          string             `json:"quantity"`
	EntryPrice        string             `json:"entryPrice,omitempty"`
	MarkPrice         string             `json:"markPrice,omitempty"`
	LiquidationPrice  string             `json:"liquidationPrice,omitempty"`
	UnrealizedPnl     string             `json:"unrealizedPnl,omitempty"`
	RealizedPnl       string             `json:"realizedPnl,omitempty"`
	Leverage          string             `json:"leverage,omitempty"`
	MarginType        enums.MarginType   `json:"marginType,omitempty"`
	InitialMargin     string             `json:"initialMargin,omitempty"`
	MaintenanceMargin string             `json:"maintenanceMargin,omitempty"`
}

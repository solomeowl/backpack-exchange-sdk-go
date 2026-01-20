package types

// Position represents a futures position from GET /api/v1/position (FuturePositionWithMargin).
type Position struct {
	BreakEvenPrice           string              `json:"breakEvenPrice"`
	EntryPrice               string              `json:"entryPrice"`
	EstLiquidationPrice      string              `json:"estLiquidationPrice"`
	IMF                      string              `json:"imf"`
	IMFFunction              PositionImfFunction `json:"imfFunction"`
	MarkPrice                string              `json:"markPrice"`
	MMF                      string              `json:"mmf"`
	MMFFunction              PositionImfFunction `json:"mmfFunction"`
	NetCost                  string              `json:"netCost"`
	NetQuantity              string              `json:"netQuantity"`
	NetExposureQuantity      string              `json:"netExposureQuantity"`
	NetExposureNotional      string              `json:"netExposureNotional"`
	PnlRealized              string              `json:"pnlRealized"`
	PnlUnrealized            string              `json:"pnlUnrealized"`
	CumulativeFundingPayment string              `json:"cumulativeFundingPayment"`
	SubaccountID             uint16              `json:"subaccountId,omitempty"`
	Symbol                   string              `json:"symbol"`
	UserID                   int32               `json:"userId"`
	PositionID               string              `json:"positionId"`
	CumulativeInterest       string              `json:"cumulativeInterest"`
}

// Note: PositionImfFunction is defined in assets.go

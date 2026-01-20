package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// BorrowLendMarket represents a borrow/lend market from GET /api/v1/borrowLend/markets.
type BorrowLendMarket struct {
	Symbol                        string `json:"symbol"`
	State                         string `json:"state"`
	AssetMarkPrice                string `json:"assetMarkPrice"`
	BorrowInterestRate            string `json:"borrowInterestRate"`
	BorrowedQuantity              string `json:"borrowedQuantity"`
	Fee                           string `json:"fee"`
	LendInterestRate              string `json:"lendInterestRate"`
	LentQuantity                  string `json:"lentQuantity"`
	MaxUtilization                string `json:"maxUtilization"`
	OpenBorrowLendLimit           string `json:"openBorrowLendLimit"`
	OptimalUtilization            string `json:"optimalUtilization"`
	Timestamp                     string `json:"timestamp"`
	ThrottleUtilizationThreshold  string `json:"throttleUtilizationThreshold"`
	ThrottleUtilizationBound      string `json:"throttleUtilizationBound"`
	ThrottleUpdateFraction        string `json:"throttleUpdateFraction"`
	Utilization                   string `json:"utilization"`
	StepSize                      string `json:"stepSize"`
}

// BorrowLendMarketHistory represents historical borrow/lend market data from GET /api/v1/borrowLend/markets/history.
type BorrowLendMarketHistory struct {
	BorrowInterestRate string `json:"borrowInterestRate"`
	BorrowedQuantity   string `json:"borrowedQuantity"`
	LendInterestRate   string `json:"lendInterestRate"`
	LentQuantity       string `json:"lentQuantity"`
	Timestamp          string `json:"timestamp"`
	Utilization        string `json:"utilization"`
}

// BorrowLendPosition represents a borrow/lend position from GET /api/v1/borrowLend/positions.
type BorrowLendPosition struct {
	CumulativeInterest   string              `json:"cumulativeInterest"`
	ID                   string              `json:"id"`
	IMF                  string              `json:"imf"`
	IMFFunction          PositionImfFunction `json:"imfFunction"`
	NetQuantity          string              `json:"netQuantity"`
	MarkPrice            string              `json:"markPrice"`
	MMF                  string              `json:"mmf"`
	MMFFunction          PositionImfFunction `json:"mmfFunction"`
	NetExposureQuantity  string              `json:"netExposureQuantity"`
	NetExposureNotional  string              `json:"netExposureNotional"`
	Symbol               string              `json:"symbol"`
}

// BorrowLendExecuteParams represents parameters for borrow/lend execution.
type BorrowLendExecuteParams struct {
	Symbol   string               `json:"symbol"`
	Side     enums.BorrowLendSide `json:"side"`
	Quantity string               `json:"quantity"`
}

// EstimatedLiquidationPrice represents estimated liquidation price from GET /api/v1/borrowLend/position/liquidationPrice.
type EstimatedLiquidationPrice struct {
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
}

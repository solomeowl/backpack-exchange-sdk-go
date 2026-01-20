package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// BorrowLendMarket represents a borrow/lend market.
type BorrowLendMarket struct {
	Symbol          string `json:"symbol"`
	BorrowRate      string `json:"borrowRate,omitempty"`
	LendRate        string `json:"lendRate,omitempty"`
	UtilizationRate string `json:"utilizationRate,omitempty"`
	TotalBorrowed   string `json:"totalBorrowed,omitempty"`
	TotalLent       string `json:"totalLent,omitempty"`
	MaxBorrow       string `json:"maxBorrow,omitempty"`
}

// BorrowLendMarketHistory represents historical borrow/lend market data.
type BorrowLendMarketHistory struct {
	Symbol     string `json:"symbol"`
	BorrowRate string `json:"borrowRate"`
	LendRate   string `json:"lendRate"`
	Timestamp  string `json:"timestamp"`
}

// BorrowLendPosition represents a borrow/lend position.
type BorrowLendPosition struct {
	Symbol          string               `json:"symbol"`
	Side            enums.BorrowLendSide `json:"side"`
	Quantity        string               `json:"quantity"`
	AccruedInterest string               `json:"accruedInterest,omitempty"`
	EntryRate       string               `json:"entryRate,omitempty"`
	Timestamp       int64                `json:"timestamp,omitempty"`
}

// BorrowLendExecuteParams represents parameters for borrow/lend execution.
type BorrowLendExecuteParams struct {
	Symbol   string               `json:"symbol"`
	Side     enums.BorrowLendSide `json:"side"`
	Quantity string               `json:"quantity"`
}

// EstimatedLiquidationPrice represents estimated liquidation price.
type EstimatedLiquidationPrice struct {
	Symbol           string `json:"symbol"`
	LiquidationPrice string `json:"liquidationPrice"`
}

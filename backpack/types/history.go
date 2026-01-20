package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Fill represents a trade fill.
type Fill struct {
	ID            string     `json:"id"`
	OrderID       string     `json:"orderId,omitempty"`
	Symbol        string     `json:"symbol"`
	Side          enums.Side `json:"side"`
	Price         string     `json:"price"`
	Quantity      string     `json:"quantity"`
	QuoteQuantity string     `json:"quoteQuantity,omitempty"`
	Fee           string     `json:"fee,omitempty"`
	FeeSymbol     string     `json:"feeSymbol,omitempty"`
	IsMaker       bool       `json:"isMaker"`
	Timestamp     string     `json:"timestamp"`
}

// FundingPayment represents a funding payment record.
type FundingPayment struct {
	Symbol      string `json:"symbol"`
	Payment     string `json:"payment"`
	FundingRate string `json:"fundingRate,omitempty"`
	MarkPrice   string `json:"markPrice,omitempty"`
	Timestamp   string `json:"timestamp"`
}

// Settlement represents a settlement record.
type Settlement struct {
	Symbol       string `json:"symbol"`
	SettlePrice  string `json:"settlePrice"`
	Quantity     string `json:"quantity"`
	PnL          string `json:"pnl"`
	Timestamp    string `json:"timestamp"`
}

// BorrowHistoryItem represents a borrow history record.
type BorrowHistoryItem struct {
	ID        string                    `json:"id"`
	Symbol    string                    `json:"symbol"`
	Quantity  string                    `json:"quantity"`
	EventType enums.BorrowLendEventType `json:"eventType"`
	Timestamp string                    `json:"timestamp"`
}

// InterestHistoryItem represents an interest payment record.
type InterestHistoryItem struct {
	Symbol    string `json:"symbol"`
	Interest  string `json:"interest"`
	Rate      string `json:"rate,omitempty"`
	Timestamp string `json:"timestamp"`
}

// DustHistoryItem represents a dust conversion record.
type DustHistoryItem struct {
	FromSymbol string `json:"fromSymbol"`
	FromAmount string `json:"fromAmount"`
	ToSymbol   string `json:"toSymbol"`
	ToAmount   string `json:"toAmount"`
	Timestamp  string `json:"timestamp"`
}

// HistoryParams represents common parameters for history queries.
type HistoryParams struct {
	Symbol    string              `json:"symbol,omitempty"`
	From      int64               `json:"from,omitempty"`
	To        int64               `json:"to,omitempty"`
	Limit     int                 `json:"limit,omitempty"`
	Offset    int                 `json:"offset,omitempty"`
	Direction enums.SortDirection `json:"direction,omitempty"`
}

// OrderHistoryItem represents an order history record.
type OrderHistoryItem struct {
	Order
	Fills []Fill `json:"fills,omitempty"`
}

// PositionHistoryItem represents a position history record.
type PositionHistoryItem struct {
	Symbol        string             `json:"symbol"`
	Side          enums.PositionSide `json:"side"`
	EntryPrice    string             `json:"entryPrice"`
	ExitPrice     string             `json:"exitPrice,omitempty"`
	Quantity      string             `json:"quantity"`
	RealizedPnl   string             `json:"realizedPnl,omitempty"`
	OpenTimestamp int64              `json:"openTimestamp"`
	CloseTimestamp int64             `json:"closeTimestamp,omitempty"`
}

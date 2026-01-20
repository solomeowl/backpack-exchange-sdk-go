package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Fill represents a trade fill from OrderFill schema.
type Fill struct {
	ClientID        string                `json:"clientId,omitempty"`
	Fee             string                `json:"fee"`
	FeeSymbol       string                `json:"feeSymbol"`
	IsMaker         bool                  `json:"isMaker"`
	OrderID         string                `json:"orderId"`
	Price           string                `json:"price"`
	Quantity        string                `json:"quantity"`
	Side            enums.Side            `json:"side"`
	Symbol          string                `json:"symbol"`
	SystemOrderType enums.SystemOrderType `json:"systemOrderType,omitempty"`
	Timestamp       string                `json:"timestamp"`
	TradeID         int64                 `json:"tradeId,omitempty"`
}

// FundingPayment represents a funding payment record.
type FundingPayment struct {
	UserID               int32  `json:"userId"`
	SubaccountID         uint16 `json:"subaccountId,omitempty"`
	Symbol               string `json:"symbol"`
	Quantity             string `json:"quantity"`
	IntervalEndTimestamp string `json:"intervalEndTimestamp"`
	FundingRate          string `json:"fundingRate"`
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
	PaymentType   enums.PaymentType `json:"paymentType"`
	InterestRate  string            `json:"interestRate"`
	Interval      string            `json:"interval"`
	MarketSymbol  string            `json:"marketSymbol"`
	PositionID    string            `json:"positionId"`
	Quantity      string            `json:"quantity"`
	Symbol        string            `json:"symbol"`
	Timestamp     string            `json:"timestamp"`
}

// DustHistoryItem represents a dust conversion record (same as DustConversion).
type DustHistoryItem = DustConversion

// HistoryParams represents common parameters for history queries.
type HistoryParams struct {
	Symbol    string              `json:"symbol,omitempty"`
	From      int64               `json:"from,omitempty"`
	To        int64               `json:"to,omitempty"`
	Limit     int                 `json:"limit,omitempty"`
	Offset    int                 `json:"offset,omitempty"`
	Direction enums.SortDirection `json:"direction,omitempty"`
}

// OrderHistoryItem represents an order history record (same as Order).
type OrderHistoryItem = Order

// PositionHistoryItem represents a position history record from PositionHistoryRow schema.
type PositionHistoryItem struct {
	ID                  string               `json:"id"`
	Symbol              string               `json:"symbol"`
	NetQuantity         string               `json:"netQuantity"`
	NetExposureQuantity string               `json:"netExposureQuantity"`
	NetExposureNotional string               `json:"netExposureNotional"`
	NetCost             string               `json:"netCost"`
	MarkPrice           string               `json:"markPrice"`
	EntryPrice          string               `json:"entryPrice"`
	CumulativePnlRealized string             `json:"cumulativePnlRealized"`
	UnrealizedPnl       string               `json:"unrealizedPnl"`
	FundingQuantity     string               `json:"fundingQuantity"`
	Interest            string               `json:"interest"`
	Liquidated          string               `json:"liquidated"`
	IMF                 string               `json:"imf"`
	Fees                string               `json:"fees"`
	State               enums.PositionState  `json:"state"`
	ClosedVolume        string               `json:"closedVolume"`
	LiquidationFees     string               `json:"liquidationFees"`
	ClosingPrice        string               `json:"closingPrice,omitempty"`
	AccountLeverage     string               `json:"accountLeverage,omitempty"`
	OpenedAt            string               `json:"openedAt,omitempty"`
	ClosedAt            string               `json:"closedAt,omitempty"`
}

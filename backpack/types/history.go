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
	Quantity     string                `json:"quantity"`
	Source       enums.SettlementSource `json:"source"`
	SubaccountID *int32                `json:"subaccountId,omitempty"`
	Timestamp    string                `json:"timestamp"`
	UserID       int32                 `json:"userId"`
	PositionID   string                `json:"positionId,omitempty"`
}

// BorrowLendMovement represents a borrow/lend movement record.
type BorrowLendMovement struct {
	EventType        enums.BorrowLendEventType `json:"eventType"`
	PositionID       string                    `json:"positionId"`
	PositionQuantity string                    `json:"positionQuantity,omitempty"`
	Quantity         string                    `json:"quantity"`
	Source           enums.BorrowLendSource    `json:"source"`
	Symbol           string                    `json:"symbol"`
	Timestamp        string                    `json:"timestamp"`
	SpotMarginOrderID string                   `json:"spotMarginOrderId,omitempty"`
}

// BorrowLendPositionHistoryRow represents borrow/lend position history.
type BorrowLendPositionHistoryRow struct {
	PositionID         string                 `json:"positionId"`
	Quantity           string                 `json:"quantity"`
	Symbol             string                 `json:"symbol"`
	Source             enums.BorrowLendSource `json:"source"`
	CumulativeInterest string                 `json:"cumulativeInterest"`
	AvgInterestRate    string                 `json:"avgInterestRate"`
	Side               enums.BorrowLendSide   `json:"side"`
	CreatedAt          string                 `json:"createdAt"`
}

// InterestHistoryItem represents an interest payment record.
type InterestHistoryItem struct {
	PaymentType   enums.PaymentType `json:"paymentType"`
	InterestRate  string            `json:"interestRate"`
	Interval      uint64            `json:"interval"`
	MarketSymbol  string            `json:"marketSymbol"`
	PositionID    string            `json:"positionId"`
	Quantity      string            `json:"quantity"`
	Symbol        string            `json:"symbol"`
	Timestamp     string            `json:"timestamp"`
}

// DustHistoryItem represents a dust conversion record (same as DustConversion).
type DustHistoryItem = DustConversion

// BorrowLendHistoryParams represents parameters for borrow/lend history queries.
type BorrowLendHistoryParams struct {
	Type          enums.BorrowLendEventType `json:"type,omitempty"`
	Sources       []enums.BorrowLendSource  `json:"sources,omitempty"`
	PositionID    string                    `json:"positionId,omitempty"`
	Symbol        string                    `json:"symbol,omitempty"`
	Limit         int                       `json:"limit,omitempty"`
	Offset        int                       `json:"offset,omitempty"`
	SortDirection enums.SortDirection       `json:"sortDirection,omitempty"`
}

// BorrowLendPositionHistoryParams represents parameters for borrow/lend position history.
type BorrowLendPositionHistoryParams struct {
	Symbol        string                    `json:"symbol,omitempty"`
	Side          enums.BorrowLendSide      `json:"side,omitempty"`
	State         enums.BorrowLendPositionState `json:"state,omitempty"`
	Limit         int                       `json:"limit,omitempty"`
	Offset        int                       `json:"offset,omitempty"`
	SortDirection enums.SortDirection       `json:"sortDirection,omitempty"`
}

// InterestHistoryParams represents parameters for interest history queries.
type InterestHistoryParams struct {
	Asset         string                    `json:"asset,omitempty"`
	Symbol        string                    `json:"symbol,omitempty"`
	PositionID    string                    `json:"positionId,omitempty"`
	Limit         int                       `json:"limit,omitempty"`
	Offset        int                       `json:"offset,omitempty"`
	Source        enums.InterestPaymentSource `json:"source,omitempty"`
	SortDirection enums.SortDirection       `json:"sortDirection,omitempty"`
}

// DustHistoryParams represents parameters for dust history queries.
type DustHistoryParams struct {
	ID            *int64                   `json:"id,omitempty"`
	Symbol        string                   `json:"symbol,omitempty"`
	Limit         int                      `json:"limit,omitempty"`
	Offset        int                      `json:"offset,omitempty"`
	SortDirection enums.SortDirection      `json:"sortDirection,omitempty"`
}

// FillHistoryParams represents parameters for fill history queries.
type FillHistoryParams struct {
	OrderID       string              `json:"orderId,omitempty"`
	StrategyID    string              `json:"strategyId,omitempty"`
	From          int64               `json:"from,omitempty"`
	To            int64               `json:"to,omitempty"`
	Symbol        string              `json:"symbol,omitempty"`
	Limit         int                 `json:"limit,omitempty"`
	Offset        int                 `json:"offset,omitempty"`
	FillType      enums.FillType      `json:"fillType,omitempty"`
	MarketType    []enums.MarketType  `json:"marketType,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// FundingHistoryParams represents parameters for funding history queries.
type FundingHistoryParams struct {
	SubaccountID  *int32             `json:"subaccountId,omitempty"`
	Symbol        string             `json:"symbol,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// OrderHistoryParams represents parameters for order history queries.
type OrderHistoryParams struct {
	OrderID       string             `json:"orderId,omitempty"`
	StrategyID    string             `json:"strategyId,omitempty"`
	Symbol        string             `json:"symbol,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	MarketType    []enums.MarketType `json:"marketType,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// PositionHistoryParams represents parameters for position history queries.
type PositionHistoryParams struct {
	Symbol        string             `json:"symbol,omitempty"`
	State         enums.PositionState `json:"state,omitempty"`
	MarketType    []enums.MarketType `json:"marketType,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// QuoteHistoryParams represents parameters for quote history queries.
type QuoteHistoryParams struct {
	QuoteID       string             `json:"quoteId,omitempty"`
	Symbol        string             `json:"symbol,omitempty"`
	Status        enums.OrderStatus  `json:"status,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// QuoteFillHistoryParams represents parameters for quote fill history queries.
type QuoteFillHistoryParams struct {
	QuoteID       string             `json:"quoteId,omitempty"`
	Symbol        string             `json:"symbol,omitempty"`
	Side          enums.Side         `json:"side,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// RFQHistoryParams represents parameters for RFQ history queries.
type RFQHistoryParams struct {
	RFQID         string             `json:"rfqId,omitempty"`
	Symbol        string             `json:"symbol,omitempty"`
	Status        enums.OrderStatus  `json:"status,omitempty"`
	Side          enums.Side         `json:"side,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// RFQFillHistoryParams represents parameters for RFQ fill history queries.
type RFQFillHistoryParams struct {
	QuoteID       string             `json:"quoteId,omitempty"`
	Symbol        string             `json:"symbol,omitempty"`
	Side          enums.Side         `json:"side,omitempty"`
	FillType      enums.RfqFillType  `json:"fillType,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
}

// SettlementHistoryParams represents parameters for settlement history queries.
type SettlementHistoryParams struct {
	Limit         int                         `json:"limit,omitempty"`
	Offset        int                         `json:"offset,omitempty"`
	Source        enums.SettlementSourceFilter `json:"source,omitempty"`
	SortDirection enums.SortDirection         `json:"sortDirection,omitempty"`
}

// StrategyHistoryParams represents parameters for strategy history queries.
type StrategyHistoryParams struct {
	StrategyID    string             `json:"strategyId,omitempty"`
	Symbol        string             `json:"symbol,omitempty"`
	Limit         int                `json:"limit,omitempty"`
	Offset        int                `json:"offset,omitempty"`
	MarketType    []enums.MarketType `json:"marketType,omitempty"`
	SortDirection enums.SortDirection `json:"sortDirection,omitempty"`
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

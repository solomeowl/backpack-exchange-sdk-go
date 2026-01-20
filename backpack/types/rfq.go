package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// RFQ represents a Request for Quote (live, from GET /api/v1/rfq).
type RFQ struct {
	RfqID          string               `json:"rfqId"`
	Symbol         string               `json:"symbol"`
	Side           enums.Side           `json:"side"`
	SubmissionTime string               `json:"submissionTime"`
	ExpiryTime     string               `json:"expiryTime"`
	Status         enums.RFQStatus      `json:"status"`
	ExecutionMode  enums.RfqExecutionMode `json:"executionMode"`
	CreatedAt      string               `json:"createdAt"`
	ClientID       *uint32              `json:"clientId,omitempty"`
	Price          string               `json:"price,omitempty"`
	Quantity       string               `json:"quantity,omitempty"`
	QuoteQuantity  string               `json:"quoteQuantity,omitempty"`
}

// RFQSubmitParams represents parameters for submitting an RFQ (POST /api/v1/rfq).
type RFQSubmitParams struct {
	Symbol         string                `json:"symbol"`
	Side           enums.Side            `json:"side"`
	ClientID       *uint32               `json:"clientId,omitempty"`
	Quantity       string                `json:"quantity,omitempty"`
	QuoteQuantity  string                `json:"quoteQuantity,omitempty"`
	Price          string                `json:"price,omitempty"`
	ExecutionMode  enums.RfqExecutionMode `json:"executionMode,omitempty"`
	AutoLend       *bool                 `json:"autoLend,omitempty"`
	AutoLendRedeem *bool                 `json:"autoLendRedeem,omitempty"`
	AutoBorrow     *bool                 `json:"autoBorrow,omitempty"`
	AutoBorrowRepay *bool                `json:"autoBorrowRepay,omitempty"`
}

// Quote represents a quote response (live, from GET /api/v1/rfq/quote).
type Quote struct {
	QuoteID   string            `json:"quoteId"`
	RfqID     string            `json:"rfqId"`
	BidPrice  string            `json:"bidPrice"`
	AskPrice  string            `json:"askPrice"`
	Status    enums.QuoteStatus `json:"status"`
	CreatedAt string            `json:"createdAt"`
	ClientID  *uint32           `json:"clientId,omitempty"`
}

// QuoteSubmitParams represents parameters for submitting a quote (POST /api/v1/rfq/quote).
type QuoteSubmitParams struct {
	RfqID    string  `json:"rfqId"`
	BidPrice string  `json:"bidPrice"`
	AskPrice string  `json:"askPrice"`
	ClientID *uint32 `json:"clientId,omitempty"`
}

// RFQHistoryItem represents an RFQ history record (RequestForQuoteHistorical).
type RFQHistoryItem struct {
	RfqID          string               `json:"rfqId"`
	Symbol         string               `json:"symbol"`
	Side           enums.Side           `json:"side"`
	SubmissionTime string               `json:"submissionTime"`
	ExpiryTime     string               `json:"expiryTime"`
	Status         enums.RFQStatus      `json:"status"`
	ExecutionMode  enums.RfqExecutionMode `json:"executionMode"`
	CreatedAt      string               `json:"createdAt"`
	ClientID       *uint32              `json:"clientId,omitempty"`
	Price          string               `json:"price,omitempty"`
	Quantity       string               `json:"quantity,omitempty"`
	QuoteQuantity  string               `json:"quoteQuantity,omitempty"`
}

// QuoteHistoryItem represents a quote history record (QuoteHistorical).
type QuoteHistoryItem struct {
	RfqID     string            `json:"rfqId"`
	QuoteID   string            `json:"quoteId"`
	BidPrice  string            `json:"bidPrice"`
	AskPrice  string            `json:"askPrice"`
	Status    enums.QuoteStatus `json:"status"`
	CreatedAt string            `json:"createdAt"`
	ClientID  *uint32           `json:"clientId,omitempty"`
}

// RFQFillHistoryItem represents an RFQ fill history record (RequestForQuoteFillHistorical).
type RFQFillHistoryItem struct {
	RfqID           string               `json:"rfqId"`
	QuoteID         string               `json:"quoteId"`
	Symbol          string               `json:"symbol"`
	Side            enums.Side           `json:"side"`
	FillPrice       string               `json:"fillPrice"`
	CreatedAt       string               `json:"createdAt"`
	FilledAt        string               `json:"filledAt"`
	ClientID        *uint32              `json:"clientId,omitempty"`
	Quantity        string               `json:"quantity,omitempty"`
	QuoteQuantity   string               `json:"quoteQuantity,omitempty"`
	SystemOrderType enums.SystemOrderType `json:"systemOrderType,omitempty"`
}

// QuoteFillHistoryItem represents a quote fill history record (QuoteFillHistorical).
type QuoteFillHistoryItem struct {
	QuoteID   string            `json:"quoteId"`
	RfqID     string            `json:"rfqId"`
	Symbol    string            `json:"symbol"`
	Side      enums.Side        `json:"side"`
	Quantity  string            `json:"quantity"`
	FillPrice string            `json:"fillPrice"`
	Fee       string            `json:"fee"`
	FeeSymbol string            `json:"feeSymbol"`
	CreatedAt string            `json:"createdAt"`
	FilledAt  string            `json:"filledAt"`
	ClientID  *uint32           `json:"clientId,omitempty"`
}

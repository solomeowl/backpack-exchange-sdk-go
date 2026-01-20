package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// RFQ represents a Request for Quote.
type RFQ struct {
	ID          string          `json:"id"`
	Symbol      string          `json:"symbol"`
	Side        enums.Side      `json:"side"`
	Quantity    string          `json:"quantity"`
	Status      enums.RFQStatus `json:"status"`
	ExpiresAt   int64           `json:"expiresAt,omitempty"`
	CreatedAt   int64           `json:"createdAt,omitempty"`
}

// RFQSubmitParams represents parameters for submitting an RFQ.
type RFQSubmitParams struct {
	Symbol   string     `json:"symbol"`
	Side     enums.Side `json:"side"`
	Quantity string     `json:"quantity"`
}

// Quote represents a quote response.
type Quote struct {
	ID        string            `json:"id"`
	RFQID     string            `json:"rfqId"`
	Price     string            `json:"price"`
	Quantity  string            `json:"quantity"`
	Side      enums.Side        `json:"side"`
	Status    enums.QuoteStatus `json:"status"`
	ExpiresAt int64             `json:"expiresAt,omitempty"`
	CreatedAt int64             `json:"createdAt,omitempty"`
}

// QuoteSubmitParams represents parameters for submitting a quote.
type QuoteSubmitParams struct {
	RFQID    string `json:"rfqId"`
	Price    string `json:"price"`
	Quantity string `json:"quantity,omitempty"`
}

// RFQHistoryItem represents an RFQ history record.
type RFQHistoryItem struct {
	RFQ
	Quotes []Quote `json:"quotes,omitempty"`
}

// QuoteHistoryItem represents a quote history record.
type QuoteHistoryItem struct {
	Quote
	RFQ *RFQ `json:"rfq,omitempty"`
}

// RFQFillHistoryItem represents an RFQ fill history record.
type RFQFillHistoryItem struct {
	RFQID      string     `json:"rfqId"`
	QuoteID    string     `json:"quoteId"`
	Symbol     string     `json:"symbol"`
	Side       enums.Side `json:"side"`
	Price      string     `json:"price"`
	Quantity   string     `json:"quantity"`
	Fee        string     `json:"fee,omitempty"`
	FeeSymbol  string     `json:"feeSymbol,omitempty"`
	Timestamp  int64      `json:"timestamp"`
}

// QuoteFillHistoryItem represents a quote fill history record.
type QuoteFillHistoryItem struct {
	QuoteID   string     `json:"quoteId"`
	RFQID     string     `json:"rfqId"`
	Symbol    string     `json:"symbol"`
	Side      enums.Side `json:"side"`
	Price     string     `json:"price"`
	Quantity  string     `json:"quantity"`
	Fee       string     `json:"fee,omitempty"`
	FeeSymbol string     `json:"feeSymbol,omitempty"`
	Timestamp int64      `json:"timestamp"`
}

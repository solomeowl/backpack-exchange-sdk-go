package types

// Trade represents a public trade.
type Trade struct {
	ID            int64  `json:"id,omitempty"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	QuoteQuantity string `json:"quoteQuantity"`
	Timestamp     int64  `json:"timestamp"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
}

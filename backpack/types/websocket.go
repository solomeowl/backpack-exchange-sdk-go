package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// WSMessage represents a generic WebSocket message.
type WSMessage struct {
	Stream string `json:"stream"`
	Data   any    `json:"data"`
}

// WSBookTicker represents a book ticker update.
type WSBookTicker struct {
	Symbol    string `json:"s"`
	BidPrice  string `json:"b"`
	BidQty    string `json:"B"`
	AskPrice  string `json:"a"`
	AskQty    string `json:"A"`
	Timestamp int64  `json:"E"`
}

// WSDepth represents an order book depth update.
type WSDepth struct {
	Symbol       string     `json:"s"`
	Bids         [][]string `json:"b"`
	Asks         [][]string `json:"a"`
	LastUpdateID int64      `json:"u"`
	Timestamp    int64      `json:"E"`
}

// WSKline represents a kline/candlestick update.
type WSKline struct {
	Symbol    string `json:"s"`
	Interval  string `json:"i"`
	OpenTime  int64  `json:"t"`
	CloseTime int64  `json:"T"`
	Open      string `json:"o"`
	High      string `json:"h"`
	Low       string `json:"l"`
	Close     string `json:"c"`
	Volume    string `json:"v"`
	IsClosed  bool   `json:"x"`
}

// WSTicker represents a ticker update.
type WSTicker struct {
	Symbol             string `json:"s"`
	LastPrice          string `json:"c"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	High               string `json:"h"`
	Low                string `json:"l"`
	Volume             string `json:"v"`
	QuoteVolume        string `json:"q"`
	Timestamp          int64  `json:"E"`
}

// WSTrade represents a trade update.
type WSTrade struct {
	Symbol       string `json:"s"`
	TradeID      int64  `json:"t"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	BuyerMaker   bool   `json:"m"`
	Timestamp    int64  `json:"E"`
}

// WSMarkPrice represents a mark price update.
type WSMarkPrice struct {
	Symbol     string `json:"s"`
	MarkPrice  string `json:"p"`
	IndexPrice string `json:"i"`
	Timestamp  int64  `json:"E"`
}

// WSOpenInterest represents an open interest update.
type WSOpenInterest struct {
	Symbol       string `json:"s"`
	OpenInterest string `json:"o"`
	Timestamp    int64  `json:"E"`
}

// WSLiquidation represents a liquidation event.
type WSLiquidation struct {
	Symbol    string     `json:"s"`
	Side      enums.Side `json:"S"`
	Price     string     `json:"p"`
	Quantity  string     `json:"q"`
	Timestamp int64      `json:"E"`
}

// WSOrderUpdate represents an order update (private stream).
type WSOrderUpdate struct {
	Symbol            string            `json:"s"`
	OrderID           string            `json:"i"`
	ClientID          int64             `json:"c,omitempty"`
	Side              enums.Side        `json:"S"`
	OrderType         enums.OrderType   `json:"o"`
	Status            enums.OrderStatus `json:"X"`
	Price             string            `json:"p,omitempty"`
	Quantity          string            `json:"q"`
	ExecutedQuantity  string            `json:"z,omitempty"`
	TimeInForce       string            `json:"f,omitempty"`
	Timestamp         int64             `json:"E"`
}

// WSPositionUpdate represents a position update (private stream).
type WSPositionUpdate struct {
	Symbol           string `json:"s"`
	Quantity         string `json:"pa"`
	EntryPrice       string `json:"ep"`
	MarkPrice        string `json:"mp,omitempty"`
	UnrealizedPnl    string `json:"up,omitempty"`
	LiquidationPrice string `json:"lp,omitempty"`
	Timestamp        int64  `json:"E"`
}

// WSRFQUpdate represents an RFQ update (private stream).
type WSRFQUpdate struct {
	RFQID     string          `json:"r"`
	Symbol    string          `json:"s"`
	Side      enums.Side      `json:"S"`
	Quantity  string          `json:"q"`
	Status    enums.RFQStatus `json:"X"`
	Timestamp int64           `json:"E"`
}

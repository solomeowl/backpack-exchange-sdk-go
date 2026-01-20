package types

import (
	"encoding/json"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
)

// WSMessage represents a generic WebSocket message.
type WSMessage struct {
	Stream string `json:"stream"`
	Data   any    `json:"data"`
}

// WSID decodes IDs that may be returned as string or number.
type WSID string

// UnmarshalJSON accepts either a JSON string or number.
func (id *WSID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*id = WSID(s)
		return nil
	}
	var num json.Number
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}
	*id = WSID(num.String())
	return nil
}

// WSBookTicker represents a book ticker update.
type WSBookTicker struct {
	Event           string `json:"e"`
	EventTime       int64  `json:"E"`
	Symbol          string `json:"s"`
	AskPrice        string `json:"a"`
	AskQty          string `json:"A"`
	BidPrice        string `json:"b"`
	BidQty          string `json:"B"`
	UpdateID        WSID   `json:"u"`
	EngineTimestamp int64  `json:"T"`
}

// WSDepth represents an order book depth update.
type WSDepth struct {
	Event           string     `json:"e"`
	EventTime       int64      `json:"E"`
	Symbol          string     `json:"s"`
	Asks            [][]string `json:"a"`
	Bids            [][]string `json:"b"`
	FirstUpdateID   int64      `json:"U"`
	LastUpdateID    int64      `json:"u"`
	EngineTimestamp int64      `json:"T"`
}

// WSKline represents a kline/candlestick update.
type WSKline struct {
	Event     string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	StartTime string `json:"t"`
	CloseTime string `json:"T"`
	Open      string `json:"o"`
	Close     string `json:"c"`
	High      string `json:"h"`
	Low       string `json:"l"`
	Volume    string `json:"v"`
	Trades    int64  `json:"n"`
	IsClosed  bool   `json:"X"`
}

// WSTicker represents a ticker update.
type WSTicker struct {
	Event       string `json:"e"`
	EventTime   int64  `json:"E"`
	Symbol      string `json:"s"`
	Open        string `json:"o"`
	Close       string `json:"c"`
	High        string `json:"h"`
	Low         string `json:"l"`
	Volume      string `json:"v"`
	QuoteVolume string `json:"V"`
	Trades      int64  `json:"n"`
}

// WSTrade represents a trade update.
type WSTrade struct {
	Event           string `json:"e"`
	EventTime       int64  `json:"E"`
	Symbol          string `json:"s"`
	Price           string `json:"p"`
	Quantity        string `json:"q"`
	BuyerOrderID    WSID   `json:"b"`
	SellerOrderID   WSID   `json:"a"`
	TradeID         int64  `json:"t"`
	EngineTimestamp int64  `json:"T"`
	IsBuyerMaker    bool   `json:"m"`
}

// WSMarkPrice represents a mark price update.
type WSMarkPrice struct {
	Event              string `json:"e"`
	EventTime          int64  `json:"E"`
	Symbol             string `json:"s"`
	MarkPrice          string `json:"p"`
	FundingRate        string `json:"f,omitempty"`
	IndexPrice         string `json:"i,omitempty"`
	NextFundingTimestamp int64 `json:"n,omitempty"`
	EngineTimestamp    int64  `json:"T"`
}

// WSOpenInterest represents an open interest update.
type WSOpenInterest struct {
	Event       string `json:"e"`
	EventTime   int64  `json:"E"`
	Symbol      string `json:"s"`
	OpenInterest string `json:"o"`
}

// WSLiquidation represents a liquidation event.
type WSLiquidation struct {
	Event           string     `json:"e"`
	EventTime       int64      `json:"E"`
	Quantity        string     `json:"q"`
	Price           string     `json:"p"`
	Side            enums.Side `json:"S"`
	Symbol          string     `json:"s"`
	EngineTimestamp int64      `json:"T"`
}

// WSOrderUpdate represents an order update (private stream).
type WSOrderUpdate struct {
	Event               string                 `json:"e,omitempty"`
	EventTime           int64                  `json:"E,omitempty"`
	Symbol              string                 `json:"s,omitempty"`
	ClientID            *uint32                `json:"c,omitempty"`
	Side                enums.Side             `json:"S,omitempty"`
	OrderType           string                 `json:"o,omitempty"`
	TimeInForce         string                 `json:"f,omitempty"`
	Quantity            string                 `json:"q,omitempty"`
	QuoteQuantity       string                 `json:"Q,omitempty"`
	Price               string                 `json:"p,omitempty"`
	TriggerPrice        string                 `json:"P,omitempty"`
	TriggerBy           string                 `json:"B,omitempty"`
	TakeProfitTriggerPrice string              `json:"a,omitempty"`
	StopLossTriggerPrice   string              `json:"b,omitempty"`
	TakeProfitLimitPrice   string              `json:"j,omitempty"`
	StopLossLimitPrice     string              `json:"k,omitempty"`
	TakeProfitTriggerBy    string              `json:"d,omitempty"`
	StopLossTriggerBy      string              `json:"g,omitempty"`
	TriggerQuantity        string              `json:"Y,omitempty"`
	Status               enums.OrderStatus     `json:"X,omitempty"`
	ExpiryReason         enums.OrderExpiryReason `json:"R,omitempty"`
	OrderID              WSID                  `json:"i,omitempty"`
	TradeID              *int64                `json:"t,omitempty"`
	FillQuantity         string                `json:"l,omitempty"`
	ExecutedQuantity     string                `json:"z,omitempty"`
	ExecutedQuoteQuantity string               `json:"Z,omitempty"`
	FillPrice            string                `json:"L,omitempty"`
	IsMaker              *bool                 `json:"m,omitempty"`
	Fee                  string                `json:"n,omitempty"`
	FeeSymbol            string                `json:"N,omitempty"`
	SelfTradePrevention  enums.SelfTradePrevention `json:"V,omitempty"`
	EngineTimestamp      int64                 `json:"T,omitempty"`
	Origin               string                `json:"O,omitempty"`
	RelatedOrderID       WSID                  `json:"I,omitempty"`
	StrategyID           WSID                  `json:"H,omitempty"`
	PostOnly             *bool                 `json:"y,omitempty"`
	ReduceOnly           *bool                 `json:"r,omitempty"`
}

// WSPositionUpdate represents a position update (private stream).
type WSPositionUpdate struct {
	Event               string `json:"e,omitempty"`
	EventTime           int64  `json:"E,omitempty"`
	Symbol              string `json:"s,omitempty"`
	BreakEvenPrice      string `json:"b,omitempty"`
	EntryPrice          string `json:"B,omitempty"`
	IMF                 string `json:"f,omitempty"`
	MarkPrice           string `json:"M,omitempty"`
	MMF                 string `json:"m,omitempty"`
	NetQuantity         string `json:"q,omitempty"`
	NetExposureQuantity string `json:"Q,omitempty"`
	NetExposureNotional string `json:"n,omitempty"`
	PositionID          WSID   `json:"i,omitempty"`
	PnlRealized         string `json:"p,omitempty"`
	PnlUnrealized       string `json:"P,omitempty"`
	EngineTimestamp     int64  `json:"T,omitempty"`
}

// WSRFQUpdate represents an RFQ update (private stream).
type WSRFQUpdate struct {
	Event          string            `json:"e,omitempty"`
	EventTime      int64             `json:"E,omitempty"`
	RFQID          WSID              `json:"R,omitempty"`
	QuoteID        WSID              `json:"u,omitempty"`
	ClientID       WSID              `json:"C,omitempty"`
	Symbol         string            `json:"s,omitempty"`
	Side           enums.Side        `json:"S,omitempty"`
	Quantity       string            `json:"q,omitempty"`
	QuoteQuantity  string            `json:"Q,omitempty"`
	Price          string            `json:"p,omitempty"`
	SubmissionTime int64             `json:"w,omitempty"`
	ExpiryTime     int64             `json:"W,omitempty"`
	Status         enums.OrderStatus `json:"X,omitempty"`
	EngineTimestamp int64            `json:"T,omitempty"`
}

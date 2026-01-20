package websocket

import (
	"encoding/json"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// Handler wraps callbacks for type-safe message handling.
type Handler struct {
	client *Client
}

// NewHandler creates a new Handler for the given client.
func NewHandler(client *Client) *Handler {
	return &Handler{client: client}
}

// OnBookTicker subscribes to book ticker updates.
func (h *Handler) OnBookTicker(symbol string, callback func(*types.WSBookTicker)) error {
	return h.client.Subscribe([]string{BookTickerStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSBookTicker
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnDepth subscribes to order book depth updates.
func (h *Handler) OnDepth(symbol string, callback func(*types.WSDepth)) error {
	return h.client.Subscribe([]string{DepthStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSDepth
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnKline subscribes to kline/candlestick updates.
func (h *Handler) OnKline(symbol, interval string, callback func(*types.WSKline)) error {
	return h.client.Subscribe([]string{KlineStream(symbol, interval)}, func(data json.RawMessage) {
		var msg types.WSKline
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnTicker subscribes to ticker updates.
func (h *Handler) OnTicker(symbol string, callback func(*types.WSTicker)) error {
	return h.client.Subscribe([]string{TickerStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSTicker
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnTrade subscribes to trade updates.
func (h *Handler) OnTrade(symbol string, callback func(*types.WSTrade)) error {
	return h.client.Subscribe([]string{TradeStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSTrade
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnMarkPrice subscribes to mark price updates.
func (h *Handler) OnMarkPrice(symbol string, callback func(*types.WSMarkPrice)) error {
	return h.client.Subscribe([]string{MarkPriceStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSMarkPrice
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnOpenInterest subscribes to open interest updates.
func (h *Handler) OnOpenInterest(symbol string, callback func(*types.WSOpenInterest)) error {
	return h.client.Subscribe([]string{OpenInterestStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSOpenInterest
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnLiquidation subscribes to liquidation updates.
func (h *Handler) OnLiquidation(symbol string, callback func(*types.WSLiquidation)) error {
	return h.client.Subscribe([]string{LiquidationStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSLiquidation
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnOrderUpdate subscribes to order updates (private stream).
// If symbol is empty, subscribes to updates for all markets.
func (h *Handler) OnOrderUpdate(symbol string, callback func(*types.WSOrderUpdate)) error {
	return h.client.SubscribePrivate([]string{OrderUpdateStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSOrderUpdate
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnPositionUpdate subscribes to position updates (private stream).
// If symbol is empty, subscribes to updates for all markets.
func (h *Handler) OnPositionUpdate(symbol string, callback func(*types.WSPositionUpdate)) error {
	return h.client.SubscribePrivate([]string{PositionUpdateStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSPositionUpdate
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

// OnRFQUpdate subscribes to RFQ updates (private stream).
// If symbol is empty, subscribes to updates for all markets.
func (h *Handler) OnRFQUpdate(symbol string, callback func(*types.WSRFQUpdate)) error {
	return h.client.SubscribePrivate([]string{RFQUpdateStream(symbol)}, func(data json.RawMessage) {
		var msg types.WSRFQUpdate
		if err := json.Unmarshal(data, &msg); err == nil {
			callback(&msg)
		}
	})
}

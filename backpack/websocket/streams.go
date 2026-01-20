package websocket

import "fmt"

// Public stream names
const (
	StreamBookTicker    = "bookTicker"
	StreamDepth         = "depth"
	StreamDepth200ms    = "depth.200ms"
	StreamDepth600ms    = "depth.600ms"
	StreamDepth1000ms   = "depth.1000ms"
	StreamKline         = "kline"
	StreamTicker        = "ticker"
	StreamTrade         = "trade"
	StreamMarkPrice     = "markPrice"
	StreamOpenInterest  = "openInterest"
	StreamLiquidation   = "liquidation"
)

// Private stream names
const (
	StreamOrderUpdate    = "account.orderUpdate"
	StreamPositionUpdate = "account.positionUpdate"
	StreamRFQUpdate      = "account.rfqUpdate"
)

// BookTickerStream returns the book ticker stream name for a symbol.
func BookTickerStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamBookTicker, symbol)
}

// DepthStream returns the depth stream name for a symbol.
func DepthStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamDepth, symbol)
}

// Depth200msStream returns the 200ms depth stream name for a symbol.
func Depth200msStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamDepth200ms, symbol)
}

// Depth600msStream returns the 600ms depth stream name for a symbol.
func Depth600msStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamDepth600ms, symbol)
}

// Depth1000msStream returns the 1000ms depth stream name for a symbol.
func Depth1000msStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamDepth1000ms, symbol)
}

// KlineStream returns the kline stream name for a symbol and interval.
func KlineStream(symbol, interval string) string {
	return fmt.Sprintf("%s.%s.%s", StreamKline, interval, symbol)
}

// TickerStream returns the ticker stream name for a symbol.
func TickerStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamTicker, symbol)
}

// TradeStream returns the trade stream name for a symbol.
func TradeStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamTrade, symbol)
}

// MarkPriceStream returns the mark price stream name for a symbol.
func MarkPriceStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamMarkPrice, symbol)
}

// OpenInterestStream returns the open interest stream name for a symbol.
func OpenInterestStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamOpenInterest, symbol)
}

// LiquidationStream returns the liquidation stream name for a symbol.
func LiquidationStream(symbol string) string {
	return fmt.Sprintf("%s.%s", StreamLiquidation, symbol)
}

// OrderUpdateStream returns the order update stream name for a symbol.
// If symbol is empty, the stream includes all markets.
func OrderUpdateStream(symbol string) string {
	if symbol == "" {
		return StreamOrderUpdate
	}
	return fmt.Sprintf("%s.%s", StreamOrderUpdate, symbol)
}

// PositionUpdateStream returns the position update stream name for a symbol.
// If symbol is empty, the stream includes all markets.
func PositionUpdateStream(symbol string) string {
	if symbol == "" {
		return StreamPositionUpdate
	}
	return fmt.Sprintf("%s.%s", StreamPositionUpdate, symbol)
}

// RFQUpdateStream returns the RFQ update stream name for a symbol.
// If symbol is empty, the stream includes all markets.
func RFQUpdateStream(symbol string) string {
	if symbol == "" {
		return StreamRFQUpdate
	}
	return fmt.Sprintf("%s.%s", StreamRFQUpdate, symbol)
}

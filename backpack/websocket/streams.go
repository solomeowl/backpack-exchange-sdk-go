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
	return fmt.Sprintf("%s@%s", symbol, StreamBookTicker)
}

// DepthStream returns the depth stream name for a symbol.
func DepthStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamDepth)
}

// Depth200msStream returns the 200ms depth stream name for a symbol.
func Depth200msStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamDepth200ms)
}

// Depth600msStream returns the 600ms depth stream name for a symbol.
func Depth600msStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamDepth600ms)
}

// Depth1000msStream returns the 1000ms depth stream name for a symbol.
func Depth1000msStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamDepth1000ms)
}

// KlineStream returns the kline stream name for a symbol and interval.
func KlineStream(symbol, interval string) string {
	return fmt.Sprintf("%s@%s_%s", symbol, StreamKline, interval)
}

// TickerStream returns the ticker stream name for a symbol.
func TickerStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamTicker)
}

// TradeStream returns the trade stream name for a symbol.
func TradeStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamTrade)
}

// MarkPriceStream returns the mark price stream name for a symbol.
func MarkPriceStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamMarkPrice)
}

// OpenInterestStream returns the open interest stream name for a symbol.
func OpenInterestStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamOpenInterest)
}

// LiquidationStream returns the liquidation stream name for a symbol.
func LiquidationStream(symbol string) string {
	return fmt.Sprintf("%s@%s", symbol, StreamLiquidation)
}

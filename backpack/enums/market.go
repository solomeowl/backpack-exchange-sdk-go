package enums

// MarketType represents the type of market.
type MarketType string

const (
	MarketTypeSpot MarketType = "SPOT"
	MarketTypePerp MarketType = "PERP"
)

// KlineInterval represents kline/candlestick intervals.
type KlineInterval string

const (
	KlineInterval1m  KlineInterval = "1m"
	KlineInterval3m  KlineInterval = "3m"
	KlineInterval5m  KlineInterval = "5m"
	KlineInterval15m KlineInterval = "15m"
	KlineInterval30m KlineInterval = "30m"
	KlineInterval1h  KlineInterval = "1h"
	KlineInterval2h  KlineInterval = "2h"
	KlineInterval4h  KlineInterval = "4h"
	KlineInterval6h  KlineInterval = "6h"
	KlineInterval8h  KlineInterval = "8h"
	KlineInterval12h KlineInterval = "12h"
	KlineInterval1d  KlineInterval = "1d"
	KlineInterval3d  KlineInterval = "3d"
	KlineInterval1w  KlineInterval = "1w"
	KlineInterval1M  KlineInterval = "1M"
)

// TickerInterval represents ticker update intervals.
type TickerInterval string

const (
	TickerInterval1h  TickerInterval = "1h"
	TickerInterval4h  TickerInterval = "4h"
	TickerInterval24h TickerInterval = "24h"
)

// DepthLimit represents order book depth limits.
type DepthLimit int

const (
	DepthLimit5   DepthLimit = 5
	DepthLimit10  DepthLimit = 10
	DepthLimit20  DepthLimit = 20
	DepthLimit50  DepthLimit = 50
	DepthLimit100 DepthLimit = 100
	DepthLimit500 DepthLimit = 500
)

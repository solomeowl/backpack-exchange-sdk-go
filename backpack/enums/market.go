package enums

// MarketType represents the type of market.
type MarketType string

const (
	MarketTypeSpot       MarketType = "SPOT"
	MarketTypePerp       MarketType = "PERP"
	MarketTypeIPerp      MarketType = "IPERP"
	MarketTypeDated      MarketType = "DATED"
	MarketTypePrediction MarketType = "PREDICTION"
	MarketTypeRFQ        MarketType = "RFQ"
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
	TickerInterval1d TickerInterval = "1d"
	TickerInterval1w TickerInterval = "1w"
)

// DepthLimit represents order book depth limits (as string values).
type DepthLimit string

const (
	DepthLimit5    DepthLimit = "5"
	DepthLimit10   DepthLimit = "10"
	DepthLimit20   DepthLimit = "20"
	DepthLimit50   DepthLimit = "50"
	DepthLimit100  DepthLimit = "100"
	DepthLimit500  DepthLimit = "500"
	DepthLimit1000 DepthLimit = "1000"
)

// KlinePriceType represents kline price types.
type KlinePriceType string

const (
	KlinePriceTypeLast  KlinePriceType = "Last"
	KlinePriceTypeIndex KlinePriceType = "Index"
	KlinePriceTypeMark  KlinePriceType = "Mark"
)

// OrderBookState represents the state of an order book.
type OrderBookState string

const (
	OrderBookStateOpen       OrderBookState = "Open"
	OrderBookStateClosed     OrderBookState = "Closed"
	OrderBookStateCancelOnly OrderBookState = "CancelOnly"
	OrderBookStateLimitOnly  OrderBookState = "LimitOnly"
	OrderBookStatePostOnly   OrderBookState = "PostOnly"
)

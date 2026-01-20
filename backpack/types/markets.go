package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Market represents a trading market from GET /api/v1/markets.
type Market struct {
	Symbol         string                 `json:"symbol"`
	BaseSymbol     enums.CustodyAsset     `json:"baseSymbol"`
	QuoteSymbol    enums.CustodyAsset     `json:"quoteSymbol"`
	MarketType     enums.MarketType       `json:"marketType"`
	Filters        OrderBookFilters       `json:"filters"`
	ImfFunction    *PositionImfFunction   `json:"imfFunction,omitempty"`
	MmfFunction    *PositionImfFunction   `json:"mmfFunction,omitempty"`
	FundingInterval       uint64          `json:"fundingInterval,omitempty"`
	FundingRateUpperBound string          `json:"fundingRateUpperBound,omitempty"`
	FundingRateLowerBound string          `json:"fundingRateLowerBound,omitempty"`
	OpenInterestLimit     string          `json:"openInterestLimit,omitempty"`
	OrderBookState enums.OrderBookState   `json:"orderBookState"`
	CreatedAt      string                 `json:"createdAt"`
	Visible        bool                   `json:"visible"`
	PositionLimitWeight string            `json:"positionLimitWeight,omitempty"`
}

// OrderBookFilters represents price and quantity rules for a market.
type OrderBookFilters struct {
	Price    PriceFilter    `json:"price"`
	Quantity QuantityFilter `json:"quantity"`
}

// PriceFilter defines the price rules for the order book.
type PriceFilter struct {
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice,omitempty"`
	TickSize string `json:"tickSize"`
}

// QuantityFilter defines the quantity rules for the order book.
type QuantityFilter struct {
	MinQuantity string `json:"minQuantity"`
	MaxQuantity string `json:"maxQuantity,omitempty"`
	StepSize    string `json:"stepSize"`
}

// Ticker represents market ticker data from GET /api/v1/ticker.
type Ticker struct {
	Symbol             string `json:"symbol"`
	FirstPrice         string `json:"firstPrice"`
	LastPrice          string `json:"lastPrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	High               string `json:"high"`
	Low                string `json:"low"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	Trades             string `json:"trades"`
}

// OrderBook represents the order book depth from GET /api/v1/depth.
type OrderBook struct {
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`
	LastUpdateID string     `json:"lastUpdateId"`
	Timestamp    int64      `json:"timestamp"`
}

// Kline represents a candlestick/kline from GET /api/v1/klines.
type Kline struct {
	Start       string `json:"start"`
	End         string `json:"end"`
	Open        string `json:"open,omitempty"`
	High        string `json:"high,omitempty"`
	Low         string `json:"low,omitempty"`
	Close       string `json:"close,omitempty"`
	Volume      string `json:"volume"`
	QuoteVolume string `json:"quoteVolume"`
	Trades      string `json:"trades"`
}

// MarkPrice represents mark price data from GET /api/v1/markPrices.
type MarkPrice struct {
	Symbol               string `json:"symbol"`
	MarkPrice            string `json:"markPrice"`
	FundingRate          string `json:"fundingRate,omitempty"`
	IndexPrice           string `json:"indexPrice,omitempty"`
	NextFundingTimestamp int64  `json:"nextFundingTimestamp,omitempty"`
}

// OpenInterest represents open interest from GET /api/v1/openInterest.
type OpenInterest struct {
	Symbol       string `json:"symbol"`
	OpenInterest string `json:"openInterest,omitempty"`
	Timestamp    int64  `json:"timestamp"`
}

// FundingRate represents funding interval rate from GET /api/v1/fundingRates.
type FundingRate struct {
	Symbol               string `json:"symbol"`
	IntervalEndTimestamp string `json:"intervalEndTimestamp"`
	FundingRate          string `json:"fundingRate"`
}

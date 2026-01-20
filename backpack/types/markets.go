package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Market represents a trading market.
type Market struct {
	Symbol              string           `json:"symbol"`
	BaseSymbol          string           `json:"baseSymbol"`
	QuoteSymbol         string           `json:"quoteSymbol"`
	MarketType          enums.MarketType `json:"marketType"`
	Status              string           `json:"status"`
	MinOrderSize        string           `json:"minOrderSize,omitempty"`
	MaxOrderSize        string           `json:"maxOrderSize,omitempty"`
	MinNotional         string           `json:"minNotional,omitempty"`
	TickSize            string           `json:"tickSize,omitempty"`
	StepSize            string           `json:"stepSize,omitempty"`
	MaxLeverage         string           `json:"maxLeverage,omitempty"`
	MaintenanceMargin   string           `json:"maintenanceMargin,omitempty"`
	InitialMargin       string           `json:"initialMargin,omitempty"`
	MakerFee            string           `json:"makerFee,omitempty"`
	TakerFee            string           `json:"takerFee,omitempty"`
}

// Ticker represents market ticker data.
type Ticker struct {
	Symbol             string `json:"symbol"`
	FirstPrice         string `json:"firstPrice,omitempty"`
	LastPrice          string `json:"lastPrice"`
	PriceChange        string `json:"priceChange,omitempty"`
	PriceChangePercent string `json:"priceChangePercent,omitempty"`
	High               string `json:"high,omitempty"`
	Low                string `json:"low,omitempty"`
	Volume             string `json:"volume,omitempty"`
	QuoteVolume        string `json:"quoteVolume,omitempty"`
	Trades             int64  `json:"trades,omitempty"`
}

// OrderBookLevel represents a price level in the order book.
type OrderBookLevel struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

// OrderBook represents the order book depth.
type OrderBook struct {
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
	LastUpdateID int64      `json:"lastUpdateId,omitempty"`
}

// Kline represents a candlestick/kline.
type Kline struct {
	OpenTime  int64  `json:"openTime"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	CloseTime int64  `json:"closeTime"`
}

// MarkPrice represents the mark price for a perpetual market.
type MarkPrice struct {
	Symbol        string `json:"symbol"`
	MarkPrice     string `json:"markPrice"`
	IndexPrice    string `json:"indexPrice,omitempty"`
	Timestamp     int64  `json:"timestamp,omitempty"`
}

// OpenInterest represents open interest for a perpetual market.
type OpenInterest struct {
	Symbol       string `json:"symbol"`
	OpenInterest string `json:"openInterest"`
	Timestamp    int64  `json:"timestamp,omitempty"`
}

// FundingRate represents funding rate information.
type FundingRate struct {
	Symbol              string `json:"symbol"`
	FundingRate         string `json:"fundingRate"`
	FundingRateTimestamp int64  `json:"fundingRateTimestamp,omitempty"`
	NextFundingTime     int64  `json:"nextFundingTime,omitempty"`
}

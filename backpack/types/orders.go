package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Order represents an order.
type Order struct {
	ID                    string             `json:"id"`
	ClientID              int64              `json:"clientId,omitempty"`
	Symbol                string             `json:"symbol"`
	Side                  enums.Side         `json:"side"`
	OrderType             enums.OrderType    `json:"orderType"`
	Status                enums.OrderStatus  `json:"status"`
	Price                 string             `json:"price,omitempty"`
	Quantity              string             `json:"quantity,omitempty"`
	QuoteQuantity         string             `json:"quoteQuantity,omitempty"`
	ExecutedQuantity      string             `json:"executedQuantity,omitempty"`
	ExecutedQuoteQuantity string             `json:"executedQuoteQuantity,omitempty"`
	TimeInForce           enums.TimeInForce  `json:"timeInForce,omitempty"`
	PostOnly              bool               `json:"postOnly,omitempty"`
	ReduceOnly            bool               `json:"reduceOnly,omitempty"`
	TriggerPrice          string             `json:"triggerPrice,omitempty"`
	TriggerQuantity       string             `json:"triggerQuantity,omitempty"`
	SelfTradePrevention   string             `json:"selfTradePrevention,omitempty"`
	CreatedAt             string             `json:"createdAt,omitempty"`
	UpdatedAt             string             `json:"updatedAt,omitempty"`
}

// ExecuteOrderParams represents parameters for executing an order.
type ExecuteOrderParams struct {
	Symbol                string                      `json:"symbol"`
	Side                  enums.Side                  `json:"side"`
	OrderType             enums.OrderType             `json:"orderType"`
	Price                 string                      `json:"price,omitempty"`
	Quantity              string                      `json:"quantity,omitempty"`
	QuoteQuantity         string                      `json:"quoteQuantity,omitempty"`
	TimeInForce           enums.TimeInForce           `json:"timeInForce,omitempty"`
	PostOnly              *bool                       `json:"postOnly,omitempty"`
	ClientID              *int64                      `json:"clientId,omitempty"`
	SelfTradePrevention   enums.SelfTradePrevention   `json:"selfTradePrevention,omitempty"`
	TriggerPrice          string                      `json:"triggerPrice,omitempty"`
	TriggerQuantity       string                      `json:"triggerQuantity,omitempty"`
	ReduceOnly            *bool                       `json:"reduceOnly,omitempty"`
	AutoBorrow            *bool                       `json:"autoBorrow,omitempty"`
	AutoBorrowRepay       *bool                       `json:"autoBorrowRepay,omitempty"`
	AutoLend              *bool                       `json:"autoLend,omitempty"`
	AutoLendRedeem        *bool                       `json:"autoLendRedeem,omitempty"`
	StopLossTriggerPrice  string                      `json:"stopLossTriggerPrice,omitempty"`
	StopLossLimitPrice    string                      `json:"stopLossLimitPrice,omitempty"`
	TakeProfitTriggerPrice string                     `json:"takeProfitTriggerPrice,omitempty"`
	TakeProfitLimitPrice  string                      `json:"takeProfitLimitPrice,omitempty"`
}

// ToMap converts ExecuteOrderParams to a map for signing.
func (p *ExecuteOrderParams) ToMap() map[string]any {
	m := map[string]any{
		"symbol":    p.Symbol,
		"side":      string(p.Side),
		"orderType": string(p.OrderType),
	}

	if p.Price != "" {
		m["price"] = p.Price
	}
	if p.Quantity != "" {
		m["quantity"] = p.Quantity
	}
	if p.QuoteQuantity != "" {
		m["quoteQuantity"] = p.QuoteQuantity
	}
	if p.TimeInForce != "" {
		m["timeInForce"] = string(p.TimeInForce)
	}
	if p.PostOnly != nil {
		m["postOnly"] = *p.PostOnly
	}
	if p.ClientID != nil {
		m["clientId"] = *p.ClientID
	}
	if p.SelfTradePrevention != "" {
		m["selfTradePrevention"] = string(p.SelfTradePrevention)
	}
	if p.TriggerPrice != "" {
		m["triggerPrice"] = p.TriggerPrice
	}
	if p.TriggerQuantity != "" {
		m["triggerQuantity"] = p.TriggerQuantity
	}
	if p.ReduceOnly != nil {
		m["reduceOnly"] = *p.ReduceOnly
	}
	if p.AutoBorrow != nil {
		m["autoBorrow"] = *p.AutoBorrow
	}
	if p.AutoBorrowRepay != nil {
		m["autoBorrowRepay"] = *p.AutoBorrowRepay
	}
	if p.AutoLend != nil {
		m["autoLend"] = *p.AutoLend
	}
	if p.AutoLendRedeem != nil {
		m["autoLendRedeem"] = *p.AutoLendRedeem
	}
	if p.StopLossTriggerPrice != "" {
		m["stopLossTriggerPrice"] = p.StopLossTriggerPrice
	}
	if p.StopLossLimitPrice != "" {
		m["stopLossLimitPrice"] = p.StopLossLimitPrice
	}
	if p.TakeProfitTriggerPrice != "" {
		m["takeProfitTriggerPrice"] = p.TakeProfitTriggerPrice
	}
	if p.TakeProfitLimitPrice != "" {
		m["takeProfitLimitPrice"] = p.TakeProfitLimitPrice
	}

	return m
}

// GetOrderParams represents parameters for getting an order.
type GetOrderParams struct {
	Symbol   string `json:"symbol"`
	OrderID  string `json:"orderId,omitempty"`
	ClientID *int64 `json:"clientId,omitempty"`
}

// CancelOrderParams represents parameters for canceling an order.
type CancelOrderParams struct {
	Symbol   string `json:"symbol"`
	OrderID  string `json:"orderId,omitempty"`
	ClientID *int64 `json:"clientId,omitempty"`
}

// GetOpenOrdersParams represents parameters for getting open orders.
type GetOpenOrdersParams struct {
	Symbol     string           `json:"symbol,omitempty"`
	MarketType enums.MarketType `json:"marketType,omitempty"`
}

// CancelAllOrdersParams represents parameters for canceling all orders.
type CancelAllOrdersParams struct {
	Symbol    string                `json:"symbol"`
	OrderType enums.CancelOrderType `json:"orderType,omitempty"`
}

// BatchOrderResult represents the result of a batch order.
type BatchOrderResult struct {
	Order *Order `json:"order,omitempty"`
	Error string `json:"error,omitempty"`
}

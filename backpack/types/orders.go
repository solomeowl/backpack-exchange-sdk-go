package types

import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Order represents an order from GET /api/v1/order or GET /api/v1/orders.
type Order struct {
	ID                     string                       `json:"id"`
	CreatedAt              TimeString                   `json:"createdAt"`
	OrderType              enums.OrderType              `json:"orderType"`
	SelfTradePrevention    enums.SelfTradePrevention    `json:"selfTradePrevention"`
	Status                 enums.OrderStatus            `json:"status"`
	Side                   enums.Side                   `json:"side"`
	Symbol                 string                       `json:"symbol"`
	TimeInForce            enums.TimeInForce            `json:"timeInForce"`
	ExecutedQuantity       string                       `json:"executedQuantity,omitempty"`
	ExecutedQuoteQuantity  string                       `json:"executedQuoteQuantity,omitempty"`
	ExpiryReason           enums.OrderExpiryReason      `json:"expiryReason,omitempty"`
	PostOnly               bool                         `json:"postOnly,omitempty"`
	Price                  string                       `json:"price,omitempty"`
	Quantity               string                       `json:"quantity,omitempty"`
	QuoteQuantity          string                       `json:"quoteQuantity,omitempty"`
	StopLossTriggerPrice   string                       `json:"stopLossTriggerPrice,omitempty"`
	StopLossLimitPrice     string                       `json:"stopLossLimitPrice,omitempty"`
	StopLossTriggerBy      string                       `json:"stopLossTriggerBy,omitempty"`
	TakeProfitTriggerPrice string                       `json:"takeProfitTriggerPrice,omitempty"`
	TakeProfitLimitPrice   string                       `json:"takeProfitLimitPrice,omitempty"`
	TakeProfitTriggerBy    string                       `json:"takeProfitTriggerBy,omitempty"`
	TriggerBy              string                       `json:"triggerBy,omitempty"`
	TriggerPrice           string                       `json:"triggerPrice,omitempty"`
	TriggerQuantity        string                       `json:"triggerQuantity,omitempty"`
	ClientID               uint32                       `json:"clientId,omitempty"`
	SystemOrderType        enums.SystemOrderType        `json:"systemOrderType,omitempty"`
	StrategyID             string                       `json:"strategyId,omitempty"`
	SlippageTolerance      string                       `json:"slippageTolerance,omitempty"`
	SlippageToleranceType  enums.SlippageToleranceType  `json:"slippageToleranceType,omitempty"`
}

// ExecuteOrderParams represents parameters for executing an order via POST /api/v1/order.
type ExecuteOrderParams struct {
	Symbol                 string                       `json:"symbol"`
	Side                   enums.Side                   `json:"side"`
	OrderType              enums.OrderType              `json:"orderType"`
	Price                  string                       `json:"price,omitempty"`
	Quantity               string                       `json:"quantity,omitempty"`
	QuoteQuantity          string                       `json:"quoteQuantity,omitempty"`
	TimeInForce            enums.TimeInForce            `json:"timeInForce,omitempty"`
	PostOnly               *bool                        `json:"postOnly,omitempty"`
	ClientID               *uint32                      `json:"clientId,omitempty"`
	BrokerID               *uint16                      `json:"brokerId,omitempty"`
	SelfTradePrevention    enums.SelfTradePrevention    `json:"selfTradePrevention,omitempty"`
	TriggerPrice           string                       `json:"triggerPrice,omitempty"`
	TriggerQuantity        string                       `json:"triggerQuantity,omitempty"`
	TriggerBy              string                       `json:"triggerBy,omitempty"`
	ReduceOnly             *bool                        `json:"reduceOnly,omitempty"`
	AutoBorrow             *bool                        `json:"autoBorrow,omitempty"`
	AutoBorrowRepay        *bool                        `json:"autoBorrowRepay,omitempty"`
	AutoLend               *bool                        `json:"autoLend,omitempty"`
	AutoLendRedeem         *bool                        `json:"autoLendRedeem,omitempty"`
	StopLossTriggerPrice   string                       `json:"stopLossTriggerPrice,omitempty"`
	StopLossLimitPrice     string                       `json:"stopLossLimitPrice,omitempty"`
	StopLossTriggerBy      string                       `json:"stopLossTriggerBy,omitempty"`
	TakeProfitTriggerPrice string                       `json:"takeProfitTriggerPrice,omitempty"`
	TakeProfitLimitPrice   string                       `json:"takeProfitLimitPrice,omitempty"`
	TakeProfitTriggerBy    string                       `json:"takeProfitTriggerBy,omitempty"`
	SlippageTolerance      string                       `json:"slippageTolerance,omitempty"`
	SlippageToleranceType  enums.SlippageToleranceType  `json:"slippageToleranceType,omitempty"`
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
	if p.BrokerID != nil {
		m["brokerId"] = *p.BrokerID
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
	if p.TriggerBy != "" {
		m["triggerBy"] = p.TriggerBy
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
	if p.StopLossTriggerBy != "" {
		m["stopLossTriggerBy"] = p.StopLossTriggerBy
	}
	if p.TakeProfitTriggerPrice != "" {
		m["takeProfitTriggerPrice"] = p.TakeProfitTriggerPrice
	}
	if p.TakeProfitLimitPrice != "" {
		m["takeProfitLimitPrice"] = p.TakeProfitLimitPrice
	}
	if p.TakeProfitTriggerBy != "" {
		m["takeProfitTriggerBy"] = p.TakeProfitTriggerBy
	}
	if p.SlippageTolerance != "" {
		m["slippageTolerance"] = p.SlippageTolerance
	}
	if p.SlippageToleranceType != "" {
		m["slippageToleranceType"] = string(p.SlippageToleranceType)
	}

	return m
}

// GetOrderParams represents parameters for getting an order.
type GetOrderParams struct {
	Symbol   string  `json:"symbol"`
	OrderID  string  `json:"orderId,omitempty"`
	ClientID *uint32 `json:"clientId,omitempty"`
}

// CancelOrderParams represents parameters for canceling an order.
type CancelOrderParams struct {
	Symbol   string  `json:"symbol"`
	OrderID  string  `json:"orderId,omitempty"`
	ClientID *uint32 `json:"clientId,omitempty"`
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

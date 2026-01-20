package enums

// OrderType represents the type of order.
type OrderType string

const (
	OrderTypeMarket OrderType = "Market"
	OrderTypeLimit  OrderType = "Limit"
)

// TimeInForce represents order time in force.
type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "GTC" // Good Till Cancelled
	TimeInForceIOC TimeInForce = "IOC" // Immediate Or Cancel
	TimeInForceFOK TimeInForce = "FOK" // Fill Or Kill
)

// OrderStatus represents the status of an order.
type OrderStatus string

const (
	OrderStatusNew             OrderStatus = "New"
	OrderStatusPartiallyFilled OrderStatus = "PartiallyFilled"
	OrderStatusFilled          OrderStatus = "Filled"
	OrderStatusCancelled       OrderStatus = "Cancelled"
	OrderStatusExpired         OrderStatus = "Expired"
	OrderStatusTriggered       OrderStatus = "Triggered"
)

// SelfTradePrevention represents self-trade prevention mode.
type SelfTradePrevention string

const (
	STPRejectTaker SelfTradePrevention = "RejectTaker"
	STPRejectMaker SelfTradePrevention = "RejectMaker"
	STPRejectBoth  SelfTradePrevention = "RejectBoth"
	STPAllow       SelfTradePrevention = "Allow"
)

// CancelOrderType represents the type of orders to cancel.
type CancelOrderType string

const (
	CancelOrderTypeResting     CancelOrderType = "RestingLimitOrder"
	CancelOrderTypeConditional CancelOrderType = "ConditionalOrder"
)

// TriggerBy represents what triggers a conditional order.
type TriggerBy string

const (
	TriggerByLastPrice TriggerBy = "LastPrice"
	TriggerByMarkPrice TriggerBy = "MarkPrice"
)

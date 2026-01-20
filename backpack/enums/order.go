package enums

// OrderType represents the type of order.
type OrderType string

const (
	OrderTypeMarket OrderType = "Market"
	OrderTypeLimit  OrderType = "Limit"
)

// OrderTypeEnum matches the OpenAPI enum name.
type OrderTypeEnum = OrderType

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
	OrderStatusTriggerPending  OrderStatus = "TriggerPending"
	OrderStatusTriggerFailed   OrderStatus = "TriggerFailed"
)

// SelfTradePrevention represents self-trade prevention mode.
type SelfTradePrevention string

const (
	SelfTradePreventionRejectTaker SelfTradePrevention = "RejectTaker"
	SelfTradePreventionRejectMaker SelfTradePrevention = "RejectMaker"
	SelfTradePreventionRejectBoth  SelfTradePrevention = "RejectBoth"
)

// CancelOrderType represents the type of orders to cancel.
type CancelOrderType string

const (
	CancelOrderTypeResting     CancelOrderType = "RestingLimitOrder"
	CancelOrderTypeConditional CancelOrderType = "ConditionalOrder"
)

// CancelOrderTypeEnum matches the OpenAPI enum name.
type CancelOrderTypeEnum = CancelOrderType

// OrderExpiryReason represents the reason an order expired.
type OrderExpiryReason string

const (
	OrderExpiryReasonAccountTradingSuspended       OrderExpiryReason = "AccountTradingSuspended"
	OrderExpiryReasonBorrowRequiresLendRedeem      OrderExpiryReason = "BorrowRequiresLendRedeem"
	OrderExpiryReasonFillOrKill                    OrderExpiryReason = "FillOrKill"
	OrderExpiryReasonInsufficientBorrowableQuantity OrderExpiryReason = "InsufficientBorrowableQuantity"
	OrderExpiryReasonInsufficientFunds             OrderExpiryReason = "InsufficientFunds"
	OrderExpiryReasonInsufficientLiquidity         OrderExpiryReason = "InsufficientLiquidity"
	OrderExpiryReasonInvalidPrice                  OrderExpiryReason = "InvalidPrice"
	OrderExpiryReasonInvalidQuantity               OrderExpiryReason = "InvalidQuantity"
	OrderExpiryReasonImmediateOrCancel             OrderExpiryReason = "ImmediateOrCancel"
	OrderExpiryReasonInsufficientMargin            OrderExpiryReason = "InsufficientMargin"
	OrderExpiryReasonLiquidation                   OrderExpiryReason = "Liquidation"
	OrderExpiryReasonNegativeEquity                OrderExpiryReason = "NegativeEquity"
	OrderExpiryReasonPostOnlyMode                  OrderExpiryReason = "PostOnlyMode"
	OrderExpiryReasonPostOnlyTaker                 OrderExpiryReason = "PostOnlyTaker"
	OrderExpiryReasonPriceOutOfBounds              OrderExpiryReason = "PriceOutOfBounds"
	OrderExpiryReasonReduceOnlyNotReduced          OrderExpiryReason = "ReduceOnlyNotReduced"
	OrderExpiryReasonSelfTradePrevention           OrderExpiryReason = "SelfTradePrevention"
	OrderExpiryReasonStopWithoutPosition           OrderExpiryReason = "StopWithoutPosition"
	OrderExpiryReasonPriceImpact                   OrderExpiryReason = "PriceImpact"
	OrderExpiryReasonUnknown                       OrderExpiryReason = "Unknown"
	OrderExpiryReasonUserPermissions               OrderExpiryReason = "UserPermissions"
	OrderExpiryReasonMaxStopOrdersPerPosition      OrderExpiryReason = "MaxStopOrdersPerPosition"
	OrderExpiryReasonPositionLimit                 OrderExpiryReason = "PositionLimit"
	OrderExpiryReasonSlippageToleranceExceeded     OrderExpiryReason = "SlippageToleranceExceeded"
)

// SystemOrderType represents the type of system order.
type SystemOrderType string

const (
	SystemOrderTypeCollateralConversion       SystemOrderType = "CollateralConversion"
	SystemOrderTypeFutureExpiry               SystemOrderType = "FutureExpiry"
	SystemOrderTypeLiquidatePositionOnAdl     SystemOrderType = "LiquidatePositionOnAdl"
	SystemOrderTypeLiquidatePositionOnBook    SystemOrderType = "LiquidatePositionOnBook"
	SystemOrderTypeLiquidatePositionOnBackstop SystemOrderType = "LiquidatePositionOnBackstop"
	SystemOrderTypeOrderBookClosed            SystemOrderType = "OrderBookClosed"
)

// SlippageToleranceType represents the type of slippage tolerance.
type SlippageToleranceType string

const (
	SlippageToleranceTypeTickSize SlippageToleranceType = "TickSize"
	SlippageToleranceTypePercent  SlippageToleranceType = "Percent"
)

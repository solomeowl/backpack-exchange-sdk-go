package enums

// RFQStatus matches the OpenAPI status type for RFQ history.
type RFQStatus = OrderStatus

// QuoteStatus matches the OpenAPI status type for quote history.
type QuoteStatus = OrderStatus

// RfqExecutionMode represents the RFQ execution mode.
type RfqExecutionMode string

const (
	RfqExecutionModeAwaitAccept RfqExecutionMode = "AwaitAccept"
	RfqExecutionModeImmediate   RfqExecutionMode = "Immediate"
)

// RfqFillType represents the RFQ fill type filter.
type RfqFillType string

const (
	RfqFillTypeUser                 RfqFillType = "User"
	RfqFillTypeCollateralConversion RfqFillType = "CollateralConversion"
)

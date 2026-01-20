package enums

// RFQStatus represents the status of an RFQ (uses OrderStatus).
type RFQStatus string

const (
	RFQStatusNew       RFQStatus = "New"
	RFQStatusFilled    RFQStatus = "Filled"
	RFQStatusCancelled RFQStatus = "Cancelled"
	RFQStatusExpired   RFQStatus = "Expired"
)

// QuoteStatus represents the status of a quote (uses OrderStatus).
type QuoteStatus string

const (
	QuoteStatusNew       QuoteStatus = "New"
	QuoteStatusFilled    QuoteStatus = "Filled"
	QuoteStatusCancelled QuoteStatus = "Cancelled"
	QuoteStatusExpired   QuoteStatus = "Expired"
)

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

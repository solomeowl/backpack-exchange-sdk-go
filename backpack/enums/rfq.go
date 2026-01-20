package enums

// RFQStatus represents the status of an RFQ.
type RFQStatus string

const (
	RFQStatusOpen       RFQStatus = "Open"
	RFQStatusFilled     RFQStatus = "Filled"
	RFQStatusCancelled  RFQStatus = "Cancelled"
	RFQStatusExpired    RFQStatus = "Expired"
)

// QuoteStatus represents the status of a quote.
type QuoteStatus string

const (
	QuoteStatusPending  QuoteStatus = "Pending"
	QuoteStatusAccepted QuoteStatus = "Accepted"
	QuoteStatusRejected QuoteStatus = "Rejected"
	QuoteStatusExpired  QuoteStatus = "Expired"
)

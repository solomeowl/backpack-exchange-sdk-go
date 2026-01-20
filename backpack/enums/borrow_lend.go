package enums

// BorrowLendSide represents the side of a borrow/lend operation.
type BorrowLendSide string

const (
	BorrowLendSideBorrow BorrowLendSide = "Borrow"
	BorrowLendSideLend   BorrowLendSide = "Lend"
)

// BorrowLendEventType represents the type of borrow/lend event.
type BorrowLendEventType string

const (
	BorrowLendEventBorrow    BorrowLendEventType = "Borrow"
	BorrowLendEventRepay     BorrowLendEventType = "Repay"
	BorrowLendEventLend      BorrowLendEventType = "Lend"
	BorrowLendEventRedeem    BorrowLendEventType = "Redeem"
	BorrowLendEventInterest  BorrowLendEventType = "Interest"
)

// BorrowLendMarketHistoryInterval represents the interval for market history.
type BorrowLendMarketHistoryInterval string

const (
	BorrowLendMarketHistoryInterval1d     BorrowLendMarketHistoryInterval = "1d"
	BorrowLendMarketHistoryInterval1w     BorrowLendMarketHistoryInterval = "1w"
	BorrowLendMarketHistoryInterval1month BorrowLendMarketHistoryInterval = "1month"
	BorrowLendMarketHistoryInterval1year  BorrowLendMarketHistoryInterval = "1year"
)

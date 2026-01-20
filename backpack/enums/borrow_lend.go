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
	BorrowLendEventTypeBorrow      BorrowLendEventType = "Borrow"
	BorrowLendEventTypeBorrowRepay BorrowLendEventType = "BorrowRepay"
	BorrowLendEventTypeLend        BorrowLendEventType = "Lend"
	BorrowLendEventTypeLendRedeem  BorrowLendEventType = "LendRedeem"
)

// BorrowLendMarketHistoryInterval represents the interval for market history.
type BorrowLendMarketHistoryInterval string

const (
	BorrowLendMarketHistoryInterval1d     BorrowLendMarketHistoryInterval = "1d"
	BorrowLendMarketHistoryInterval1w     BorrowLendMarketHistoryInterval = "1w"
	BorrowLendMarketHistoryInterval1month BorrowLendMarketHistoryInterval = "1month"
	BorrowLendMarketHistoryInterval1year  BorrowLendMarketHistoryInterval = "1year"
)

// BorrowLendBookState represents the state of a borrow/lend book.
type BorrowLendBookState string

const (
	BorrowLendBookStateOpen      BorrowLendBookState = "Open"
	BorrowLendBookStateClosed    BorrowLendBookState = "Closed"
	BorrowLendBookStateRepayOnly BorrowLendBookState = "RepayOnly"
)

// BorrowLendPositionState represents the state of a borrow/lend position.
type BorrowLendPositionState string

const (
	BorrowLendPositionStateOpen   BorrowLendPositionState = "Open"
	BorrowLendPositionStateClosed BorrowLendPositionState = "Closed"
)

// BorrowLendSource represents the source of a borrow/lend movement.
type BorrowLendSource string

const (
	BorrowLendSourceAdlProvider        BorrowLendSource = "AdlProvider"
	BorrowLendSourceAutoBorrowRepay    BorrowLendSource = "AutoBorrowRepay"
	BorrowLendSourceAutoLend           BorrowLendSource = "AutoLend"
	BorrowLendSourceBackstopProvider   BorrowLendSource = "BackstopProvider"
	BorrowLendSourceDustConversion     BorrowLendSource = "DustConversion"
	BorrowLendSourceInterest           BorrowLendSource = "Interest"
	BorrowLendSourceLiquidation        BorrowLendSource = "Liquidation"
	BorrowLendSourceLiquidationAdl     BorrowLendSource = "LiquidationAdl"
	BorrowLendSourceLiquidationBackstop BorrowLendSource = "LiquidationBackstop"
	BorrowLendSourceManual             BorrowLendSource = "Manual"
	BorrowLendSourceRealization        BorrowLendSource = "Realization"
)

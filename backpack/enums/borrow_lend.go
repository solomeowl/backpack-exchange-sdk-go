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

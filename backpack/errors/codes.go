package errors

// ErrorCode represents known API error codes.
type ErrorCode string

// Known error codes from Backpack Exchange API.
const (
	// Authentication errors
	ErrCodeInvalidSignature   ErrorCode = "INVALID_SIGNATURE"
	ErrCodeInvalidAPIKey      ErrorCode = "INVALID_API_KEY"
	ErrCodeExpiredTimestamp   ErrorCode = "EXPIRED_TIMESTAMP"
	ErrCodeInvalidTimestamp   ErrorCode = "INVALID_TIMESTAMP"
	ErrCodeUnauthorized       ErrorCode = "UNAUTHORIZED"

	// Order errors
	ErrCodeInsufficientFunds  ErrorCode = "INSUFFICIENT_FUNDS"
	ErrCodeOrderNotFound      ErrorCode = "ORDER_NOT_FOUND"
	ErrCodeInvalidOrderType   ErrorCode = "INVALID_ORDER_TYPE"
	ErrCodeInvalidSide        ErrorCode = "INVALID_SIDE"
	ErrCodeInvalidQuantity    ErrorCode = "INVALID_QUANTITY"
	ErrCodeInvalidPrice       ErrorCode = "INVALID_PRICE"
	ErrCodeMinNotional        ErrorCode = "MIN_NOTIONAL"
	ErrCodeMaxOrders          ErrorCode = "MAX_ORDERS"
	ErrCodeSelfTrade          ErrorCode = "SELF_TRADE"
	ErrCodePostOnlyRejected   ErrorCode = "POST_ONLY_REJECTED"

	// Market errors
	ErrCodeMarketNotFound     ErrorCode = "MARKET_NOT_FOUND"
	ErrCodeMarketClosed       ErrorCode = "MARKET_CLOSED"

	// Rate limiting
	ErrCodeRateLimited        ErrorCode = "RATE_LIMITED"
	ErrCodeTooManyRequests    ErrorCode = "TOO_MANY_REQUESTS"

	// System errors
	ErrCodeInternalError      ErrorCode = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"

	// Withdrawal errors
	ErrCodeWithdrawalDisabled ErrorCode = "WITHDRAWAL_DISABLED"
	ErrCodeInvalidAddress     ErrorCode = "INVALID_ADDRESS"
	ErrCodeMinWithdrawal      ErrorCode = "MIN_WITHDRAWAL"

	// Position errors
	ErrCodePositionNotFound   ErrorCode = "POSITION_NOT_FOUND"
	ErrCodeMaxLeverage        ErrorCode = "MAX_LEVERAGE"
	ErrCodeLiquidation        ErrorCode = "LIQUIDATION"
)

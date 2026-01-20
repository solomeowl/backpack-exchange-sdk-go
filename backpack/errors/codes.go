package errors

// ApiErrorCode represents API error codes defined by the OpenAPI spec.
type ApiErrorCode string

// ErrorCode is kept for backwards compatibility with previous SDK versions.
type ErrorCode = ApiErrorCode

// Known error codes from Backpack Exchange API.
const (
	ErrCodeAccountDeactivated        ApiErrorCode = "ACCOUNT_DEACTIVATED"
	ErrCodeAccountLiquidating        ApiErrorCode = "ACCOUNT_LIQUIDATING"
	ErrCodeBorrowLimit               ApiErrorCode = "BORROW_LIMIT"
	ErrCodeBorrowRequiresLendRedeem  ApiErrorCode = "BORROW_REQUIRES_LEND_REDEEM"
	ErrCodeForbidden                 ApiErrorCode = "FORBIDDEN"
	ErrCodeInsufficientFunds         ApiErrorCode = "INSUFFICIENT_FUNDS"
	ErrCodeInsufficientMargin        ApiErrorCode = "INSUFFICIENT_MARGIN"
	ErrCodeInsufficientSupply        ApiErrorCode = "INSUFFICIENT_SUPPLY"
	ErrCodeInvalidAsset              ApiErrorCode = "INVALID_ASSET"
	ErrCodeInvalidClientRequest      ApiErrorCode = "INVALID_CLIENT_REQUEST"
	ErrCodeInvalidMarket             ApiErrorCode = "INVALID_MARKET"
	ErrCodeInvalidOrder              ApiErrorCode = "INVALID_ORDER"
	ErrCodeInvalidPositionID         ApiErrorCode = "INVALID_POSITION_ID"
	ErrCodeInvalidQuantity           ApiErrorCode = "INVALID_QUANTITY"
	ErrCodeInvalidRange              ApiErrorCode = "INVALID_RANGE"
	ErrCodeInvalidSignature          ApiErrorCode = "INVALID_SIGNATURE"
	ErrCodeInvalidSource             ApiErrorCode = "INVALID_SOURCE"
	ErrCodeInvalidSymbol             ApiErrorCode = "INVALID_SYMBOL"
	ErrCodeInvalidTwoFactorCode      ApiErrorCode = "INVALID_TWO_FACTOR_CODE"
	ErrCodeLendLimit                 ApiErrorCode = "LEND_LIMIT"
	ErrCodeLendRequiresBorrowRepay   ApiErrorCode = "LEND_REQUIRES_BORROW_REPAY"
	ErrCodeMaintenance               ApiErrorCode = "MAINTENANCE"
	ErrCodeMaxLeverageReached        ApiErrorCode = "MAX_LEVERAGE_REACHED"
	ErrCodeNotImplemented            ApiErrorCode = "NOT_IMPLEMENTED"
	ErrCodeOrderLimit                ApiErrorCode = "ORDER_LIMIT"
	ErrCodePositionLimit             ApiErrorCode = "POSITION_LIMIT"
	ErrCodePreconditionFailed        ApiErrorCode = "PRECONDITION_FAILED"
	ErrCodeResourceNotFound          ApiErrorCode = "RESOURCE_NOT_FOUND"
	ErrCodeServerError               ApiErrorCode = "SERVER_ERROR"
	ErrCodeTimeout                   ApiErrorCode = "TIMEOUT"
	ErrCodeTooManyRequests           ApiErrorCode = "TOO_MANY_REQUESTS"
	ErrCodeTradingPaused             ApiErrorCode = "TRADING_PAUSED"
	ErrCodeUnauthorized              ApiErrorCode = "UNAUTHORIZED"
)

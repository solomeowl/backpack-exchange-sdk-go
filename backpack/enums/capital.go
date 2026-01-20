package enums

// WithdrawalStatus represents the status of a withdrawal.
type WithdrawalStatus string

const (
	WithdrawalStatusPending   WithdrawalStatus = "Pending"
	WithdrawalStatusConfirmed WithdrawalStatus = "Confirmed"
	WithdrawalStatusCompleted WithdrawalStatus = "Completed"
	WithdrawalStatusFailed    WithdrawalStatus = "Failed"
	WithdrawalStatusCancelled WithdrawalStatus = "Cancelled"
)

// DepositStatus represents the status of a deposit.
type DepositStatus string

const (
	DepositStatusPending   DepositStatus = "Pending"
	DepositStatusConfirmed DepositStatus = "Confirmed"
	DepositStatusCompleted DepositStatus = "Completed"
)

// TwoFactorType represents the type of 2FA.
type TwoFactorType string

const (
	TwoFactorTypeTOTP TwoFactorType = "totp"
)

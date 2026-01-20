package enums

// WithdrawalStatus represents the status of a withdrawal.
type WithdrawalStatus string

const (
	WithdrawalStatusConfirmed                    WithdrawalStatus = "confirmed"
	WithdrawalStatusOwnershipVerificationRequired WithdrawalStatus = "ownershipVerificationRequired"
	WithdrawalStatusPending                      WithdrawalStatus = "pending"
	WithdrawalStatusRecipientInformationProvided WithdrawalStatus = "recipientInformationProvided"
	WithdrawalStatusRecipientInformationRequired WithdrawalStatus = "recipientInformationRequired"
)

// DepositStatus represents the status of a deposit.
type DepositStatus string

const (
	DepositStatusCancelled                      DepositStatus = "cancelled"
	DepositStatusConfirmed                      DepositStatus = "confirmed"
	DepositStatusDeclined                       DepositStatus = "declined"
	DepositStatusExpired                        DepositStatus = "expired"
	DepositStatusInitiated                      DepositStatus = "initiated"
	DepositStatusOwnershipVerificationRequired  DepositStatus = "ownershipVerificationRequired"
	DepositStatusPending                        DepositStatus = "pending"
	DepositStatusRefunded                       DepositStatus = "refunded"
	DepositStatusSenderVerificationCompleted    DepositStatus = "senderVerificationCompleted"
	DepositStatusSenderVerificationRequired     DepositStatus = "senderVerificationRequired"
)

// DepositSource represents the source of a deposit (blockchain or payment processor).
type DepositSource string

const (
	DepositSourceAdministrator DepositSource = "administrator"
	DepositSource0G            DepositSource = "0G"
	DepositSourceAptos         DepositSource = "aptos"
	DepositSourceArbitrum      DepositSource = "arbitrum"
	DepositSourceAvalanche     DepositSource = "avalanche"
	DepositSourceBase          DepositSource = "base"
	DepositSourceBerachain     DepositSource = "berachain"
	DepositSourceBitcoin       DepositSource = "bitcoin"
	DepositSourceBitcoinCash   DepositSource = "bitcoinCash"
	DepositSourceBSC           DepositSource = "bsc"
	DepositSourceCardano       DepositSource = "cardano"
	DepositSourceDogecoin      DepositSource = "dogecoin"
	DepositSourceEclipse       DepositSource = "eclipse"
	DepositSourceEthereum      DepositSource = "ethereum"
	DepositSourceFogo          DepositSource = "fogo"
	DepositSourceHyperEVM      DepositSource = "hyperEVM"
	DepositSourceHyperliquid   DepositSource = "hyperliquid"
	DepositSourceLinea         DepositSource = "linea"
	DepositSourceLitecoin      DepositSource = "litecoin"
	DepositSourceMonad         DepositSource = "monad"
	DepositSourceNear          DepositSource = "near"
	DepositSourcePolygon       DepositSource = "polygon"
	DepositSourceOptimism      DepositSource = "optimism"
	DepositSourcePlasma        DepositSource = "plasma"
	DepositSourceSei           DepositSource = "sei"
	DepositSourceStable        DepositSource = "stable"
	DepositSourceSui           DepositSource = "sui"
	DepositSourceSolana        DepositSource = "solana"
	DepositSourceStory         DepositSource = "story"
	DepositSourceTron          DepositSource = "tron"
	DepositSourceXRP           DepositSource = "xRP"
	DepositSourceZcash         DepositSource = "zcash"
	DepositSourceEqualsMoney   DepositSource = "equalsMoney"
	DepositSourceBanxa         DepositSource = "banxa"
	DepositSourceInternal      DepositSource = "internal"
)

// EqualsMoneyWithdrawalState represents the fiat withdrawal state for Equals Money.
type EqualsMoneyWithdrawalState string

const (
	EqualsMoneyWithdrawalStateInitialized                 EqualsMoneyWithdrawalState = "initialized"
	EqualsMoneyWithdrawalStatePending                     EqualsMoneyWithdrawalState = "pending"
	EqualsMoneyWithdrawalStateFulfilling                  EqualsMoneyWithdrawalState = "fulfilling"
	EqualsMoneyWithdrawalStateProcessing                  EqualsMoneyWithdrawalState = "processing"
	EqualsMoneyWithdrawalStateComplete                    EqualsMoneyWithdrawalState = "complete"
	EqualsMoneyWithdrawalStateDeclined                    EqualsMoneyWithdrawalState = "declined"
	EqualsMoneyWithdrawalStateCancelled                   EqualsMoneyWithdrawalState = "cancelled"
	EqualsMoneyWithdrawalStateReview                      EqualsMoneyWithdrawalState = "review"
	EqualsMoneyWithdrawalStateAwaitingDocuments           EqualsMoneyWithdrawalState = "awaitingDocuments"
	EqualsMoneyWithdrawalStateAwaitingComplianceQuestions EqualsMoneyWithdrawalState = "awaitingComplianceQuestions"
	EqualsMoneyWithdrawalStateRefundedInternal            EqualsMoneyWithdrawalState = "refundedInternal"
	EqualsMoneyWithdrawalStateRefundedExternal            EqualsMoneyWithdrawalState = "refundedExternal"
)

// TwoFactorType represents the type of 2FA.
type TwoFactorType string

const (
	TwoFactorTypeTOTP TwoFactorType = "totp"
)

// PaymentType represents the type of interest payment.
type PaymentType string

const (
	PaymentTypeEntryFee             PaymentType = "EntryFee"
	PaymentTypeBorrow               PaymentType = "Borrow"
	PaymentTypeLend                 PaymentType = "Lend"
	PaymentTypeUnrealizedPositivePnl PaymentType = "UnrealizedPositivePnl"
	PaymentTypeUnrealizedNegativePnl PaymentType = "UnrealizedNegativePnl"
)

// InterestPaymentSource represents the source of interest payment.
type InterestPaymentSource string

const (
	InterestPaymentSourceUnrealizedPnl InterestPaymentSource = "UnrealizedPnl"
	InterestPaymentSourceBorrowLend    InterestPaymentSource = "BorrowLend"
)

package enums

// FillType represents fill types for fill history filters.
type FillType string

const (
	FillTypeUser                               FillType = "User"
	FillTypeBookLiquidation                    FillType = "BookLiquidation"
	FillTypeAdl                                FillType = "Adl"
	FillTypeBackstop                           FillType = "Backstop"
	FillTypeLiquidation                        FillType = "Liquidation"
	FillTypeAllLiquidation                     FillType = "AllLiquidation"
	FillTypeCollateralConversion               FillType = "CollateralConversion"
	FillTypeCollateralConversionAndSpotLiquidation FillType = "CollateralConversionAndSpotLiquidation"
)

// SettlementSource represents settlement sources.
type SettlementSource string

const (
	SettlementSourceTradingFees                     SettlementSource = "TradingFees"
	SettlementSourceTradingFeesSystem               SettlementSource = "TradingFeesSystem"
	SettlementSourceFundingPayment                  SettlementSource = "FundingPayment"
	SettlementSourceCulledBorrowInterest            SettlementSource = "CulledBorrowInterest"
	SettlementSourceCulledRealizePnlAuto            SettlementSource = "CulledRealizePnlAuto"
	SettlementSourceCulledRealizePnlBookUtilisation SettlementSource = "CulledRealizePnlBookUtilisation"
	SettlementSourceCulledRealizePnlAccountThreshold SettlementSource = "CulledRealizePnlAccountThreshold"
	SettlementSourceCulledRealizePnlSystemThreshold SettlementSource = "CulledRealizePnlSystemThreshold"
	SettlementSourceRealizePnl                      SettlementSource = "RealizePnl"
	SettlementSourceBackstopProviderLiquidation     SettlementSource = "BackstopProviderLiquidation"
	SettlementSourceBackstopAdlLiquidation          SettlementSource = "BackstopAdlLiquidation"
	SettlementSourceBackstopLiquidityFundProceeds   SettlementSource = "BackstopLiquidityFundProceeds"
	SettlementSourceSystemLiabilityTransfer         SettlementSource = "SystemLiabilityTransfer"
)

// SettlementSourceFilter represents settlement history filters.
type SettlementSourceFilter string

const (
	SettlementSourceFilterBackstopLiquidation       SettlementSourceFilter = "BackstopLiquidation"
	SettlementSourceFilterCulledBorrowInterest      SettlementSourceFilter = "CulledBorrowInterest"
	SettlementSourceFilterCulledRealizePnl          SettlementSourceFilter = "CulledRealizePnl"
	SettlementSourceFilterCulledRealizePnlBookUtilization SettlementSourceFilter = "CulledRealizePnlBookUtilization"
	SettlementSourceFilterFundingPayment            SettlementSourceFilter = "FundingPayment"
	SettlementSourceFilterRealizePnl                SettlementSourceFilter = "RealizePnl"
	SettlementSourceFilterTradingFees               SettlementSourceFilter = "TradingFees"
	SettlementSourceFilterTradingFeesSystem         SettlementSourceFilter = "TradingFeesSystem"
)

// SeriesRecurrence represents prediction series recurrence.
type SeriesRecurrence string

const (
	SeriesRecurrenceHourly  SeriesRecurrence = "hourly"
	SeriesRecurrenceDaily   SeriesRecurrence = "daily"
	SeriesRecurrenceWeekly  SeriesRecurrence = "weekly"
	SeriesRecurrenceMonthly SeriesRecurrence = "monthly"
	SeriesRecurrenceAnnual  SeriesRecurrence = "annual"
)

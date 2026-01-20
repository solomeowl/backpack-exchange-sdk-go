package services

import (
	"context"
	"strconv"
	"strings"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// HistoryService provides history operations.
type HistoryService struct {
	client HTTPClient
}

// NewHistoryService creates a new HistoryService.
func NewHistoryService(client HTTPClient) *HistoryService {
	return &HistoryService{client: client}
}

func joinMarketTypes(values []enums.MarketType) string {
	if len(values) == 0 {
		return ""
	}
	parts := make([]string, 0, len(values))
	for _, v := range values {
		parts = append(parts, string(v))
	}
	return strings.Join(parts, ",")
}

func joinBorrowLendSources(values []enums.BorrowLendSource) string {
	if len(values) == 0 {
		return ""
	}
	parts := make([]string, 0, len(values))
	for _, v := range values {
		parts = append(parts, string(v))
	}
	return strings.Join(parts, ",")
}

// GetBorrowHistory retrieves borrow history.
func (s *HistoryService) GetBorrowHistory(ctx context.Context, params *types.BorrowLendHistoryParams) ([]types.BorrowLendMovement, error) {
	var result []types.BorrowLendMovement
	queryParams := make(map[string]string)
	if params != nil {
		if params.Type != "" {
			queryParams["type"] = string(params.Type)
		}
		if sources := joinBorrowLendSources(params.Sources); sources != "" {
			queryParams["sources"] = sources
		}
		if params.PositionID != "" {
			queryParams["positionId"] = params.PositionID
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/borrowLend", queryParams, "borrowHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInterestHistory retrieves interest payment history.
func (s *HistoryService) GetInterestHistory(ctx context.Context, params *types.InterestHistoryParams) ([]types.InterestHistoryItem, error) {
	var result []types.InterestHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.Asset != "" {
			queryParams["asset"] = params.Asset
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.PositionID != "" {
			queryParams["positionId"] = params.PositionID
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.Source != "" {
			queryParams["source"] = string(params.Source)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/interest", queryParams, "interestHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBorrowPositionHistory retrieves borrow position history.
func (s *HistoryService) GetBorrowPositionHistory(ctx context.Context, params *types.BorrowLendPositionHistoryParams) ([]types.BorrowLendPositionHistoryRow, error) {
	var result []types.BorrowLendPositionHistoryRow
	queryParams := make(map[string]string)
	if params != nil {
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Side != "" {
			queryParams["side"] = string(params.Side)
		}
		if params.State != "" {
			queryParams["state"] = string(params.State)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/borrowLend/positions", queryParams, "borrowPositionHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFillHistory retrieves fill/trade history.
func (s *HistoryService) GetFillHistory(ctx context.Context, params *types.FillHistoryParams) ([]types.Fill, error) {
	var result []types.Fill
	queryParams := make(map[string]string)
	if params != nil {
		if params.OrderID != "" {
			queryParams["orderId"] = params.OrderID
		}
		if params.StrategyID != "" {
			queryParams["strategyId"] = params.StrategyID
		}
		if params.From > 0 {
			queryParams["from"] = strconv.FormatInt(params.From, 10)
		}
		if params.To > 0 {
			queryParams["to"] = strconv.FormatInt(params.To, 10)
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.FillType != "" {
			queryParams["fillType"] = string(params.FillType)
		}
		if marketTypes := joinMarketTypes(params.MarketType); marketTypes != "" {
			queryParams["marketType"] = marketTypes
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/fills", queryParams, "fillHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFundingPayments retrieves funding payment history.
func (s *HistoryService) GetFundingPayments(ctx context.Context, params *types.FundingHistoryParams) ([]types.FundingPayment, error) {
	var result []types.FundingPayment
	queryParams := make(map[string]string)
	if params != nil {
		if params.SubaccountID != nil {
			queryParams["subaccountId"] = strconv.FormatInt(int64(*params.SubaccountID), 10)
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/funding", queryParams, "fundingHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOrderHistory retrieves order history.
func (s *HistoryService) GetOrderHistory(ctx context.Context, params *types.OrderHistoryParams) ([]types.OrderHistoryItem, error) {
	var result []types.OrderHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.OrderID != "" {
			queryParams["orderId"] = params.OrderID
		}
		if params.StrategyID != "" {
			queryParams["strategyId"] = params.StrategyID
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if marketTypes := joinMarketTypes(params.MarketType); marketTypes != "" {
			queryParams["marketType"] = marketTypes
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/orders", queryParams, "orderHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSettlementHistory retrieves settlement history.
func (s *HistoryService) GetSettlementHistory(ctx context.Context, params *types.SettlementHistoryParams) ([]types.Settlement, error) {
	var result []types.Settlement
	queryParams := make(map[string]string)
	if params != nil {
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.Source != "" {
			queryParams["source"] = string(params.Source)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/settlement", queryParams, "settlementHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDustHistory retrieves dust conversion history.
func (s *HistoryService) GetDustHistory(ctx context.Context, params *types.DustHistoryParams) ([]types.DustHistoryItem, error) {
	var result []types.DustHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.ID != nil {
			queryParams["id"] = strconv.FormatInt(*params.ID, 10)
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/dust", queryParams, "dustHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetRFQHistory retrieves RFQ history.
func (s *HistoryService) GetRFQHistory(ctx context.Context, params *types.RFQHistoryParams) ([]types.RFQHistoryItem, error) {
	var result []types.RFQHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.RFQID != "" {
			queryParams["rfqId"] = params.RFQID
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Status != "" {
			queryParams["status"] = string(params.Status)
		}
		if params.Side != "" {
			queryParams["side"] = string(params.Side)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/rfq", queryParams, "rfqHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuoteHistory retrieves quote history.
func (s *HistoryService) GetQuoteHistory(ctx context.Context, params *types.QuoteHistoryParams) ([]types.QuoteHistoryItem, error) {
	var result []types.QuoteHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.QuoteID != "" {
			queryParams["quoteId"] = params.QuoteID
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Status != "" {
			queryParams["status"] = string(params.Status)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/quote", queryParams, "quoteHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetRFQFillHistory retrieves RFQ fill history.
func (s *HistoryService) GetRFQFillHistory(ctx context.Context, params *types.RFQFillHistoryParams) ([]types.RFQFillHistoryItem, error) {
	var result []types.RFQFillHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.QuoteID != "" {
			queryParams["quoteId"] = params.QuoteID
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Side != "" {
			queryParams["side"] = string(params.Side)
		}
		if params.FillType != "" {
			queryParams["fillType"] = string(params.FillType)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/rfq/fill", queryParams, "rfqFillHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuoteFillHistory retrieves quote fill history.
func (s *HistoryService) GetQuoteFillHistory(ctx context.Context, params *types.QuoteFillHistoryParams) ([]types.QuoteFillHistoryItem, error) {
	var result []types.QuoteFillHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.QuoteID != "" {
			queryParams["quoteId"] = params.QuoteID
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Side != "" {
			queryParams["side"] = string(params.Side)
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/quote/fill", queryParams, "quoteFillHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetStrategyHistory retrieves strategy history.
func (s *HistoryService) GetStrategyHistory(ctx context.Context, params *types.StrategyHistoryParams) ([]types.StrategyHistoryItem, error) {
	var result []types.StrategyHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.StrategyID != "" {
			queryParams["strategyId"] = params.StrategyID
		}
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if marketTypes := joinMarketTypes(params.MarketType); marketTypes != "" {
			queryParams["marketType"] = marketTypes
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/strategies", queryParams, "strategyHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPositionHistory retrieves position history.
func (s *HistoryService) GetPositionHistory(ctx context.Context, params *types.PositionHistoryParams) ([]types.PositionHistoryItem, error) {
	var result []types.PositionHistoryItem
	queryParams := make(map[string]string)
	if params != nil {
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.State != "" {
			queryParams["state"] = string(params.State)
		}
		if marketTypes := joinMarketTypes(params.MarketType); marketTypes != "" {
			queryParams["marketType"] = marketTypes
		}
		if params.Limit > 0 {
			queryParams["limit"] = strconv.Itoa(params.Limit)
		}
		if params.Offset > 0 {
			queryParams["offset"] = strconv.Itoa(params.Offset)
		}
		if params.SortDirection != "" {
			queryParams["sortDirection"] = string(params.SortDirection)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/position", queryParams, "positionHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

package services

import (
	"context"
	"strconv"

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

func (s *HistoryService) buildQueryParams(params *types.HistoryParams) map[string]string {
	queryParams := make(map[string]string)
	if params == nil {
		return queryParams
	}
	if params.Symbol != "" {
		queryParams["symbol"] = params.Symbol
	}
	if params.From > 0 {
		queryParams["from"] = strconv.FormatInt(params.From, 10)
	}
	if params.To > 0 {
		queryParams["to"] = strconv.FormatInt(params.To, 10)
	}
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(params.Limit)
	}
	if params.Offset > 0 {
		queryParams["offset"] = strconv.Itoa(params.Offset)
	}
	if params.Direction != "" {
		queryParams["direction"] = string(params.Direction)
	}
	return queryParams
}

// GetBorrowHistory retrieves borrow history.
func (s *HistoryService) GetBorrowHistory(ctx context.Context, params *types.HistoryParams) ([]types.BorrowHistoryItem, error) {
	var result []types.BorrowHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/borrowLend", s.buildQueryParams(params), "borrowHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetInterestHistory retrieves interest payment history.
func (s *HistoryService) GetInterestHistory(ctx context.Context, params *types.HistoryParams) ([]types.InterestHistoryItem, error) {
	var result []types.InterestHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/interest", s.buildQueryParams(params), "interestHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBorrowPositionHistory retrieves borrow position history.
func (s *HistoryService) GetBorrowPositionHistory(ctx context.Context, params *types.HistoryParams) ([]types.BorrowLendPosition, error) {
	var result []types.BorrowLendPosition
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/borrowLend/positions", s.buildQueryParams(params), "borrowPositionHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFillHistory retrieves fill/trade history.
func (s *HistoryService) GetFillHistory(ctx context.Context, params *types.HistoryParams) ([]types.Fill, error) {
	var result []types.Fill
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/fills", s.buildQueryParams(params), "fillHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetFundingPayments retrieves funding payment history.
func (s *HistoryService) GetFundingPayments(ctx context.Context, params *types.HistoryParams) ([]types.FundingPayment, error) {
	var result []types.FundingPayment
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/funding", s.buildQueryParams(params), "fundingHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOrderHistory retrieves order history.
func (s *HistoryService) GetOrderHistory(ctx context.Context, params *types.HistoryParams) ([]types.OrderHistoryItem, error) {
	var result []types.OrderHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/orders", s.buildQueryParams(params), "orderHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetSettlementHistory retrieves settlement history.
func (s *HistoryService) GetSettlementHistory(ctx context.Context, params *types.HistoryParams) ([]types.Settlement, error) {
	var result []types.Settlement
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/settlement", s.buildQueryParams(params), "settlementHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetDustHistory retrieves dust conversion history.
func (s *HistoryService) GetDustHistory(ctx context.Context, params *types.HistoryParams) ([]types.DustHistoryItem, error) {
	var result []types.DustHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/dust", s.buildQueryParams(params), "dustHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetRFQHistory retrieves RFQ history.
func (s *HistoryService) GetRFQHistory(ctx context.Context, params *types.HistoryParams) ([]types.RFQHistoryItem, error) {
	var result []types.RFQHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/rfq", s.buildQueryParams(params), "rfqHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuoteHistory retrieves quote history.
func (s *HistoryService) GetQuoteHistory(ctx context.Context, params *types.HistoryParams) ([]types.QuoteHistoryItem, error) {
	var result []types.QuoteHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/quote", s.buildQueryParams(params), "quoteHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetRFQFillHistory retrieves RFQ fill history.
func (s *HistoryService) GetRFQFillHistory(ctx context.Context, params *types.HistoryParams) ([]types.RFQFillHistoryItem, error) {
	var result []types.RFQFillHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/rfq/fill", s.buildQueryParams(params), "rfqFillHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuoteFillHistory retrieves quote fill history.
func (s *HistoryService) GetQuoteFillHistory(ctx context.Context, params *types.HistoryParams) ([]types.QuoteFillHistoryItem, error) {
	var result []types.QuoteFillHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/quote/fill", s.buildQueryParams(params), "quoteFillHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetStrategyHistory retrieves strategy history.
func (s *HistoryService) GetStrategyHistory(ctx context.Context, params *types.HistoryParams) ([]types.StrategyHistoryItem, error) {
	var result []types.StrategyHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/strategies", s.buildQueryParams(params), "strategyHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPositionHistory retrieves position history.
func (s *HistoryService) GetPositionHistory(ctx context.Context, params *types.HistoryParams) ([]types.PositionHistoryItem, error) {
	var result []types.PositionHistoryItem
	if err := s.client.GetAuthenticated(ctx, "wapi/v1/history/position", s.buildQueryParams(params), "positionHistoryQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

package services

import (
	"context"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// RFQService provides RFQ (Request for Quote) operations.
type RFQService struct {
	client HTTPClient
}

// NewRFQService creates a new RFQService.
func NewRFQService(client HTTPClient) *RFQService {
	return &RFQService{client: client}
}

// SubmitRFQ submits a new RFQ.
func (s *RFQService) SubmitRFQ(ctx context.Context, params types.RFQSubmitParams) (*types.RFQ, error) {
	var result types.RFQ
	body := map[string]any{
		"symbol": params.Symbol,
		"side":   string(params.Side),
	}
	if params.ClientID != nil {
		body["clientId"] = *params.ClientID
	}
	if params.Quantity != "" {
		body["quantity"] = params.Quantity
	}
	if params.QuoteQuantity != "" {
		body["quoteQuantity"] = params.QuoteQuantity
	}
	if params.Price != "" {
		body["price"] = params.Price
	}
	if params.ExecutionMode != "" {
		body["executionMode"] = string(params.ExecutionMode)
	}
	if params.AutoLend != nil {
		body["autoLend"] = *params.AutoLend
	}
	if params.AutoLendRedeem != nil {
		body["autoLendRedeem"] = *params.AutoLendRedeem
	}
	if params.AutoBorrow != nil {
		body["autoBorrow"] = *params.AutoBorrow
	}
	if params.AutoBorrowRepay != nil {
		body["autoBorrowRepay"] = *params.AutoBorrowRepay
	}
	if err := s.client.PostAuthenticated(ctx, "api/v1/rfq", body, "rfqSubmit", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SubmitQuote submits a quote for an RFQ.
func (s *RFQService) SubmitQuote(ctx context.Context, params types.QuoteSubmitParams) (*types.Quote, error) {
	var result types.Quote
	body := map[string]any{
		"rfqId":    params.RfqID,
		"bidPrice": params.BidPrice,
		"askPrice": params.AskPrice,
	}
	if params.ClientID != nil {
		body["clientId"] = *params.ClientID
	}
	if err := s.client.PostAuthenticated(ctx, "api/v1/rfq/quote", body, "quoteSubmit", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AcceptQuote accepts a quote.
func (s *RFQService) AcceptQuote(ctx context.Context, quoteID string) (*types.RFQ, error) {
	var result types.RFQ
	body := map[string]any{"quoteId": quoteID}
	if err := s.client.PostAuthenticated(ctx, "api/v1/rfq/accept", body, "quoteAccept", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RefreshRFQ refreshes an RFQ to extend its expiration.
func (s *RFQService) RefreshRFQ(ctx context.Context, rfqID string) (*types.RFQ, error) {
	var result types.RFQ
	body := map[string]any{"rfqId": rfqID}
	if err := s.client.PostAuthenticated(ctx, "api/v1/rfq/refresh", body, "rfqRefresh", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelRFQ cancels an RFQ.
func (s *RFQService) CancelRFQ(ctx context.Context, rfqID string) (*types.RFQ, error) {
	var result types.RFQ
	body := map[string]any{"rfqId": rfqID}
	if err := s.client.PostAuthenticated(ctx, "api/v1/rfq/cancel", body, "rfqCancel", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

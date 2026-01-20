package services

import (
	"context"
	"strconv"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

// OrdersService provides order operations.
type OrdersService struct {
	client HTTPClient
}

// NewOrdersService creates a new OrdersService.
func NewOrdersService(client HTTPClient) *OrdersService {
	return &OrdersService{client: client}
}

// GetOrder retrieves an open order from the order book.
func (s *OrdersService) GetOrder(ctx context.Context, params types.GetOrderParams) (*types.Order, error) {
	var result types.Order
	queryParams := map[string]string{"symbol": params.Symbol}
	if params.OrderID != "" {
		queryParams["orderId"] = params.OrderID
	}
	if params.ClientID != nil {
		queryParams["clientId"] = strconv.FormatInt(*params.ClientID, 10)
	}
	if err := s.client.GetAuthenticated(ctx, "api/v1/order", queryParams, "orderQuery", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ExecuteOrder executes a new order.
func (s *OrdersService) ExecuteOrder(ctx context.Context, params types.ExecuteOrderParams) (*types.Order, error) {
	var result types.Order
	if err := s.client.PostAuthenticated(ctx, "api/v1/order", params.ToMap(), "orderExecute", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelOrder cancels an open order.
func (s *OrdersService) CancelOrder(ctx context.Context, params types.CancelOrderParams) (*types.Order, error) {
	var result types.Order
	body := map[string]any{"symbol": params.Symbol}
	if params.OrderID != "" {
		body["orderId"] = params.OrderID
	}
	if params.ClientID != nil {
		body["clientId"] = *params.ClientID
	}
	if err := s.client.DeleteAuthenticated(ctx, "api/v1/order", body, "orderCancel", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetOpenOrders retrieves all open orders.
func (s *OrdersService) GetOpenOrders(ctx context.Context, params *types.GetOpenOrdersParams) ([]types.Order, error) {
	var result []types.Order
	queryParams := make(map[string]string)
	if params != nil {
		if params.Symbol != "" {
			queryParams["symbol"] = params.Symbol
		}
		if params.MarketType != "" {
			queryParams["marketType"] = string(params.MarketType)
		}
	}
	if err := s.client.GetAuthenticated(ctx, "api/v1/orders", queryParams, "orderQueryAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CancelAllOrders cancels all open orders for a market.
func (s *OrdersService) CancelAllOrders(ctx context.Context, params types.CancelAllOrdersParams) ([]types.Order, error) {
	var result []types.Order
	body := map[string]any{"symbol": params.Symbol}
	if params.OrderType != "" {
		body["orderType"] = string(params.OrderType)
	}
	if err := s.client.DeleteAuthenticated(ctx, "api/v1/orders", body, "orderCancelAll", &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExecuteBatchOrders executes multiple orders at once.
func (s *OrdersService) ExecuteBatchOrders(ctx context.Context, orders []types.ExecuteOrderParams) ([]types.BatchOrderResult, error) {
	var result []types.BatchOrderResult

	// Convert to map format for signing
	orderMaps := make([]map[string]any, len(orders))
	for i, order := range orders {
		orderMaps[i] = order.ToMap()
	}

	if err := s.client.PostBatchOrders(ctx, "api/v1/orders", orderMaps, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// MarketOrder is a helper to create a market order.
func (s *OrdersService) MarketOrder(ctx context.Context, symbol string, side enums.Side, quantity string) (*types.Order, error) {
	return s.ExecuteOrder(ctx, types.ExecuteOrderParams{
		Symbol:    symbol,
		Side:      side,
		OrderType: enums.OrderTypeMarket,
		Quantity:  quantity,
	})
}

// LimitOrder is a helper to create a limit order.
func (s *OrdersService) LimitOrder(ctx context.Context, symbol string, side enums.Side, price, quantity string) (*types.Order, error) {
	return s.ExecuteOrder(ctx, types.ExecuteOrderParams{
		Symbol:    symbol,
		Side:      side,
		OrderType: enums.OrderTypeLimit,
		Price:     price,
		Quantity:  quantity,
	})
}

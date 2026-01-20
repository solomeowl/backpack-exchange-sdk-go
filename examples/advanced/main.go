// Example: Advanced usage patterns
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/errors"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/services"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

func main() {
	publicKey := os.Getenv("BACKPACK_API_KEY")
	secretKey := os.Getenv("BACKPACK_SECRET_KEY")

	if publicKey == "" || secretKey == "" {
		log.Fatal("Please set BACKPACK_API_KEY and BACKPACK_SECRET_KEY environment variables")
	}

	// Create client with custom options
	client, err := backpack.NewClient(
		backpack.WithCredentials(publicKey, secretKey),
		backpack.WithTimeout(60*time.Second),
		backpack.WithWindow(10000), // 10 second signature window
		backpack.WithDebug(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Example 1: Batch Orders
	fmt.Println("=== Batch Orders Example ===")
	batchOrdersExample(ctx, client)

	// Example 2: Error Handling Patterns
	fmt.Println("\n=== Error Handling Patterns ===")
	errorHandlingExample(ctx, client)

	// Example 3: Pagination
	fmt.Println("\n=== Pagination Example ===")
	paginationExample(ctx, client)

	// Example 4: Working with Positions
	fmt.Println("\n=== Positions Example ===")
	positionsExample(ctx, client)

	// Example 5: Borrow/Lend
	fmt.Println("\n=== Borrow/Lend Example ===")
	borrowLendExample(ctx, client)

	fmt.Println("\nDone!")
}

func batchOrdersExample(ctx context.Context, client *backpack.Client) {
	// Prepare batch orders
	orders := []types.ExecuteOrderParams{
		{
			Symbol:    "SOL_USDC",
			Side:      enums.SideBid,
			OrderType: enums.OrderTypeLimit,
			Price:     "50.00",
			Quantity:  "0.1",
		},
		{
			Symbol:    "SOL_USDC",
			Side:      enums.SideBid,
			OrderType: enums.OrderTypeLimit,
			Price:     "49.00",
			Quantity:  "0.1",
		},
		{
			Symbol:    "SOL_USDC",
			Side:      enums.SideBid,
			OrderType: enums.OrderTypeLimit,
			Price:     "48.00",
			Quantity:  "0.1",
		},
	}

	fmt.Println("Would submit batch orders:")
	for i, order := range orders {
		fmt.Printf("  %d. %s %s %s @ %s\n", i+1, order.Side, order.OrderType, order.Quantity, order.Price)
	}

	// Uncomment to actually execute:
	/*
		results, err := client.Orders.ExecuteBatchOrders(ctx, orders)
		if err != nil {
			log.Printf("Batch order error: %v", err)
			return
		}
		for i, result := range results {
			if result.Error != "" {
				fmt.Printf("  Order %d failed: %s\n", i+1, result.Error)
			} else {
				fmt.Printf("  Order %d: ID=%s, Status=%s\n", i+1, result.Order.ID, result.Order.Status)
			}
		}
	*/
}

func errorHandlingExample(ctx context.Context, client *backpack.Client) {
	// Example: Handling specific error codes
	_, err := client.Orders.GetOrder(ctx, types.GetOrderParams{
		Symbol:  "SOL_USDC",
		OrderID: "invalid-order-id",
	})

	if err != nil {
		// Check if it's an API error
		if apiErr, ok := errors.IsAPIError(err); ok {
			fmt.Printf("API Error: Status=%d, Code=%s, Message=%s\n",
				apiErr.StatusCode, apiErr.Code, apiErr.Message)

			// Handle specific error codes
			switch {
			case apiErr.HasCode(errors.ErrCodeInsufficientFunds):
				fmt.Println("Action: Deposit more funds")
			case apiErr.HasCode(errors.ErrCodeOrderNotFound):
				fmt.Println("Action: Order may have been filled or cancelled")
			case apiErr.HasCode(errors.ErrCodeRateLimited):
				fmt.Println("Action: Wait before retrying")
			default:
				fmt.Printf("Unhandled error code: %s\n", apiErr.Code)
			}
		} else if reqErr, ok := errors.IsRequestError(err); ok {
			// Network or request-level error
			fmt.Printf("Request Error: %s %s - %s\n", reqErr.Method, reqErr.URL, reqErr.Message)
		} else {
			// Unknown error
			fmt.Printf("Unknown error: %v\n", err)
		}
	}
}

func paginationExample(ctx context.Context, client *backpack.Client) {
	// Fetch trade history with pagination
	params := &types.HistoryParams{
		Symbol: "SOL_USDC",
		Limit:  10,
		Offset: 0,
	}

	fmt.Printf("Fetching first %d fills...\n", params.Limit)
	fills, err := client.History.GetFillHistory(ctx, params)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Got %d fills\n", len(fills))
	for i, fill := range fills {
		fmt.Printf("  %d. %s %s @ %s (Fee: %s %s)\n",
			i+1, fill.Side, fill.Quantity, fill.Price, fill.Fee, fill.FeeSymbol)
	}

	// Get next page
	if len(fills) == params.Limit {
		params.Offset = params.Limit
		fmt.Printf("\nFetching next page (offset=%d)...\n", params.Offset)
		nextFills, err := client.History.GetFillHistory(ctx, params)
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}
		fmt.Printf("Got %d more fills\n", len(nextFills))
	}
}

func positionsExample(ctx context.Context, client *backpack.Client) {
	// Get all positions
	positions, err := client.Positions.GetPositions(ctx, nil)
	if err != nil {
		log.Printf("Error getting positions: %v", err)
		return
	}

	if len(positions) == 0 {
		fmt.Println("No open positions")
		return
	}

	for _, pos := range positions {
		fmt.Printf("Position: %s\n", pos.Symbol)
		fmt.Printf("  Net Quantity: %s\n", pos.NetQuantity)
		fmt.Printf("  Entry Price: %s\n", pos.EntryPrice)
		fmt.Printf("  Mark Price: %s\n", pos.MarkPrice)
		fmt.Printf("  Unrealized PnL: %s\n", pos.PnlUnrealized)
		fmt.Printf("  Est. Liquidation Price: %s\n", pos.EstLiquidationPrice)
	}
}

func borrowLendExample(ctx context.Context, client *backpack.Client) {
	// Get borrow/lend markets (public)
	markets, err := client.BorrowLendMarkets.GetMarkets(ctx)
	if err != nil {
		log.Printf("Error getting borrow/lend markets: %v", err)
		return
	}

	fmt.Printf("Found %d borrow/lend markets\n", len(markets))
	for i, market := range markets {
		if i >= 5 {
			fmt.Println("...")
			break
		}
		fmt.Printf("  %s: Borrow Rate=%s, Lend Rate=%s\n",
			market.Symbol, market.BorrowInterestRate, market.LendInterestRate)
	}

	// Get borrow/lend positions (authenticated)
	positions, err := client.BorrowLend.GetPositions(ctx)
	if err != nil {
		log.Printf("Error getting borrow/lend positions: %v", err)
		return
	}

	if len(positions) == 0 {
		fmt.Println("\nNo borrow/lend positions")
	} else {
		fmt.Println("\nBorrow/Lend Positions:")
		for _, pos := range positions {
			fmt.Printf("  %s: Net Qty=%s, Mark Price=%s (Interest: %s)\n",
				pos.Symbol, pos.NetQuantity, pos.MarkPrice, pos.CumulativeInterest)
		}
	}

	// Get historical market data
	history, err := client.BorrowLendMarkets.GetMarketHistory(ctx, services.GetMarketHistoryParams{
		Interval: enums.BorrowLendMarketHistoryInterval1d,
		Symbol:   "USDC",
	})
	if err != nil {
		log.Printf("Error getting market history: %v", err)
		return
	}

	fmt.Println("\nUSDC Borrow/Lend Rate History:")
	for _, h := range history {
		fmt.Printf("  Borrow: %s, Lend: %s\n", h.BorrowInterestRate, h.LendInterestRate)
	}
}

// Example: Authenticated API usage
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/errors"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

func main() {
	// Get credentials from environment variables
	publicKey := os.Getenv("BACKPACK_API_KEY")
	secretKey := os.Getenv("BACKPACK_SECRET_KEY")

	if publicKey == "" || secretKey == "" {
		log.Fatal("Please set BACKPACK_API_KEY and BACKPACK_SECRET_KEY environment variables")
	}

	// Create an authenticated client
	client, err := backpack.NewClient(
		backpack.WithCredentials(publicKey, secretKey),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Get account information
	fmt.Println("=== Account Info ===")
	account, err := client.Account.GetAccount(ctx)
	if err != nil {
		log.Printf("Error getting account: %v", err)
	} else {
		fmt.Printf("Total Equity: %s\n", account.TotalEquity)
		fmt.Printf("Available Collateral: %s\n", account.AvailableCollateral)
	}

	// Get balances
	fmt.Println("\n=== Balances ===")
	balances, err := client.Capital.GetBalances(ctx)
	if err != nil {
		log.Printf("Error getting balances: %v", err)
	} else {
		for symbol, balance := range balances {
			if balance.Available != "0" || balance.Locked != "0" {
				fmt.Printf("  %s: Available=%s, Locked=%s\n",
					symbol, balance.Available, balance.Locked)
			}
		}
	}

	// Get open orders
	fmt.Println("\n=== Open Orders ===")
	openOrders, err := client.Orders.GetOpenOrders(ctx, nil)
	if err != nil {
		log.Printf("Error getting open orders: %v", err)
	} else {
		if len(openOrders) == 0 {
			fmt.Println("  No open orders")
		} else {
			for _, order := range openOrders {
				fmt.Printf("  %s: %s %s @ %s (%s)\n",
					order.Symbol, order.Side, order.Quantity, order.Price, order.Status)
			}
		}
	}

	// Example: Place a limit order (commented out for safety)
	fmt.Println("\n=== Place Order Example (dry run) ===")
	fmt.Println("Would execute: Buy 0.1 SOL @ $50.00")

	// Uncomment to actually place an order:
	/*
		order, err := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
			Symbol:    "SOL_USDC",
			Side:      enums.SideBid,
			OrderType: enums.OrderTypeLimit,
			Price:     "50.00",
			Quantity:  "0.1",
		})
		if err != nil {
			if apiErr, ok := errors.IsAPIError(err); ok {
				if apiErr.HasCode(errors.ErrCodeInsufficientFunds) {
					fmt.Println("Insufficient funds!")
				} else {
					fmt.Printf("API Error: %s\n", apiErr.Message)
				}
			} else {
				log.Printf("Error placing order: %v", err)
			}
		} else {
			fmt.Printf("Order placed: ID=%s, Status=%s\n", order.ID, order.Status)
		}
	*/

	// Using helper functions
	fmt.Println("\n=== Helper Functions Example ===")
	fmt.Println("client.Orders.MarketOrder(ctx, \"SOL_USDC\", enums.SideBid, \"1.0\")")
	fmt.Println("client.Orders.LimitOrder(ctx, \"SOL_USDC\", enums.SideAsk, \"200.00\", \"1.0\")")

	// Get order history
	fmt.Println("\n=== Order History ===")
	history, err := client.History.GetOrderHistory(ctx, &types.HistoryParams{
		Limit: 5,
	})
	if err != nil {
		log.Printf("Error getting order history: %v", err)
	} else {
		if len(history) == 0 {
			fmt.Println("  No order history")
		} else {
			for _, order := range history {
				fmt.Printf("  %s: %s %s %s @ %s (%s)\n",
					order.Symbol, order.Side, order.OrderType, order.Quantity, order.Price, order.Status)
			}
		}
	}

	// Error handling example
	fmt.Println("\n=== Error Handling Example ===")
	_, err = client.Orders.GetOrder(ctx, types.GetOrderParams{
		Symbol:  "SOL_USDC",
		OrderID: "non-existent-order-id",
	})
	if err != nil {
		if apiErr, ok := errors.IsAPIError(err); ok {
			fmt.Printf("API Error Code: %s\n", apiErr.Code)
			fmt.Printf("API Error Message: %s\n", apiErr.Message)
		} else {
			fmt.Printf("Other error: %v\n", err)
		}
	}

	fmt.Println("\nDone!")

	// Prevent unused import errors
	_ = enums.SideBid
	_ = types.ExecuteOrderParams{}
	_ = errors.ErrCodeInsufficientFunds
}

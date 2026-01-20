// Package backpack provides a Go SDK for Backpack Exchange API.
//
// The SDK provides complete coverage of all Backpack Exchange API endpoints including:
//   - Public APIs: System, Assets, Markets, Trades, BorrowLend Markets, Prediction Markets
//   - Authenticated APIs: Account, Capital, Orders, Positions, BorrowLend, RFQ, Strategy, History
//   - WebSocket streaming for real-time data
//
// Basic Usage:
//
//	// Create a public client (no authentication)
//	client, err := backpack.NewClient()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Get all markets
//	markets, err := client.Markets.GetMarkets(context.Background(), nil)
//
//	// Create an authenticated client
//	client, err := backpack.NewClient(
//	    backpack.WithCredentials(publicKey, secretKey),
//	)
//
//	// Execute an order
//	order, err := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
//	    Symbol:    "SOL_USDC",
//	    Side:      enums.SideBid,
//	    OrderType: enums.OrderTypeLimit,
//	    Price:     "100.00",
//	    Quantity:  "1.0",
//	})
//
// For more examples, see the examples directory.
package backpack

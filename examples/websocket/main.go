// Example: WebSocket usage
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/websocket"
)

func main() {
	// Create WebSocket client
	wsClient, err := websocket.NewClient(
		websocket.WithAutoReconnect(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Connect
	ctx := context.Background()
	if err := wsClient.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer wsClient.Close()

	fmt.Println("Connected to WebSocket")

	// Create type-safe handler
	handler := websocket.NewHandler(wsClient)

	// Subscribe to trade updates
	fmt.Println("Subscribing to SOL_USDC trades...")
	err = handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
		side := "Buy"
		if trade.BuyerMaker {
			side = "Sell"
		}
		fmt.Printf("[Trade] %s: %s @ %s\n", side, trade.Quantity, trade.Price)
	})
	if err != nil {
		log.Printf("Error subscribing to trades: %v", err)
	}

	// Subscribe to ticker updates
	fmt.Println("Subscribing to SOL_USDC ticker...")
	err = handler.OnTicker("SOL_USDC", func(ticker *types.WSTicker) {
		fmt.Printf("[Ticker] Price: %s, Change: %s%%\n",
			ticker.LastPrice, ticker.PriceChangePercent)
	})
	if err != nil {
		log.Printf("Error subscribing to ticker: %v", err)
	}

	// Subscribe to order book updates
	fmt.Println("Subscribing to SOL_USDC depth...")
	err = handler.OnDepth("SOL_USDC", func(depth *types.WSDepth) {
		bestBid := "N/A"
		bestAsk := "N/A"
		if len(depth.Bids) > 0 && len(depth.Bids[0]) >= 1 {
			bestBid = depth.Bids[0][0]
		}
		if len(depth.Asks) > 0 && len(depth.Asks[0]) >= 1 {
			bestAsk = depth.Asks[0][0]
		}
		fmt.Printf("[Depth] Best Bid: %s, Best Ask: %s\n", bestBid, bestAsk)
	})
	if err != nil {
		log.Printf("Error subscribing to depth: %v", err)
	}

	// Example: Private streams (requires authentication)
	publicKey := os.Getenv("BACKPACK_API_KEY")
	secretKey := os.Getenv("BACKPACK_SECRET_KEY")

	if publicKey != "" && secretKey != "" {
		fmt.Println("\nSubscribing to private streams...")

		// Create authenticated WebSocket client
		authWsClient, err := websocket.NewClient(
			websocket.WithCredentials(publicKey, secretKey),
			websocket.WithAutoReconnect(true),
		)
		if err != nil {
			log.Printf("Error creating authenticated WS client: %v", err)
		} else {
			if err := authWsClient.Connect(ctx); err != nil {
				log.Printf("Error connecting authenticated WS: %v", err)
			} else {
				defer authWsClient.Close()

				authHandler := websocket.NewHandler(authWsClient)

				// Subscribe to order updates
				err = authHandler.OnOrderUpdate(func(order *types.WSOrderUpdate) {
					fmt.Printf("[Order Update] %s: %s %s @ %s (%s)\n",
						order.Symbol, order.Side, order.Quantity, order.Price, order.Status)
				})
				if err != nil {
					log.Printf("Error subscribing to order updates: %v", err)
				}

				// Subscribe to position updates
				err = authHandler.OnPositionUpdate(func(pos *types.WSPositionUpdate) {
					fmt.Printf("[Position Update] %s: %s @ %s, PnL: %s\n",
						pos.Symbol, pos.Quantity, pos.EntryPrice, pos.UnrealizedPnl)
				})
				if err != nil {
					log.Printf("Error subscribing to position updates: %v", err)
				}

				fmt.Println("Subscribed to private streams")
			}
		}
	} else {
		fmt.Println("\nSkipping private streams (no credentials)")
	}

	// Wait for interrupt signal
	fmt.Println("\nPress Ctrl+C to exit...")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nShutting down...")
}

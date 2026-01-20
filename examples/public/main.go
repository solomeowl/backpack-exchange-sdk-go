// Example: Public API usage (no authentication required)
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/solomeowl/backpack-exchange-sdk-go/backpack"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
	"github.com/solomeowl/backpack-exchange-sdk-go/backpack/services"
)

func main() {
	// Create a client without credentials (public APIs only)
	client, err := backpack.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Get system status
	fmt.Println("=== System Status ===")
	status, err := client.System.GetStatus(ctx)
	if err != nil {
		log.Printf("Error getting status: %v", err)
	} else {
		fmt.Printf("Status: %s\n", status.Status)
	}

	// Get server time
	fmt.Println("\n=== Server Time ===")
	serverTime, err := client.System.GetTime(ctx)
	if err != nil {
		log.Printf("Error getting time: %v", err)
	} else {
		fmt.Printf("Server time: %d\n", serverTime.ServerTime)
	}

	// Get all markets
	fmt.Println("\n=== Markets ===")
	markets, err := client.Markets.GetMarkets(ctx)
	if err != nil {
		log.Printf("Error getting markets: %v", err)
	} else {
		fmt.Printf("Found %d markets\n", len(markets))
		for i, market := range markets {
			if i >= 5 {
				fmt.Println("...")
				break
			}
			fmt.Printf("  %s (%s)\n", market.Symbol, market.MarketType)
		}
	}

	// Get ticker for a specific market
	fmt.Println("\n=== Ticker (SOL_USDC) ===")
	ticker, err := client.Markets.GetTicker(ctx, "SOL_USDC")
	if err != nil {
		log.Printf("Error getting ticker: %v", err)
	} else {
		fmt.Printf("Last Price: %s\n", ticker.LastPrice)
		fmt.Printf("24h Change: %s%%\n", ticker.PriceChangePercent)
		fmt.Printf("24h Volume: %s\n", ticker.Volume)
	}

	// Get order book depth
	fmt.Println("\n=== Order Book (SOL_USDC) ===")
	depth, err := client.Markets.GetDepth(ctx, services.GetDepthParams{
		Symbol: "SOL_USDC",
		Limit:  enums.DepthLimit5,
	})
	if err != nil {
		log.Printf("Error getting depth: %v", err)
	} else {
		fmt.Println("Bids:")
		for _, bid := range depth.Bids {
			if len(bid) >= 2 {
				fmt.Printf("  %s @ %s\n", bid[1], bid[0])
			}
		}
		fmt.Println("Asks:")
		for _, ask := range depth.Asks {
			if len(ask) >= 2 {
				fmt.Printf("  %s @ %s\n", ask[1], ask[0])
			}
		}
	}

	// Get recent trades
	fmt.Println("\n=== Recent Trades (SOL_USDC) ===")
	trades, err := client.Trades.GetRecentTrades(ctx, services.GetRecentTradesParams{
		Symbol: "SOL_USDC",
		Limit:  5,
	})
	if err != nil {
		log.Printf("Error getting trades: %v", err)
	} else {
		for _, trade := range trades {
			side := "Buy"
			if trade.IsBuyerMaker {
				side = "Sell"
			}
			fmt.Printf("  %s: %s @ %s\n", side, trade.Quantity, trade.Price)
		}
	}

	// Get klines
	fmt.Println("\n=== Klines (SOL_USDC, 1h) ===")
	klines, err := client.Markets.GetKlines(ctx, services.GetKlinesParams{
		Symbol:   "SOL_USDC",
		Interval: enums.KlineInterval1h,
		Limit:    3,
	})
	if err != nil {
		log.Printf("Error getting klines: %v", err)
	} else {
		for _, kline := range klines {
			fmt.Printf("  O: %s, H: %s, L: %s, C: %s\n",
				kline.Open, kline.High, kline.Low, kline.Close)
		}
	}

	fmt.Println("\nDone!")
}

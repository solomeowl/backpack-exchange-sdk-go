# Backpack Exchange Go SDK

![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

**English** | [简体中文](./README_zh-Hans.md) | [繁體中文](./README_zh-Hant.md) | [日本語](./README_ja.md) | [한국어](./README_ko.md) | [Español](./README_es.md) | [Português](./README_pt.md)

A complete Go SDK for the [Backpack Exchange](https://backpack.exchange/) API. Supports all 70 API endpoints including REST and WebSocket.

## Features

- **Complete Coverage**: All 70 API endpoints implemented (59 unique paths)
- **ED25519 Authentication**: Secure signing for authenticated endpoints
- **Type Safety**: Strongly typed request/response structs
- **WebSocket Client**: Real-time streaming for market data and account updates
- **Enums**: Comprehensive enums for type-safe API calls
- **Functional Options**: Flexible client configuration

## Installation

```bash
go get github.com/solomeowl/backpack-exchange-sdk-go
```

## Quick Start

### Public Client

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/solomeowl/backpack-exchange-sdk-go/backpack"
    "github.com/solomeowl/backpack-exchange-sdk-go/backpack/services"
)

func main() {
    client, err := backpack.NewClient()
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Get all markets
    markets, _ := client.Markets.GetMarkets(ctx, nil)

    // Get ticker
    ticker, _ := client.Markets.GetTicker(ctx, services.GetTickerParams{Symbol: "SOL_USDC"})

    // Get order book depth
    depth, _ := client.Markets.GetDepth(ctx, services.GetDepthParams{Symbol: "SOL_USDC"})

    fmt.Printf("Markets: %d, Ticker: %s, Depth bids: %d\n",
        len(markets), ticker.LastPrice, len(depth.Bids))
}
```

### Authentication Client

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/solomeowl/backpack-exchange-sdk-go/backpack"
    "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
    "github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

func main() {
    client, err := backpack.NewClient(
        backpack.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Get account balances
    balances, _ := client.Capital.GetBalances(ctx)

    // Place an order
    order, _ := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
        Symbol:    "SOL_USDC",
        Side:      enums.SideBid,
        OrderType: enums.OrderTypeLimit,
        Price:     "100",
        Quantity:  "1",
    })

    // Get order history
    history, _ := client.History.GetOrderHistory(ctx, &types.OrderHistoryParams{
        Symbol: "SOL_USDC",
    })

    fmt.Printf("Balances: %d, Order: %s, History: %d\n",
        len(balances), order.ID, len(history))
}
```

### Using Enums

```go
package main

import (
    "context"

    "github.com/solomeowl/backpack-exchange-sdk-go/backpack"
    "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"
    "github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
)

func main() {
    client, _ := backpack.NewClient(
        backpack.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    )

    timeInForce := enums.TimeInForceGTC
    order, _ := client.Orders.ExecuteOrder(context.Background(), types.ExecuteOrderParams{
        Symbol:      "SOL_USDC",
        Side:        enums.SideBid,
        OrderType:   enums.OrderTypeLimit,
        Price:       "100",
        Quantity:    "1",
        TimeInForce: &timeInForce,
    })

    _ = order
}
```

## API Reference

### Public Client Methods

| Category | Method | Description |
|----------|--------|-------------|
| **System** | `System.GetStatus(ctx)` | Get system status |
| | `System.Ping(ctx)` | Ping the server |
| | `System.GetTime(ctx)` | Get server time |
| | `System.GetWallets(ctx)` | Get supported wallets |
| **Assets** | `Assets.GetAssets(ctx)` | Get all assets |
| | `Assets.GetCollateral(ctx)` | Get collateral info |
| **Markets** | `Markets.GetMarkets(ctx, params)` | Get all markets |
| | `Markets.GetMarket(ctx, symbol)` | Get specific market |
| | `Markets.GetTicker(ctx, params)` | Get ticker |
| | `Markets.GetTickers(ctx, params)` | Get all tickers |
| | `Markets.GetDepth(ctx, params)` | Get order book |
| | `Markets.GetKlines(ctx, params)` | Get candlesticks |
| | `Markets.GetMarkPrice(ctx, symbol)` | Get mark price |
| | `Markets.GetMarkPrices(ctx)` | Get all mark prices |
| | `Markets.GetOpenInterest(ctx, symbol)` | Get open interest |
| | `Markets.GetFundingRates(ctx, symbol)` | Get funding rates |
| **Trades** | `Trades.GetRecentTrades(ctx, symbol)` | Get recent trades |
| | `Trades.GetHistoricalTrades(ctx, params)` | Get trade history |
| **Borrow/Lend** | `BorrowLendMarkets.GetMarkets(ctx)` | Get lending markets |
| | `BorrowLendMarkets.GetMarketHistory(ctx, params)` | Get lending history |
| **Prediction** | `Prediction.GetMarkets(ctx)` | Get prediction markets |
| | `Prediction.GetTags(ctx, nil)` | Get prediction tags |

### Authentication Client Methods

| Category | Method | Description |
|----------|--------|-------------|
| **Account** | `Account.GetAccount(ctx)` | Get account settings |
| | `Account.UpdateAccount(ctx, settings)` | Update account settings |
| | `Account.GetMaxBorrowQuantity(ctx, symbol)` | Get max borrow amount |
| | `Account.GetMaxOrderQuantity(ctx, params)` | Get max order size |
| | `Account.GetMaxWithdrawalQuantity(ctx, params)` | Get max withdrawal |
| | `Account.ConvertDust(ctx, &symbol)` | Convert dust to USDC |
| **Capital** | `Capital.GetBalances(ctx)` | Get balances |
| | `Capital.GetCollateral(ctx, nil)` | Get collateral |
| | `Capital.GetDeposits(ctx, params)` | Get deposit history |
| | `Capital.GetDepositAddress(ctx, blockchain)` | Get deposit address |
| | `Capital.GetWithdrawals(ctx, params)` | Get withdrawal history |
| | `Capital.RequestWithdrawal(ctx, req)` | Request withdrawal |
| | `Capital.GetWithdrawalDelay(ctx)` | Get withdrawal delay |
| | `Capital.CreateWithdrawalDelay(ctx, params)` | Create withdrawal delay |
| | `Capital.UpdateWithdrawalDelay(ctx, params)` | Update withdrawal delay |
| **Orders** | `Orders.ExecuteOrder(ctx, params)` | Place single order |
| | `Orders.ExecuteBatchOrders(ctx, orders)` | Place batch orders |
| | `Orders.GetOpenOrders(ctx, params)` | Get open orders |
| | `Orders.GetOrder(ctx, params)` | Get specific order |
| | `Orders.CancelOrder(ctx, params)` | Cancel single order |
| | `Orders.CancelAllOrders(ctx, symbol)` | Cancel all orders |
| **History** | `History.GetOrderHistory(ctx, params)` | Get order history |
| | `History.GetFillHistory(ctx, params)` | Get fill history |
| | `History.GetBorrowHistory(ctx, params)` | Get borrow history |
| | `History.GetInterestHistory(ctx, params)` | Get interest history |
| | `History.GetBorrowPositionHistory(ctx, params)` | Get borrow positions |
| | `History.GetFundingPayments(ctx, params)` | Get funding payments |
| | `History.GetSettlementHistory(ctx, params)` | Get settlements |
| | `History.GetDustHistory(ctx, params)` | Get dust conversions |
| | `History.GetPositionHistory(ctx, params)` | Get position history |
| | `History.GetStrategyHistory(ctx, params)` | Get strategy history |
| | `History.GetRFQHistory(ctx, params)` | Get RFQ history |
| | `History.GetQuoteHistory(ctx, params)` | Get quote history |
| | `History.GetRFQFillHistory(ctx, params)` | Get RFQ fills |
| | `History.GetQuoteFillHistory(ctx, params)` | Get quote fills |
| **Borrow/Lend** | `BorrowLend.GetPositions(ctx)` | Get positions |
| | `BorrowLend.Execute(ctx, params)` | Borrow or lend |
| | `BorrowLend.GetLiquidationPrice(ctx, params)` | Get liquidation price |
| **Positions** | `Positions.GetPositions(ctx, params)` | Get open positions |
| **RFQ** | `RFQ.SubmitRFQ(ctx, params)` | Submit RFQ |
| | `RFQ.SubmitQuote(ctx, params)` | Submit quote |
| | `RFQ.AcceptQuote(ctx, types.QuoteAcceptParams{QuoteID: quoteId})` | Accept quote |
| | `RFQ.RefreshRFQ(ctx, rfqId)` | Refresh RFQ |
| | `RFQ.CancelRFQ(ctx, types.RFQCancelParams{RfqID: rfqId})` | Cancel RFQ |
| **Strategy** | `Strategy.CreateStrategy(ctx, params)` | Create strategy |
| | `Strategy.GetStrategy(ctx, strategyId)` | Get strategy |
| | `Strategy.GetOpenStrategies(ctx, symbol)` | Get open strategies |
| | `Strategy.CancelStrategy(ctx, strategyId)` | Cancel strategy |
| | `Strategy.CancelAllStrategies(ctx, symbol)` | Cancel all strategies |

### WebSocket Client

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/solomeowl/backpack-exchange-sdk-go/backpack/types"
    "github.com/solomeowl/backpack-exchange-sdk-go/backpack/websocket"
)

func main() {
    // Public streams (no auth required)
    wsClient := websocket.NewClient()

    // Private streams (auth required)
    // wsClient := websocket.NewClient(
    //     websocket.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    // )

    if err := wsClient.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()

    handler := websocket.NewHandler(wsClient)

    // Subscribe to book ticker
    handler.OnBookTicker("SOL_USDC", func(ticker *types.WSBookTicker) {
        fmt.Printf("Best bid: %s, Best ask: %s\n", ticker.BidPrice, ticker.AskPrice)
    })

    // Subscribe to trades
    handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
        fmt.Printf("Trade: %s @ %s\n", trade.Quantity, trade.Price)
    })

    // Private stream example
    // handler.OnOrderUpdate("", func(order *types.WSOrderUpdate) {
    //     fmt.Printf("Order update: %s %s\n", order.OrderID, order.Status)
    // })

    // Keep running
    select {}
}
```

## Available Enums

```go
import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Order related
enums.OrderTypeLimit                  // Limit, Market
enums.SideBid                         // Bid, Ask
enums.TimeInForceGTC                  // GTC, IOC, FOK
enums.SelfTradePreventionRejectTaker  // RejectTaker, RejectMaker, RejectBoth
enums.OrderStatusNew                  // New, Filled, Cancelled, Expired, PartiallyFilled, TriggerPending, TriggerFailed

// Market related
enums.MarketTypeSpot       // SPOT, PERP, IPERP, DATED, PREDICTION, RFQ
enums.KlineInterval1m      // 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1month
enums.DepthLimit5          // "5", "10", "20", "50", "100", "500", "1000"
enums.OrderBookStateOpen   // Open, Closed, CancelOnly, LimitOnly, PostOnly

// Status related
enums.DepositStatusPending            // cancelled, confirmed, declined, expired, initiated, etc.
enums.WithdrawalStatusPending         // confirmed, ownershipVerificationRequired, pending, etc.

// Blockchain (30+ supported)
enums.BlockchainSolana
enums.BlockchainEthereum
enums.BlockchainBitcoin
enums.BlockchainArbitrum
enums.BlockchainBase
// ... and more
```

## Configuration Options

```go
client, err := backpack.NewClient(
    backpack.WithCredentials(publicKey, secretKey),  // API credentials
    backpack.WithBaseURL("https://api.backpack.exchange"),  // Custom base URL
    backpack.WithTimeout(30 * time.Second),  // HTTP timeout
    backpack.WithWindow(5000),  // Signature window in ms
    backpack.WithDebug(true),  // Enable debug logging
    backpack.WithHTTPClient(customClient),  // Custom HTTP client
)
```

## Examples

See the [examples](./examples) directory for complete usage examples:

- `public/main.go` - Public API examples
- `authenticated/main.go` - Authenticated API examples
- `websocket/main.go` - WebSocket streaming examples
- `advanced/main.go` - Advanced usage examples

## Documentation

For detailed API documentation, visit the [Backpack Exchange API Docs](https://docs.backpack.exchange/).

## Changelog

### v1.1.0
- Sync all types and enums with OpenAPI specification
- Update enum values to match API (SortDirection, OrderStatus, MarketType, etc.)
- Expand Blockchain enum from 4 to 30+ networks
- Add new enums: PositionState, OrderBookState, PaymentType, RfqExecutionMode, etc.
- Fix type definitions for Order, Position, Strategy, RFQ, Capital
- Update service method signatures

### v1.0.0
- Initial release with full API support
- All 70 API endpoints implemented
- WebSocket client with auto-reconnect
- ED25519 authentication
- Comprehensive enums and types

## Support

If this SDK has been helpful, please consider:

1. Using my referral link to register: [Register on Backpack Exchange](https://backpack.exchange/refer/solomeowl)
2. Giving this repo a star on GitHub

## License

MIT License

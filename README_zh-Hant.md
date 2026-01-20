# Backpack Exchange Go SDK

![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

[English](./README.md) | [简体中文](./README_zh-Hans.md) | **繁體中文** | [日本語](./README_ja.md) | [한국어](./README_ko.md) | [Español](./README_es.md) | [Português](./README_pt.md)

完整的 [Backpack Exchange](https://backpack.exchange/) API Go SDK。支援所有 70 個 API 端點，包含 REST 和 WebSocket。

## 功能特色

- **完整覆蓋**：實作所有 70 個 API 端點（59 個唯一路徑）
- **ED25519 認證**：安全的簽名機制用於認證端點
- **型別安全**：強型別的請求/回應結構
- **WebSocket 客戶端**：即時串流市場資料和帳戶更新
- **列舉型別**：完整的列舉定義確保 API 呼叫的型別安全
- **函數式選項**：彈性的客戶端配置

## 安裝

```bash
go get github.com/solomeowl/backpack-exchange-sdk-go
```

## 快速開始

### 公開客戶端

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/solomeowl/backpack-exchange-sdk-go/backpack"
)

func main() {
    client, err := backpack.NewClient()
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // 取得所有市場
    markets, _ := client.Markets.GetMarkets(ctx)

    // 取得行情
    ticker, _ := client.Markets.GetTicker(ctx, "SOL_USDC")

    // 取得訂單簿深度
    depth, _ := client.Markets.GetDepth(ctx, "SOL_USDC")

    fmt.Printf("市場數: %d, 最新價: %s, 買單深度: %d\n",
        len(markets), ticker.LastPrice, len(depth.Bids))
}
```

### 認證客戶端

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

    // 取得帳戶餘額
    balances, _ := client.Capital.GetBalances(ctx)

    // 下單
    order, _ := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
        Symbol:    "SOL_USDC",
        Side:      enums.SideBid,
        OrderType: enums.OrderTypeLimit,
        Price:     "100",
        Quantity:  "1",
    })

    // 取得訂單歷史
    history, _ := client.History.GetOrderHistory(ctx, &types.HistoryParams{
        Symbol: "SOL_USDC",
    })

    fmt.Printf("餘額數: %d, 訂單: %s, 歷史數: %d\n",
        len(balances), order.ID, len(history))
}
```

### 使用列舉

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

## API 參考

### 公開客戶端方法

| 分類 | 方法 | 說明 |
|------|------|------|
| **系統** | `System.GetStatus(ctx)` | 取得系統狀態 |
| | `System.Ping(ctx)` | Ping 伺服器 |
| | `System.GetTime(ctx)` | 取得伺服器時間 |
| | `System.GetWallets(ctx)` | 取得支援的錢包 |
| **資產** | `Assets.GetAssets(ctx)` | 取得所有資產 |
| | `Assets.GetCollateral(ctx)` | 取得抵押品資訊 |
| **市場** | `Markets.GetMarkets(ctx)` | 取得所有市場 |
| | `Markets.GetMarket(ctx, symbol)` | 取得特定市場 |
| | `Markets.GetTicker(ctx, symbol)` | 取得行情 |
| | `Markets.GetTickers(ctx)` | 取得所有行情 |
| | `Markets.GetDepth(ctx, symbol)` | 取得訂單簿 |
| | `Markets.GetKlines(ctx, params)` | 取得 K 線 |
| | `Markets.GetMarkPrice(ctx, symbol)` | 取得標記價格 |
| | `Markets.GetMarkPrices(ctx)` | 取得所有標記價格 |
| | `Markets.GetOpenInterest(ctx, symbol)` | 取得未平倉量 |
| | `Markets.GetFundingRates(ctx, symbol)` | 取得資金費率 |
| **交易** | `Trades.GetRecentTrades(ctx, symbol)` | 取得最近成交 |
| | `Trades.GetHistoricalTrades(ctx, params)` | 取得歷史成交 |
| **借貸** | `BorrowLendMarkets.GetMarkets(ctx)` | 取得借貸市場 |
| | `BorrowLendMarkets.GetMarketHistory(ctx, params)` | 取得借貸歷史 |
| **預測** | `Prediction.GetMarkets(ctx)` | 取得預測市場 |
| | `Prediction.GetTags(ctx)` | 取得預測標籤 |

### 認證客戶端方法

| 分類 | 方法 | 說明 |
|------|------|------|
| **帳戶** | `Account.GetAccount(ctx)` | 取得帳戶設定 |
| | `Account.UpdateAccount(ctx, settings)` | 更新帳戶設定 |
| | `Account.GetMaxBorrowQuantity(ctx, symbol)` | 取得最大借入數量 |
| | `Account.GetMaxOrderQuantity(ctx, params)` | 取得最大下單數量 |
| | `Account.GetMaxWithdrawalQuantity(ctx, params)` | 取得最大提現數量 |
| | `Account.ConvertDust(ctx, symbol)` | 將塵埃轉換為 USDC |
| **資金** | `Capital.GetBalances(ctx)` | 取得餘額 |
| | `Capital.GetCollateral(ctx)` | 取得抵押品 |
| | `Capital.GetDeposits(ctx, params)` | 取得充值歷史 |
| | `Capital.GetDepositAddress(ctx, blockchain)` | 取得充值地址 |
| | `Capital.GetWithdrawals(ctx, params)` | 取得提現歷史 |
| | `Capital.RequestWithdrawal(ctx, req)` | 請求提現 |
| | `Capital.GetWithdrawalDelay(ctx)` | 取得提現延遲設定 |
| | `Capital.CreateWithdrawalDelay(ctx, params)` | 建立提現延遲 |
| | `Capital.UpdateWithdrawalDelay(ctx, params)` | 更新提現延遲 |
| **訂單** | `Orders.ExecuteOrder(ctx, params)` | 下單 |
| | `Orders.ExecuteBatchOrders(ctx, orders)` | 批量下單 |
| | `Orders.GetOpenOrders(ctx, params)` | 取得未成交訂單 |
| | `Orders.GetOrder(ctx, params)` | 取得特定訂單 |
| | `Orders.CancelOrder(ctx, params)` | 取消訂單 |
| | `Orders.CancelAllOrders(ctx, symbol)` | 取消所有訂單 |
| **歷史** | `History.GetOrderHistory(ctx, params)` | 取得訂單歷史 |
| | `History.GetFillHistory(ctx, params)` | 取得成交歷史 |
| | `History.GetBorrowHistory(ctx, params)` | 取得借入歷史 |
| | `History.GetInterestHistory(ctx, params)` | 取得利息歷史 |
| | `History.GetBorrowPositionHistory(ctx, params)` | 取得借貸倉位歷史 |
| | `History.GetFundingPayments(ctx, params)` | 取得資金費用歷史 |
| | `History.GetSettlementHistory(ctx, params)` | 取得結算歷史 |
| | `History.GetDustHistory(ctx, params)` | 取得塵埃轉換歷史 |
| | `History.GetPositionHistory(ctx, params)` | 取得倉位歷史 |
| | `History.GetStrategyHistory(ctx, params)` | 取得策略歷史 |
| | `History.GetRFQHistory(ctx, params)` | 取得 RFQ 歷史 |
| | `History.GetQuoteHistory(ctx, params)` | 取得報價歷史 |
| | `History.GetRFQFillHistory(ctx, params)` | 取得 RFQ 成交歷史 |
| | `History.GetQuoteFillHistory(ctx, params)` | 取得報價成交歷史 |
| **借貸** | `BorrowLend.GetPositions(ctx)` | 取得倉位 |
| | `BorrowLend.Execute(ctx, params)` | 借入或借出 |
| | `BorrowLend.GetLiquidationPrice(ctx, symbol)` | 取得清算價格 |
| **倉位** | `Positions.GetPositions(ctx)` | 取得未平倉位 |
| | `Positions.GetPosition(ctx, symbol)` | 取得特定倉位 |
| **RFQ** | `RFQ.SubmitRFQ(ctx, params)` | 提交 RFQ |
| | `RFQ.SubmitQuote(ctx, params)` | 提交報價 |
| | `RFQ.AcceptQuote(ctx, params)` | 接受報價 |
| | `RFQ.RefreshRFQ(ctx, rfqId)` | 刷新 RFQ |
| | `RFQ.CancelRFQ(ctx, rfqId)` | 取消 RFQ |
| **策略** | `Strategy.CreateStrategy(ctx, params)` | 建立策略 |
| | `Strategy.GetStrategy(ctx, symbol, strategyId)` | 取得策略 |
| | `Strategy.GetOpenStrategies(ctx, symbol)` | 取得未完成策略 |
| | `Strategy.CancelStrategy(ctx, symbol, strategyId)` | 取消策略 |
| | `Strategy.CancelAllStrategies(ctx, symbol)` | 取消所有策略 |

### WebSocket 客戶端

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
    // 公開串流（無需認證）
    wsClient := websocket.NewClient()

    // 私有串流（需要認證）
    // wsClient := websocket.NewClient(
    //     websocket.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    // )

    if err := wsClient.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()

    handler := websocket.NewHandler(wsClient)

    // 訂閱最佳買賣價
    handler.OnBookTicker("SOL_USDC", func(ticker *types.WSBookTicker) {
        fmt.Printf("最佳買價: %s, 最佳賣價: %s\n", ticker.BidPrice, ticker.AskPrice)
    })

    // 訂閱成交
    handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
        fmt.Printf("成交: %s @ %s\n", trade.Quantity, trade.Price)
    })

    // 私有串流範例
    // handler.OnOrderUpdate(func(order *types.WSOrderUpdate) {
    //     fmt.Printf("訂單更新: %s %s\n", order.OrderID, order.Status)
    // })

    // 保持運行
    select {}
}
```

## 可用列舉

```go
import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// 訂單相關
enums.OrderTypeLimit       // Limit, Market
enums.SideBid              // Bid, Ask
enums.TimeInForceGTC       // GTC, IOC, FOK
enums.SelfTradePreventionAllow

// 市場相關
enums.MarketTypeSpot       // Spot, Perp
enums.KlineInterval1m      // 1m, 5m, 15m, 1h, 4h, 1d, 1w

// 狀態相關
enums.OrderStatusNew       // New, Filled, Cancelled, etc.
enums.DepositStatusPending
enums.WithdrawalStatusPending

// 區塊鏈
enums.BlockchainSolana
enums.BlockchainEthereum
enums.BlockchainPolygon
enums.BlockchainBitcoin
```

## 配置選項

```go
client, err := backpack.NewClient(
    backpack.WithCredentials(publicKey, secretKey),  // API 憑證
    backpack.WithBaseURL("https://api.backpack.exchange"),  // 自訂 Base URL
    backpack.WithTimeout(30 * time.Second),  // HTTP 超時
    backpack.WithWindow(5000),  // 簽名有效期（毫秒）
    backpack.WithDebug(true),  // 啟用除錯日誌
    backpack.WithHTTPClient(customClient),  // 自訂 HTTP client
)
```

## 範例

請參閱 [examples](./examples) 目錄以取得完整使用範例：

- `public/main.go` - 公開 API 範例
- `authenticated/main.go` - 認證 API 範例
- `websocket/main.go` - WebSocket 串流範例
- `advanced/main.go` - 進階用法範例

## 文件

詳細 API 文件請參閱 [Backpack Exchange API Docs](https://docs.backpack.exchange/)。

## 更新日誌

### v1.0.0
- 初始版本，完整 API 支援
- 實作所有 70 個 API 端點
- WebSocket 客戶端含自動重連
- ED25519 認證
- 完整的列舉和型別定義

## 支持

如果這個 SDK 對您有幫助，請考慮：

1. 使用我的推薦連結註冊：[註冊 Backpack Exchange](https://backpack.exchange/refer/solomeowl)
2. 在 GitHub 上給這個專案一顆星

## 授權

MIT License

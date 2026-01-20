# Backpack Exchange Go SDK

![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

[English](./README.md) | **简体中文** | [繁體中文](./README_zh-Hant.md) | [日本語](./README_ja.md) | [한국어](./README_ko.md) | [Español](./README_es.md) | [Português](./README_pt.md)

完整的 [Backpack Exchange](https://backpack.exchange/) API Go SDK。支持所有 70 个 API 端点，包含 REST 和 WebSocket。

## 功能特色

- **完整覆盖**：实现所有 70 个 API 端点（59 个唯一路径）
- **ED25519 认证**：安全的签名机制用于认证端点
- **类型安全**：强类型的请求/响应结构
- **WebSocket 客户端**：实时流式市场数据和账户更新
- **枚举类型**：完整的枚举定义确保 API 调用的类型安全
- **函数式选项**：灵活的客户端配置

## 安装

```bash
go get github.com/solomeowl/backpack-exchange-sdk-go
```

## 快速开始

### 公开客户端

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

    // 获取所有市场
    markets, _ := client.Markets.GetMarkets(ctx, nil)

    // 获取行情
    ticker, _ := client.Markets.GetTicker(ctx, services.GetTickerParams{Symbol: "SOL_USDC"})

    // 获取订单簿深度
    depth, _ := client.Markets.GetDepth(ctx, services.GetDepthParams{Symbol: "SOL_USDC"})

    fmt.Printf("市场数: %d, 最新价: %s, 买单深度: %d\n",
        len(markets), ticker.LastPrice, len(depth.Bids))
}
```

### 认证客户端

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

    // 获取账户余额
    balances, _ := client.Capital.GetBalances(ctx)

    // 下单
    order, _ := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
        Symbol:    "SOL_USDC",
        Side:      enums.SideBid,
        OrderType: enums.OrderTypeLimit,
        Price:     "100",
        Quantity:  "1",
    })

    // 获取订单历史
    history, _ := client.History.GetOrderHistory(ctx, &types.OrderHistoryParams{
        Symbol: "SOL_USDC",
    })

    fmt.Printf("余额数: %d, 订单: %s, 历史数: %d\n",
        len(balances), order.ID, len(history))
}
```

### 使用枚举

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

## API 参考

### 公开客户端方法

| 分类 | 方法 | 说明 |
|------|------|------|
| **系统** | `System.GetStatus(ctx)` | 获取系统状态 |
| | `System.Ping(ctx)` | Ping 服务器 |
| | `System.GetTime(ctx)` | 获取服务器时间 |
| | `System.GetWallets(ctx)` | 获取支持的钱包 |
| **资产** | `Assets.GetAssets(ctx)` | 获取所有资产 |
| | `Assets.GetCollateral(ctx)` | 获取抵押品信息 |
| **市场** | `Markets.GetMarkets(ctx, params)` | 获取所有市场 |
| | `Markets.GetMarket(ctx, symbol)` | 获取特定市场 |
| | `Markets.GetTicker(ctx, params)` | 获取行情 |
| | `Markets.GetTickers(ctx, params)` | 获取所有行情 |
| | `Markets.GetDepth(ctx, params)` | 获取订单簿 |
| | `Markets.GetKlines(ctx, params)` | 获取 K 线 |
| | `Markets.GetMarkPrice(ctx, symbol)` | 获取标记价格 |
| | `Markets.GetMarkPrices(ctx)` | 获取所有标记价格 |
| | `Markets.GetOpenInterest(ctx, symbol)` | 获取未平仓量 |
| | `Markets.GetFundingRates(ctx, symbol)` | 获取资金费率 |
| **交易** | `Trades.GetRecentTrades(ctx, symbol)` | 获取最近成交 |
| | `Trades.GetHistoricalTrades(ctx, params)` | 获取历史成交 |
| **借贷** | `BorrowLendMarkets.GetMarkets(ctx)` | 获取借贷市场 |
| | `BorrowLendMarkets.GetMarketHistory(ctx, params)` | 获取借贷历史 |
| **预测** | `Prediction.GetMarkets(ctx)` | 获取预测市场 |
| | `Prediction.GetTags(ctx, nil)` | 获取预测标签 |

### 认证客户端方法

| 分类 | 方法 | 说明 |
|------|------|------|
| **账户** | `Account.GetAccount(ctx)` | 获取账户设置 |
| | `Account.UpdateAccount(ctx, settings)` | 更新账户设置 |
| | `Account.GetMaxBorrowQuantity(ctx, symbol)` | 获取最大借入数量 |
| | `Account.GetMaxOrderQuantity(ctx, params)` | 获取最大下单数量 |
| | `Account.GetMaxWithdrawalQuantity(ctx, params)` | 获取最大提现数量 |
| | `Account.ConvertDust(ctx, &symbol)` | 将尘埃转换为 USDC |
| **资金** | `Capital.GetBalances(ctx)` | 获取余额 |
| | `Capital.GetCollateral(ctx, nil)` | 获取抵押品 |
| | `Capital.GetDeposits(ctx, params)` | 获取充值历史 |
| | `Capital.GetDepositAddress(ctx, blockchain)` | 获取充值地址 |
| | `Capital.GetWithdrawals(ctx, params)` | 获取提现历史 |
| | `Capital.RequestWithdrawal(ctx, req)` | 请求提现 |
| | `Capital.GetWithdrawalDelay(ctx)` | 获取提现延迟设置 |
| | `Capital.CreateWithdrawalDelay(ctx, params)` | 创建提现延迟 |
| | `Capital.UpdateWithdrawalDelay(ctx, params)` | 更新提现延迟 |
| **订单** | `Orders.ExecuteOrder(ctx, params)` | 下单 |
| | `Orders.ExecuteBatchOrders(ctx, orders)` | 批量下单 |
| | `Orders.GetOpenOrders(ctx, params)` | 获取未成交订单 |
| | `Orders.GetOrder(ctx, params)` | 获取特定订单 |
| | `Orders.CancelOrder(ctx, params)` | 取消订单 |
| | `Orders.CancelAllOrders(ctx, symbol)` | 取消所有订单 |
| **历史** | `History.GetOrderHistory(ctx, params)` | 获取订单历史 |
| | `History.GetFillHistory(ctx, params)` | 获取成交历史 |
| | `History.GetBorrowHistory(ctx, params)` | 获取借入历史 |
| | `History.GetInterestHistory(ctx, params)` | 获取利息历史 |
| | `History.GetBorrowPositionHistory(ctx, params)` | 获取借贷仓位历史 |
| | `History.GetFundingPayments(ctx, params)` | 获取资金费用历史 |
| | `History.GetSettlementHistory(ctx, params)` | 获取结算历史 |
| | `History.GetDustHistory(ctx, params)` | 获取尘埃转换历史 |
| | `History.GetPositionHistory(ctx, params)` | 获取仓位历史 |
| | `History.GetStrategyHistory(ctx, params)` | 获取策略历史 |
| | `History.GetRFQHistory(ctx, params)` | 获取 RFQ 历史 |
| | `History.GetQuoteHistory(ctx, params)` | 获取报价历史 |
| | `History.GetRFQFillHistory(ctx, params)` | 获取 RFQ 成交历史 |
| | `History.GetQuoteFillHistory(ctx, params)` | 获取报价成交历史 |
| **借贷** | `BorrowLend.GetPositions(ctx)` | 获取仓位 |
| | `BorrowLend.Execute(ctx, params)` | 借入或借出 |
| | `BorrowLend.GetLiquidationPrice(ctx, params)` | 获取清算价格 |
| **仓位** | `Positions.GetPositions(ctx, params)` | 获取未平仓位 |
| **RFQ** | `RFQ.SubmitRFQ(ctx, params)` | 提交 RFQ |
| | `RFQ.SubmitQuote(ctx, params)` | 提交报价 |
| | `RFQ.AcceptQuote(ctx, types.QuoteAcceptParams{QuoteID: quoteId})` | 接受报价 |
| | `RFQ.RefreshRFQ(ctx, rfqId)` | 刷新 RFQ |
| | `RFQ.CancelRFQ(ctx, types.RFQCancelParams{RfqID: rfqId})` | 取消 RFQ |
| **策略** | `Strategy.CreateStrategy(ctx, params)` | 创建策略 |
| | `Strategy.GetStrategy(ctx, strategyId)` | 获取策略 |
| | `Strategy.GetOpenStrategies(ctx, symbol)` | 获取未完成策略 |
| | `Strategy.CancelStrategy(ctx, strategyId)` | 取消策略 |
| | `Strategy.CancelAllStrategies(ctx, symbol)` | 取消所有策略 |

### WebSocket 客户端

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
    // 公开流（无需认证）
    wsClient := websocket.NewClient()

    // 私有流（需要认证）
    // wsClient := websocket.NewClient(
    //     websocket.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    // )

    if err := wsClient.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()

    handler := websocket.NewHandler(wsClient)

    // 订阅最佳买卖价
    handler.OnBookTicker("SOL_USDC", func(ticker *types.WSBookTicker) {
        fmt.Printf("最佳买价: %s, 最佳卖价: %s\n", ticker.BidPrice, ticker.AskPrice)
    })

    // 订阅成交
    handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
        fmt.Printf("成交: %s @ %s\n", trade.Quantity, trade.Price)
    })

    // 保持运行
    select {}
}
```

## 可用枚举

```go
import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// 订单相关
enums.OrderTypeLimit                  // Limit, Market
enums.SideBid                         // Bid, Ask
enums.TimeInForceGTC                  // GTC, IOC, FOK
enums.SelfTradePreventionRejectTaker  // RejectTaker, RejectMaker, RejectBoth
enums.OrderStatusNew                  // New, Filled, Cancelled, Expired, PartiallyFilled, TriggerPending, TriggerFailed

// 市场相关
enums.MarketTypeSpot       // SPOT, PERP, IPERP, DATED, PREDICTION, RFQ
enums.KlineInterval1m      // 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1month
enums.DepthLimit5          // "5", "10", "20", "50", "100", "500", "1000"
enums.OrderBookStateOpen   // Open, Closed, CancelOnly, LimitOnly, PostOnly

// 状态相关
enums.DepositStatusPending            // cancelled, confirmed, declined, expired, initiated, etc.
enums.WithdrawalStatusPending         // confirmed, ownershipVerificationRequired, pending, etc.

// 区块链（支持 30+ 网络）
enums.BlockchainSolana
enums.BlockchainEthereum
enums.BlockchainBitcoin
enums.BlockchainArbitrum
enums.BlockchainBase
// ... 更多
```

## 配置选项

```go
client, err := backpack.NewClient(
    backpack.WithCredentials(publicKey, secretKey),  // API 凭证
    backpack.WithBaseURL("https://api.backpack.exchange"),  // 自定义 Base URL
    backpack.WithTimeout(30 * time.Second),  // HTTP 超时
    backpack.WithWindow(5000),  // 签名有效期（毫秒）
    backpack.WithDebug(true),  // 启用调试日志
    backpack.WithHTTPClient(customClient),  // 自定义 HTTP client
)
```

## 示例

请参阅 [examples](./examples) 目录以获取完整使用示例：

- `public/main.go` - 公开 API 示例
- `authenticated/main.go` - 认证 API 示例
- `websocket/main.go` - WebSocket 流式示例
- `advanced/main.go` - 高级用法示例

## 文档

详细 API 文档请参阅 [Backpack Exchange API Docs](https://docs.backpack.exchange/)。

## 更新日志

### v1.1.0
- 同步所有类型和枚举至 OpenAPI 规格
- 更新枚举值以符合 API（SortDirection、OrderStatus、MarketType 等）
- 扩展区块链枚举从 4 个到 30+ 个网络
- 新增枚举：PositionState、OrderBookState、PaymentType、RfqExecutionMode 等
- 修正类型定义：Order、Position、Strategy、RFQ、Capital
- 更新服务方法签名

### v1.0.0
- 初始版本，完整 API 支持
- 实现所有 70 个 API 端点
- WebSocket 客户端含自动重连
- ED25519 认证
- 完整的枚举和类型定义

## 支持

如果这个 SDK 对您有帮助，请考虑：

1. 使用我的推荐链接注册：[注册 Backpack Exchange](https://backpack.exchange/refer/solomeowl)
2. 在 GitHub 上给这个项目一颗星

## 许可证

MIT License

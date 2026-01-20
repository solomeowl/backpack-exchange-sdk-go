# Backpack Exchange Go SDK

![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

[English](./README.md) | [简体中文](./README_zh-Hans.md) | [繁體中文](./README_zh-Hant.md) | **日本語** | [한국어](./README_ko.md) | [Español](./README_es.md) | [Português](./README_pt.md)

[Backpack Exchange](https://backpack.exchange/) API の完全な Go SDK。REST と WebSocket を含む全 70 の API エンドポイントをサポート。

## 特徴

- **完全カバレッジ**：全 70 API エンドポイント実装（59 ユニークパス）
- **ED25519 認証**：認証エンドポイント用のセキュアな署名
- **型安全**：強い型付けのリクエスト/レスポンス構造体
- **WebSocket クライアント**：マーケットデータとアカウント更新のリアルタイムストリーミング
- **列挙型**：型安全な API 呼び出しのための包括的な列挙型
- **関数オプション**：柔軟なクライアント設定

## インストール

```bash
go get github.com/solomeowl/backpack-exchange-sdk-go
```

## クイックスタート

### パブリッククライアント

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

    // 全マーケット取得
    markets, _ := client.Markets.GetMarkets(ctx, nil)

    // ティッカー取得
    ticker, _ := client.Markets.GetTicker(ctx, services.GetTickerParams{Symbol: "SOL_USDC"})

    // オーダーブック深度取得
    depth, _ := client.Markets.GetDepth(ctx, services.GetDepthParams{Symbol: "SOL_USDC"})

    fmt.Printf("マーケット数: %d, 最終価格: %s, 買い板深度: %d\n",
        len(markets), ticker.LastPrice, len(depth.Bids))
}
```

### 認証クライアント

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

    // 残高取得
    balances, _ := client.Capital.GetBalances(ctx)

    // 注文実行
    order, _ := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
        Symbol:    "SOL_USDC",
        Side:      enums.SideBid,
        OrderType: enums.OrderTypeLimit,
        Price:     "100",
        Quantity:  "1",
    })

    // 注文履歴取得
    history, _ := client.History.GetOrderHistory(ctx, &types.OrderHistoryParams{
        Symbol: "SOL_USDC",
    })

    fmt.Printf("残高数: %d, 注文: %s, 履歴数: %d\n",
        len(balances), order.ID, len(history))
}
```

### 列挙型の使用

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

## API リファレンス

### パブリッククライアントメソッド

| カテゴリ | メソッド | 説明 |
|----------|--------|------|
| **システム** | `System.GetStatus(ctx)` | システムステータス取得 |
| | `System.Ping(ctx)` | サーバーへ Ping |
| | `System.GetTime(ctx)` | サーバー時間取得 |
| | `System.GetWallets(ctx)` | 対応ウォレット取得 |
| **資産** | `Assets.GetAssets(ctx)` | 全資産取得 |
| | `Assets.GetCollateral(ctx)` | 担保情報取得 |
| **マーケット** | `Markets.GetMarkets(ctx, params)` | 全マーケット取得 |
| | `Markets.GetMarket(ctx, symbol)` | 特定マーケット取得 |
| | `Markets.GetTicker(ctx, params)` | ティッカー取得 |
| | `Markets.GetTickers(ctx, params)` | 全ティッカー取得 |
| | `Markets.GetDepth(ctx, params)` | オーダーブック取得 |
| | `Markets.GetKlines(ctx, params)` | ローソク足取得 |
| | `Markets.GetMarkPrice(ctx, symbol)` | マーク価格取得 |
| | `Markets.GetMarkPrices(ctx)` | 全マーク価格取得 |
| | `Markets.GetOpenInterest(ctx, symbol)` | 建玉取得 |
| | `Markets.GetFundingRates(ctx, symbol)` | ファンディングレート取得 |
| **取引** | `Trades.GetRecentTrades(ctx, symbol)` | 最近の取引取得 |
| | `Trades.GetHistoricalTrades(ctx, params)` | 取引履歴取得 |
| **借入/貸出** | `BorrowLendMarkets.GetMarkets(ctx)` | 貸借マーケット取得 |
| | `BorrowLendMarkets.GetMarketHistory(ctx, params)` | 貸借履歴取得 |
| **予測** | `Prediction.GetMarkets(ctx)` | 予測マーケット取得 |
| | `Prediction.GetTags(ctx, nil)` | 予測タグ取得 |

### 認証クライアントメソッド

| カテゴリ | メソッド | 説明 |
|----------|--------|------|
| **アカウント** | `Account.GetAccount(ctx)` | アカウント設定取得 |
| | `Account.UpdateAccount(ctx, settings)` | アカウント設定更新 |
| | `Account.GetMaxBorrowQuantity(ctx, symbol)` | 最大借入数量取得 |
| | `Account.GetMaxOrderQuantity(ctx, params)` | 最大注文数量取得 |
| | `Account.GetMaxWithdrawalQuantity(ctx, params)` | 最大出金数量取得 |
| | `Account.ConvertDust(ctx, &symbol)` | ダストを USDC に変換 |
| **資金** | `Capital.GetBalances(ctx)` | 残高取得 |
| | `Capital.GetCollateral(ctx, nil)` | 担保取得 |
| | `Capital.GetDeposits(ctx, params)` | 入金履歴取得 |
| | `Capital.GetDepositAddress(ctx, blockchain)` | 入金アドレス取得 |
| | `Capital.GetWithdrawals(ctx, params)` | 出金履歴取得 |
| | `Capital.RequestWithdrawal(ctx, req)` | 出金リクエスト |
| | `Capital.GetWithdrawalDelay(ctx)` | 出金遅延設定取得 |
| | `Capital.CreateWithdrawalDelay(ctx, params)` | 出金遅延作成 |
| | `Capital.UpdateWithdrawalDelay(ctx, params)` | 出金遅延更新 |
| **注文** | `Orders.ExecuteOrder(ctx, params)` | 注文実行 |
| | `Orders.ExecuteBatchOrders(ctx, orders)` | バッチ注文実行 |
| | `Orders.GetOpenOrders(ctx, params)` | オープン注文取得 |
| | `Orders.GetOrder(ctx, params)` | 特定注文取得 |
| | `Orders.CancelOrder(ctx, params)` | 注文キャンセル |
| | `Orders.CancelAllOrders(ctx, symbol)` | 全注文キャンセル |
| **履歴** | `History.GetOrderHistory(ctx, params)` | 注文履歴取得 |
| | `History.GetFillHistory(ctx, params)` | 約定履歴取得 |
| | `History.GetBorrowHistory(ctx, params)` | 借入履歴取得 |
| | `History.GetInterestHistory(ctx, params)` | 金利履歴取得 |
| | `History.GetBorrowPositionHistory(ctx, params)` | 借入ポジション履歴取得 |
| | `History.GetFundingPayments(ctx, params)` | ファンディング支払履歴取得 |
| | `History.GetSettlementHistory(ctx, params)` | 決済履歴取得 |
| | `History.GetDustHistory(ctx, params)` | ダスト変換履歴取得 |
| | `History.GetPositionHistory(ctx, params)` | ポジション履歴取得 |
| | `History.GetStrategyHistory(ctx, params)` | ストラテジー履歴取得 |
| | `History.GetRFQHistory(ctx, params)` | RFQ 履歴取得 |
| | `History.GetQuoteHistory(ctx, params)` | クォート履歴取得 |
| | `History.GetRFQFillHistory(ctx, params)` | RFQ 約定履歴取得 |
| | `History.GetQuoteFillHistory(ctx, params)` | クォート約定履歴取得 |
| **借入/貸出** | `BorrowLend.GetPositions(ctx)` | ポジション取得 |
| | `BorrowLend.Execute(ctx, params)` | 借入または貸出実行 |
| | `BorrowLend.GetLiquidationPrice(ctx, params)` | 清算価格取得 |
| **ポジション** | `Positions.GetPositions(ctx, params)` | オープンポジション取得 |
| **RFQ** | `RFQ.SubmitRFQ(ctx, params)` | RFQ 提出 |
| | `RFQ.SubmitQuote(ctx, params)` | クォート提出 |
| | `RFQ.AcceptQuote(ctx, types.QuoteAcceptParams{QuoteID: quoteId})` | クォート承認 |
| | `RFQ.RefreshRFQ(ctx, rfqId)` | RFQ 更新 |
| | `RFQ.CancelRFQ(ctx, types.RFQCancelParams{RfqID: rfqId})` | RFQ キャンセル |
| **ストラテジー** | `Strategy.CreateStrategy(ctx, params)` | ストラテジー作成 |
| | `Strategy.GetStrategy(ctx, strategyId)` | ストラテジー取得 |
| | `Strategy.GetOpenStrategies(ctx, symbol)` | オープンストラテジー取得 |
| | `Strategy.CancelStrategy(ctx, strategyId)` | ストラテジーキャンセル |
| | `Strategy.CancelAllStrategies(ctx, symbol)` | 全ストラテジーキャンセル |

### WebSocket クライアント

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
    // パブリックストリーム（認証不要）
    wsClient := websocket.NewClient()

    // プライベートストリーム（認証必要）
    // wsClient := websocket.NewClient(
    //     websocket.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    // )

    if err := wsClient.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()

    handler := websocket.NewHandler(wsClient)

    // ブックティッカー購読
    handler.OnBookTicker("SOL_USDC", func(ticker *types.WSBookTicker) {
        fmt.Printf("最良買気配: %s, 最良売気配: %s\n", ticker.BidPrice, ticker.AskPrice)
    })

    // 取引購読
    handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
        fmt.Printf("取引: %s @ %s\n", trade.Quantity, trade.Price)
    })

    // 実行継続
    select {}
}
```

## 利用可能な列挙型

```go
import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// 注文関連
enums.OrderTypeLimit                  // Limit, Market
enums.SideBid                         // Bid, Ask
enums.TimeInForceGTC                  // GTC, IOC, FOK
enums.SelfTradePreventionRejectTaker  // RejectTaker, RejectMaker, RejectBoth
enums.OrderStatusNew                  // New, Filled, Cancelled, Expired, PartiallyFilled, TriggerPending, TriggerFailed

// マーケット関連
enums.MarketTypeSpot       // SPOT, PERP, IPERP, DATED, PREDICTION, RFQ
enums.KlineInterval1m      // 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1month
enums.DepthLimit5          // "5", "10", "20", "50", "100", "500", "1000"
enums.OrderBookStateOpen   // Open, Closed, CancelOnly, LimitOnly, PostOnly

// ステータス関連
enums.DepositStatusPending            // cancelled, confirmed, declined, expired, initiated, etc.
enums.WithdrawalStatusPending         // confirmed, ownershipVerificationRequired, pending, etc.

// ブロックチェーン（30+ ネットワーク対応）
enums.BlockchainSolana
enums.BlockchainEthereum
enums.BlockchainBitcoin
enums.BlockchainArbitrum
enums.BlockchainBase
// ... その他
```

## 設定オプション

```go
client, err := backpack.NewClient(
    backpack.WithCredentials(publicKey, secretKey),  // API 認証情報
    backpack.WithBaseURL("https://api.backpack.exchange"),  // カスタム Base URL
    backpack.WithTimeout(30 * time.Second),  // HTTP タイムアウト
    backpack.WithWindow(5000),  // 署名有効期間（ミリ秒）
    backpack.WithDebug(true),  // デバッグログ有効化
    backpack.WithHTTPClient(customClient),  // カスタム HTTP クライアント
)
```

## サンプル

完全な使用例については [examples](./examples) ディレクトリを参照：

- `public/main.go` - パブリック API サンプル
- `authenticated/main.go` - 認証 API サンプル
- `websocket/main.go` - WebSocket ストリーミングサンプル
- `advanced/main.go` - 高度な使用法サンプル

## ドキュメント

詳細な API ドキュメントは [Backpack Exchange API Docs](https://docs.backpack.exchange/) を参照。

## 変更履歴

### v1.1.0
- OpenAPI 仕様に全型と列挙型を同期
- 列挙値を API に合わせて更新（SortDirection、OrderStatus、MarketType 等）
- Blockchain 列挙型を 4 から 30+ ネットワークに拡張
- 新規列挙型追加：PositionState、OrderBookState、PaymentType、RfqExecutionMode 等
- 型定義修正：Order、Position、Strategy、RFQ、Capital
- サービスメソッドシグネチャ更新

### v1.0.0
- 初回リリース、完全 API サポート
- 全 70 API エンドポイント実装
- 自動再接続付き WebSocket クライアント
- ED25519 認証
- 包括的な列挙型と型定義

## サポート

この SDK が役立った場合は、以下をご検討ください：

1. 紹介リンクで登録：[Backpack Exchange に登録](https://backpack.exchange/refer/solomeowl)
2. GitHub でスターを付ける

## ライセンス

MIT License

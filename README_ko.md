# Backpack Exchange Go SDK

![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

[English](./README.md) | [简体中文](./README_zh-Hans.md) | [繁體中文](./README_zh-Hant.md) | [日本語](./README_ja.md) | **한국어** | [Español](./README_es.md) | [Português](./README_pt.md)

[Backpack Exchange](https://backpack.exchange/) API를 위한 완전한 Go SDK. REST 및 WebSocket을 포함한 모든 70개 API 엔드포인트 지원.

## 특징

- **완전한 커버리지**: 모든 70개 API 엔드포인트 구현 (59개 고유 경로)
- **ED25519 인증**: 인증된 엔드포인트를 위한 보안 서명
- **타입 안전성**: 강력한 타입의 요청/응답 구조체
- **WebSocket 클라이언트**: 시장 데이터 및 계정 업데이트의 실시간 스트리밍
- **열거형**: 타입 안전한 API 호출을 위한 포괄적인 열거형
- **함수형 옵션**: 유연한 클라이언트 구성

## 설치

```bash
go get github.com/solomeowl/backpack-exchange-sdk-go
```

## 빠른 시작

### 공개 클라이언트

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

    // 모든 마켓 조회
    markets, _ := client.Markets.GetMarkets(ctx)

    // 티커 조회
    ticker, _ := client.Markets.GetTicker(ctx, "SOL_USDC")

    // 오더북 깊이 조회
    depth, _ := client.Markets.GetDepth(ctx, "SOL_USDC")

    fmt.Printf("마켓 수: %d, 최근가: %s, 매수 깊이: %d\n",
        len(markets), ticker.LastPrice, len(depth.Bids))
}
```

### 인증 클라이언트

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

    // 계정 잔고 조회
    balances, _ := client.Capital.GetBalances(ctx)

    // 주문 실행
    order, _ := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
        Symbol:    "SOL_USDC",
        Side:      enums.SideBid,
        OrderType: enums.OrderTypeLimit,
        Price:     "100",
        Quantity:  "1",
    })

    // 주문 내역 조회
    history, _ := client.History.GetOrderHistory(ctx, &types.HistoryParams{
        Symbol: "SOL_USDC",
    })

    fmt.Printf("잔고 수: %d, 주문: %s, 내역 수: %d\n",
        len(balances), order.ID, len(history))
}
```

### 열거형 사용

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

## API 레퍼런스

### 공개 클라이언트 메서드

| 카테고리 | 메서드 | 설명 |
|----------|--------|------|
| **시스템** | `System.GetStatus(ctx)` | 시스템 상태 조회 |
| | `System.Ping(ctx)` | 서버 Ping |
| | `System.GetTime(ctx)` | 서버 시간 조회 |
| | `System.GetWallets(ctx)` | 지원 지갑 조회 |
| **자산** | `Assets.GetAssets(ctx)` | 모든 자산 조회 |
| | `Assets.GetCollateral(ctx)` | 담보 정보 조회 |
| **마켓** | `Markets.GetMarkets(ctx)` | 모든 마켓 조회 |
| | `Markets.GetMarket(ctx, symbol)` | 특정 마켓 조회 |
| | `Markets.GetTicker(ctx, symbol)` | 티커 조회 |
| | `Markets.GetTickers(ctx)` | 모든 티커 조회 |
| | `Markets.GetDepth(ctx, symbol)` | 오더북 조회 |
| | `Markets.GetKlines(ctx, params)` | 캔들스틱 조회 |
| | `Markets.GetMarkPrice(ctx, symbol)` | 마크 가격 조회 |
| | `Markets.GetMarkPrices(ctx)` | 모든 마크 가격 조회 |
| | `Markets.GetOpenInterest(ctx, symbol)` | 미결제약정 조회 |
| | `Markets.GetFundingRates(ctx, symbol)` | 펀딩비율 조회 |
| **거래** | `Trades.GetRecentTrades(ctx, symbol)` | 최근 거래 조회 |
| | `Trades.GetHistoricalTrades(ctx, params)` | 거래 내역 조회 |
| **대출** | `BorrowLendMarkets.GetMarkets(ctx)` | 대출 마켓 조회 |
| | `BorrowLendMarkets.GetMarketHistory(ctx, params)` | 대출 내역 조회 |
| **예측** | `Prediction.GetMarkets(ctx)` | 예측 마켓 조회 |
| | `Prediction.GetTags(ctx)` | 예측 태그 조회 |

### 인증 클라이언트 메서드

| 카테고리 | 메서드 | 설명 |
|----------|--------|------|
| **계정** | `Account.GetAccount(ctx)` | 계정 설정 조회 |
| | `Account.UpdateAccount(ctx, settings)` | 계정 설정 업데이트 |
| | `Account.GetMaxBorrowQuantity(ctx, symbol)` | 최대 대출 수량 조회 |
| | `Account.GetMaxOrderQuantity(ctx, params)` | 최대 주문 수량 조회 |
| | `Account.GetMaxWithdrawalQuantity(ctx, params)` | 최대 출금 수량 조회 |
| | `Account.ConvertDust(ctx, symbol)` | 더스트를 USDC로 변환 |
| **자금** | `Capital.GetBalances(ctx)` | 잔고 조회 |
| | `Capital.GetCollateral(ctx)` | 담보 조회 |
| | `Capital.GetDeposits(ctx, params)` | 입금 내역 조회 |
| | `Capital.GetDepositAddress(ctx, blockchain)` | 입금 주소 조회 |
| | `Capital.GetWithdrawals(ctx, params)` | 출금 내역 조회 |
| | `Capital.RequestWithdrawal(ctx, req)` | 출금 요청 |
| | `Capital.GetWithdrawalDelay(ctx)` | 출금 지연 설정 조회 |
| | `Capital.CreateWithdrawalDelay(ctx, params)` | 출금 지연 생성 |
| | `Capital.UpdateWithdrawalDelay(ctx, params)` | 출금 지연 업데이트 |
| **주문** | `Orders.ExecuteOrder(ctx, params)` | 주문 실행 |
| | `Orders.ExecuteBatchOrders(ctx, orders)` | 배치 주문 실행 |
| | `Orders.GetOpenOrders(ctx, params)` | 미체결 주문 조회 |
| | `Orders.GetOrder(ctx, params)` | 특정 주문 조회 |
| | `Orders.CancelOrder(ctx, params)` | 주문 취소 |
| | `Orders.CancelAllOrders(ctx, symbol)` | 모든 주문 취소 |
| **내역** | `History.GetOrderHistory(ctx, params)` | 주문 내역 조회 |
| | `History.GetFillHistory(ctx, params)` | 체결 내역 조회 |
| | `History.GetBorrowHistory(ctx, params)` | 대출 내역 조회 |
| | `History.GetInterestHistory(ctx, params)` | 이자 내역 조회 |
| | `History.GetBorrowPositionHistory(ctx, params)` | 대출 포지션 내역 조회 |
| | `History.GetFundingPayments(ctx, params)` | 펀딩 지불 내역 조회 |
| | `History.GetSettlementHistory(ctx, params)` | 정산 내역 조회 |
| | `History.GetDustHistory(ctx, params)` | 더스트 변환 내역 조회 |
| | `History.GetPositionHistory(ctx, params)` | 포지션 내역 조회 |
| | `History.GetStrategyHistory(ctx, params)` | 전략 내역 조회 |
| | `History.GetRFQHistory(ctx, params)` | RFQ 내역 조회 |
| | `History.GetQuoteHistory(ctx, params)` | 견적 내역 조회 |
| | `History.GetRFQFillHistory(ctx, params)` | RFQ 체결 내역 조회 |
| | `History.GetQuoteFillHistory(ctx, params)` | 견적 체결 내역 조회 |
| **대출** | `BorrowLend.GetPositions(ctx)` | 포지션 조회 |
| | `BorrowLend.Execute(ctx, params)` | 대출 또는 대여 실행 |
| | `BorrowLend.GetLiquidationPrice(ctx, symbol)` | 청산 가격 조회 |
| **포지션** | `Positions.GetPositions(ctx)` | 오픈 포지션 조회 |
| | `Positions.GetPosition(ctx, symbol)` | 특정 포지션 조회 |
| **RFQ** | `RFQ.SubmitRFQ(ctx, params)` | RFQ 제출 |
| | `RFQ.SubmitQuote(ctx, params)` | 견적 제출 |
| | `RFQ.AcceptQuote(ctx, params)` | 견적 수락 |
| | `RFQ.RefreshRFQ(ctx, rfqId)` | RFQ 갱신 |
| | `RFQ.CancelRFQ(ctx, rfqId)` | RFQ 취소 |
| **전략** | `Strategy.CreateStrategy(ctx, params)` | 전략 생성 |
| | `Strategy.GetStrategy(ctx, symbol, strategyId)` | 전략 조회 |
| | `Strategy.GetOpenStrategies(ctx, symbol)` | 오픈 전략 조회 |
| | `Strategy.CancelStrategy(ctx, symbol, strategyId)` | 전략 취소 |
| | `Strategy.CancelAllStrategies(ctx, symbol)` | 모든 전략 취소 |

### WebSocket 클라이언트

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
    // 공개 스트림 (인증 불필요)
    wsClient := websocket.NewClient()

    // 비공개 스트림 (인증 필요)
    // wsClient := websocket.NewClient(
    //     websocket.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    // )

    if err := wsClient.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()

    handler := websocket.NewHandler(wsClient)

    // 북 티커 구독
    handler.OnBookTicker("SOL_USDC", func(ticker *types.WSBookTicker) {
        fmt.Printf("최우선 매수가: %s, 최우선 매도가: %s\n", ticker.BidPrice, ticker.AskPrice)
    })

    // 거래 구독
    handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
        fmt.Printf("거래: %s @ %s\n", trade.Quantity, trade.Price)
    })

    // 계속 실행
    select {}
}
```

## 사용 가능한 열거형

```go
import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// 주문 관련
enums.OrderTypeLimit       // Limit, Market
enums.SideBid              // Bid, Ask
enums.TimeInForceGTC       // GTC, IOC, FOK
enums.SelfTradePreventionAllow

// 마켓 관련
enums.MarketTypeSpot       // Spot, Perp
enums.KlineInterval1m      // 1m, 5m, 15m, 1h, 4h, 1d, 1w

// 상태 관련
enums.OrderStatusNew       // New, Filled, Cancelled, etc.
enums.DepositStatusPending
enums.WithdrawalStatusPending

// 블록체인
enums.BlockchainSolana
enums.BlockchainEthereum
enums.BlockchainPolygon
enums.BlockchainBitcoin
```

## 구성 옵션

```go
client, err := backpack.NewClient(
    backpack.WithCredentials(publicKey, secretKey),  // API 자격 증명
    backpack.WithBaseURL("https://api.backpack.exchange"),  // 사용자 지정 Base URL
    backpack.WithTimeout(30 * time.Second),  // HTTP 타임아웃
    backpack.WithWindow(5000),  // 서명 유효 기간 (밀리초)
    backpack.WithDebug(true),  // 디버그 로깅 활성화
    backpack.WithHTTPClient(customClient),  // 사용자 지정 HTTP 클라이언트
)
```

## 예제

완전한 사용 예제는 [examples](./examples) 디렉토리 참조:

- `public/main.go` - 공개 API 예제
- `authenticated/main.go` - 인증 API 예제
- `websocket/main.go` - WebSocket 스트리밍 예제
- `advanced/main.go` - 고급 사용법 예제

## 문서

자세한 API 문서는 [Backpack Exchange API Docs](https://docs.backpack.exchange/) 참조.

## 변경 로그

### v1.0.0
- 초기 릴리스, 전체 API 지원
- 모든 70개 API 엔드포인트 구현
- 자동 재연결 WebSocket 클라이언트
- ED25519 인증
- 포괄적인 열거형 및 타입 정의

## 지원

이 SDK가 도움이 되었다면 다음을 고려해 주세요:

1. 추천 링크로 가입: [Backpack Exchange 가입](https://backpack.exchange/refer/solomeowl)
2. GitHub에서 스타 주기

## 라이선스

MIT License

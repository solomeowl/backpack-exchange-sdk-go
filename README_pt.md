# Backpack Exchange Go SDK

![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

[English](./README.md) | [简体中文](./README_zh-Hans.md) | [繁體中文](./README_zh-Hant.md) | [日本語](./README_ja.md) | [한국어](./README_ko.md) | [Español](./README_es.md) | **Português**

SDK completo em Go para a API do [Backpack Exchange](https://backpack.exchange/). Suporta todos os 70 endpoints de API incluindo REST e WebSocket.

## Recursos

- **Cobertura Completa**: Implementação de todos os 70 endpoints de API (59 caminhos únicos)
- **Autenticação ED25519**: Assinatura segura para endpoints autenticados
- **Segurança de Tipos**: Estruturas de requisição/resposta fortemente tipadas
- **Cliente WebSocket**: Streaming em tempo real de dados de mercado e atualizações de conta
- **Enumerações**: Enumerações abrangentes para chamadas de API com segurança de tipos
- **Opções Funcionais**: Configuração flexível do cliente

## Instalação

```bash
go get github.com/solomeowl/backpack-exchange-sdk-go
```

## Início Rápido

### Cliente Público

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

    // Obter todos os mercados
    markets, _ := client.Markets.GetMarkets(ctx, nil)

    // Obter ticker
    ticker, _ := client.Markets.GetTicker(ctx, services.GetTickerParams{Symbol: "SOL_USDC"})

    // Obter profundidade do livro de ordens
    depth, _ := client.Markets.GetDepth(ctx, services.GetDepthParams{Symbol: "SOL_USDC"})

    fmt.Printf("Mercados: %d, Último preço: %s, Profundidade de compra: %d\n",
        len(markets), ticker.LastPrice, len(depth.Bids))
}
```

### Cliente Autenticado

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

    // Obter saldos da conta
    balances, _ := client.Capital.GetBalances(ctx)

    // Executar ordem
    order, _ := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
        Symbol:    "SOL_USDC",
        Side:      enums.SideBid,
        OrderType: enums.OrderTypeLimit,
        Price:     "100",
        Quantity:  "1",
    })

    // Obter histórico de ordens
    history, _ := client.History.GetOrderHistory(ctx, &types.OrderHistoryParams{
        Symbol: "SOL_USDC",
    })

    fmt.Printf("Saldos: %d, Ordem: %s, Histórico: %d\n",
        len(balances), order.ID, len(history))
}
```

### Usando Enumerações

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

## Referência da API

### Métodos do Cliente Público

| Categoria | Método | Descrição |
|-----------|--------|-----------|
| **Sistema** | `System.GetStatus(ctx)` | Obter status do sistema |
| | `System.Ping(ctx)` | Ping ao servidor |
| | `System.GetTime(ctx)` | Obter hora do servidor |
| | `System.GetWallets(ctx)` | Obter carteiras suportadas |
| **Ativos** | `Assets.GetAssets(ctx)` | Obter todos os ativos |
| | `Assets.GetCollateral(ctx)` | Obter informações de colateral |
| **Mercados** | `Markets.GetMarkets(ctx, params)` | Obter todos os mercados |
| | `Markets.GetMarket(ctx, symbol)` | Obter mercado específico |
| | `Markets.GetTicker(ctx, params)` | Obter ticker |
| | `Markets.GetTickers(ctx, params)` | Obter todos os tickers |
| | `Markets.GetDepth(ctx, params)` | Obter livro de ordens |
| | `Markets.GetKlines(ctx, params)` | Obter candlesticks |
| | `Markets.GetMarkPrice(ctx, symbol)` | Obter preço de marca |
| | `Markets.GetMarkPrices(ctx)` | Obter todos os preços de marca |
| | `Markets.GetOpenInterest(ctx, symbol)` | Obter interesse aberto |
| | `Markets.GetFundingRates(ctx, symbol)` | Obter taxas de financiamento |
| **Trades** | `Trades.GetRecentTrades(ctx, symbol)` | Obter trades recentes |
| | `Trades.GetHistoricalTrades(ctx, params)` | Obter histórico de trades |
| **Empréstimo** | `BorrowLendMarkets.GetMarkets(ctx)` | Obter mercados de empréstimo |
| | `BorrowLendMarkets.GetMarketHistory(ctx, params)` | Obter histórico de empréstimo |
| **Previsão** | `Prediction.GetMarkets(ctx)` | Obter mercados de previsão |
| | `Prediction.GetTags(ctx, nil)` | Obter tags de previsão |

### Métodos do Cliente Autenticado

| Categoria | Método | Descrição |
|-----------|--------|-----------|
| **Conta** | `Account.GetAccount(ctx)` | Obter configurações da conta |
| | `Account.UpdateAccount(ctx, settings)` | Atualizar configurações |
| | `Account.GetMaxBorrowQuantity(ctx, symbol)` | Obter quantidade máxima de empréstimo |
| | `Account.GetMaxOrderQuantity(ctx, params)` | Obter tamanho máximo de ordem |
| | `Account.GetMaxWithdrawalQuantity(ctx, params)` | Obter saque máximo |
| | `Account.ConvertDust(ctx, &symbol)` | Converter poeira para USDC |
| **Capital** | `Capital.GetBalances(ctx)` | Obter saldos |
| | `Capital.GetCollateral(ctx, nil)` | Obter colateral |
| | `Capital.GetDeposits(ctx, params)` | Obter histórico de depósitos |
| | `Capital.GetDepositAddress(ctx, blockchain)` | Obter endereço de depósito |
| | `Capital.GetWithdrawals(ctx, params)` | Obter histórico de saques |
| | `Capital.RequestWithdrawal(ctx, req)` | Solicitar saque |
| | `Capital.GetWithdrawalDelay(ctx)` | Obter atraso de saque |
| | `Capital.CreateWithdrawalDelay(ctx, params)` | Criar atraso de saque |
| | `Capital.UpdateWithdrawalDelay(ctx, params)` | Atualizar atraso de saque |
| **Ordens** | `Orders.ExecuteOrder(ctx, params)` | Executar ordem |
| | `Orders.ExecuteBatchOrders(ctx, orders)` | Executar ordens em lote |
| | `Orders.GetOpenOrders(ctx, params)` | Obter ordens abertas |
| | `Orders.GetOrder(ctx, params)` | Obter ordem específica |
| | `Orders.CancelOrder(ctx, params)` | Cancelar ordem |
| | `Orders.CancelAllOrders(ctx, symbol)` | Cancelar todas as ordens |
| **Histórico** | `History.GetOrderHistory(ctx, params)` | Obter histórico de ordens |
| | `History.GetFillHistory(ctx, params)` | Obter histórico de fills |
| | `History.GetBorrowHistory(ctx, params)` | Obter histórico de empréstimos |
| | `History.GetInterestHistory(ctx, params)` | Obter histórico de juros |
| | `History.GetBorrowPositionHistory(ctx, params)` | Obter histórico de posições |
| | `History.GetFundingPayments(ctx, params)` | Obter pagamentos de financiamento |
| | `History.GetSettlementHistory(ctx, params)` | Obter histórico de liquidações |
| | `History.GetDustHistory(ctx, params)` | Obter histórico de conversão de poeira |
| | `History.GetPositionHistory(ctx, params)` | Obter histórico de posições |
| | `History.GetStrategyHistory(ctx, params)` | Obter histórico de estratégias |
| | `History.GetRFQHistory(ctx, params)` | Obter histórico de RFQ |
| | `History.GetQuoteHistory(ctx, params)` | Obter histórico de cotações |
| | `History.GetRFQFillHistory(ctx, params)` | Obter fills de RFQ |
| | `History.GetQuoteFillHistory(ctx, params)` | Obter fills de cotações |
| **Empréstimo** | `BorrowLend.GetPositions(ctx)` | Obter posições |
| | `BorrowLend.Execute(ctx, params)` | Emprestar ou tomar emprestado |
| | `BorrowLend.GetLiquidationPrice(ctx, params)` | Obter preço de liquidação |
| **Posições** | `Positions.GetPositions(ctx, params)` | Obter posições abertas |
| **RFQ** | `RFQ.SubmitRFQ(ctx, params)` | Enviar RFQ |
| | `RFQ.SubmitQuote(ctx, params)` | Enviar cotação |
| | `RFQ.AcceptQuote(ctx, types.QuoteAcceptParams{QuoteID: quoteId})` | Aceitar cotação |
| | `RFQ.RefreshRFQ(ctx, rfqId)` | Atualizar RFQ |
| | `RFQ.CancelRFQ(ctx, types.RFQCancelParams{RfqID: rfqId})` | Cancelar RFQ |
| **Estratégia** | `Strategy.CreateStrategy(ctx, params)` | Criar estratégia |
| | `Strategy.GetStrategy(ctx, strategyId)` | Obter estratégia |
| | `Strategy.GetOpenStrategies(ctx, symbol)` | Obter estratégias abertas |
| | `Strategy.CancelStrategy(ctx, strategyId)` | Cancelar estratégia |
| | `Strategy.CancelAllStrategies(ctx, symbol)` | Cancelar todas as estratégias |

### Cliente WebSocket

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
    // Streams públicos (sem autenticação)
    wsClient := websocket.NewClient()

    // Streams privados (com autenticação)
    // wsClient := websocket.NewClient(
    //     websocket.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    // )

    if err := wsClient.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()

    handler := websocket.NewHandler(wsClient)

    // Inscrever-se no book ticker
    handler.OnBookTicker("SOL_USDC", func(ticker *types.WSBookTicker) {
        fmt.Printf("Melhor compra: %s, Melhor venda: %s\n", ticker.BidPrice, ticker.AskPrice)
    })

    // Inscrever-se em trades
    handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
        fmt.Printf("Trade: %s @ %s\n", trade.Quantity, trade.Price)
    })

    // Manter em execução
    select {}
}
```

## Enumerações Disponíveis

```go
import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Relacionadas a ordens
enums.OrderTypeLimit                  // Limit, Market
enums.SideBid                         // Bid, Ask
enums.TimeInForceGTC                  // GTC, IOC, FOK
enums.SelfTradePreventionRejectTaker  // RejectTaker, RejectMaker, RejectBoth
enums.OrderStatusNew                  // New, Filled, Cancelled, Expired, PartiallyFilled, TriggerPending, TriggerFailed

// Relacionadas ao mercado
enums.MarketTypeSpot       // SPOT, PERP, IPERP, DATED, PREDICTION, RFQ
enums.KlineInterval1m      // 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1month
enums.DepthLimit5          // "5", "10", "20", "50", "100", "500", "1000"
enums.OrderBookStateOpen   // Open, Closed, CancelOnly, LimitOnly, PostOnly

// Relacionadas ao status
enums.DepositStatusPending            // cancelled, confirmed, declined, expired, initiated, etc.
enums.WithdrawalStatusPending         // confirmed, ownershipVerificationRequired, pending, etc.

// Blockchain (30+ redes suportadas)
enums.BlockchainSolana
enums.BlockchainEthereum
enums.BlockchainBitcoin
enums.BlockchainArbitrum
enums.BlockchainBase
// ... e mais
```

## Opções de Configuração

```go
client, err := backpack.NewClient(
    backpack.WithCredentials(publicKey, secretKey),  // Credenciais da API
    backpack.WithBaseURL("https://api.backpack.exchange"),  // URL base personalizada
    backpack.WithTimeout(30 * time.Second),  // Timeout HTTP
    backpack.WithWindow(5000),  // Janela de assinatura em ms
    backpack.WithDebug(true),  // Habilitar logs de depuração
    backpack.WithHTTPClient(customClient),  // Cliente HTTP personalizado
)
```

## Exemplos

Veja o diretório [examples](./examples) para exemplos completos de uso:

- `public/main.go` - Exemplos de API pública
- `authenticated/main.go` - Exemplos de API autenticada
- `websocket/main.go` - Exemplos de streaming WebSocket
- `advanced/main.go` - Exemplos de uso avançado

## Documentação

Para documentação detalhada da API, visite [Backpack Exchange API Docs](https://docs.backpack.exchange/).

## Registro de Alterações

### v1.1.0
- Sincronização de todos os tipos e enumerações com a especificação OpenAPI
- Atualização de valores de enumeração para corresponder à API (SortDirection, OrderStatus, MarketType, etc.)
- Expansão da enumeração Blockchain de 4 para 30+ redes
- Novas enumerações adicionadas: PositionState, OrderBookState, PaymentType, RfqExecutionMode, etc.
- Correção de definições de tipos: Order, Position, Strategy, RFQ, Capital
- Atualização de assinaturas de métodos de serviço

### v1.0.0
- Lançamento inicial com suporte completo à API
- Implementação de todos os 70 endpoints de API
- Cliente WebSocket com reconexão automática
- Autenticação ED25519
- Enumerações e tipos abrangentes

## Suporte

Se este SDK foi útil para você, por favor considere:

1. Usar meu link de referência para se registrar: [Registrar no Backpack Exchange](https://backpack.exchange/refer/solomeowl)
2. Dar uma estrela a este repositório no GitHub

## Licença

MIT License

# Backpack Exchange Go SDK

![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

[English](./README.md) | [简体中文](./README_zh-Hans.md) | [繁體中文](./README_zh-Hant.md) | [日本語](./README_ja.md) | [한국어](./README_ko.md) | **Español** | [Português](./README_pt.md)

SDK completo de Go para la API de [Backpack Exchange](https://backpack.exchange/). Soporta los 70 endpoints de API incluyendo REST y WebSocket.

## Características

- **Cobertura Completa**: Implementación de los 70 endpoints de API (59 rutas únicas)
- **Autenticación ED25519**: Firma segura para endpoints autenticados
- **Seguridad de Tipos**: Estructuras de solicitud/respuesta fuertemente tipadas
- **Cliente WebSocket**: Streaming en tiempo real de datos de mercado y actualizaciones de cuenta
- **Enumeraciones**: Enumeraciones completas para llamadas de API con seguridad de tipos
- **Opciones Funcionales**: Configuración flexible del cliente

## Instalación

```bash
go get github.com/solomeowl/backpack-exchange-sdk-go
```

## Inicio Rápido

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

    // Obtener todos los mercados
    markets, _ := client.Markets.GetMarkets(ctx, nil)

    // Obtener ticker
    ticker, _ := client.Markets.GetTicker(ctx, services.GetTickerParams{Symbol: "SOL_USDC"})

    // Obtener profundidad del libro de órdenes
    depth, _ := client.Markets.GetDepth(ctx, services.GetDepthParams{Symbol: "SOL_USDC"})

    fmt.Printf("Mercados: %d, Último precio: %s, Profundidad de compra: %d\n",
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

    // Obtener saldos de cuenta
    balances, _ := client.Capital.GetBalances(ctx)

    // Ejecutar orden
    order, _ := client.Orders.ExecuteOrder(ctx, types.ExecuteOrderParams{
        Symbol:    "SOL_USDC",
        Side:      enums.SideBid,
        OrderType: enums.OrderTypeLimit,
        Price:     "100",
        Quantity:  "1",
    })

    // Obtener historial de órdenes
    history, _ := client.History.GetOrderHistory(ctx, &types.OrderHistoryParams{
        Symbol: "SOL_USDC",
    })

    fmt.Printf("Saldos: %d, Orden: %s, Historial: %d\n",
        len(balances), order.ID, len(history))
}
```

### Uso de Enumeraciones

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

## Referencia de API

### Métodos del Cliente Público

| Categoría | Método | Descripción |
|-----------|--------|-------------|
| **Sistema** | `System.GetStatus(ctx)` | Obtener estado del sistema |
| | `System.Ping(ctx)` | Ping al servidor |
| | `System.GetTime(ctx)` | Obtener hora del servidor |
| | `System.GetWallets(ctx)` | Obtener billeteras soportadas |
| **Activos** | `Assets.GetAssets(ctx)` | Obtener todos los activos |
| | `Assets.GetCollateral(ctx)` | Obtener información de colateral |
| **Mercados** | `Markets.GetMarkets(ctx, params)` | Obtener todos los mercados |
| | `Markets.GetMarket(ctx, symbol)` | Obtener mercado específico |
| | `Markets.GetTicker(ctx, params)` | Obtener ticker |
| | `Markets.GetTickers(ctx, params)` | Obtener todos los tickers |
| | `Markets.GetDepth(ctx, params)` | Obtener libro de órdenes |
| | `Markets.GetKlines(ctx, params)` | Obtener velas |
| | `Markets.GetMarkPrice(ctx, symbol)` | Obtener precio de marca |
| | `Markets.GetMarkPrices(ctx)` | Obtener todos los precios de marca |
| | `Markets.GetOpenInterest(ctx, symbol)` | Obtener interés abierto |
| | `Markets.GetFundingRates(ctx, symbol)` | Obtener tasas de financiamiento |
| **Trades** | `Trades.GetRecentTrades(ctx, symbol)` | Obtener trades recientes |
| | `Trades.GetHistoricalTrades(ctx, params)` | Obtener historial de trades |
| **Préstamo** | `BorrowLendMarkets.GetMarkets(ctx)` | Obtener mercados de préstamo |
| | `BorrowLendMarkets.GetMarketHistory(ctx, params)` | Obtener historial de préstamo |
| **Predicción** | `Prediction.GetMarkets(ctx)` | Obtener mercados de predicción |
| | `Prediction.GetTags(ctx, nil)` | Obtener etiquetas de predicción |

### Métodos del Cliente Autenticado

| Categoría | Método | Descripción |
|-----------|--------|-------------|
| **Cuenta** | `Account.GetAccount(ctx)` | Obtener configuración de cuenta |
| | `Account.UpdateAccount(ctx, settings)` | Actualizar configuración |
| | `Account.GetMaxBorrowQuantity(ctx, symbol)` | Obtener cantidad máxima de préstamo |
| | `Account.GetMaxOrderQuantity(ctx, params)` | Obtener tamaño máximo de orden |
| | `Account.GetMaxWithdrawalQuantity(ctx, params)` | Obtener retiro máximo |
| | `Account.ConvertDust(ctx, &symbol)` | Convertir polvo a USDC |
| **Capital** | `Capital.GetBalances(ctx)` | Obtener saldos |
| | `Capital.GetCollateral(ctx, nil)` | Obtener colateral |
| | `Capital.GetDeposits(ctx, params)` | Obtener historial de depósitos |
| | `Capital.GetDepositAddress(ctx, blockchain)` | Obtener dirección de depósito |
| | `Capital.GetWithdrawals(ctx, params)` | Obtener historial de retiros |
| | `Capital.RequestWithdrawal(ctx, req)` | Solicitar retiro |
| | `Capital.GetWithdrawalDelay(ctx)` | Obtener retraso de retiro |
| | `Capital.CreateWithdrawalDelay(ctx, params)` | Crear retraso de retiro |
| | `Capital.UpdateWithdrawalDelay(ctx, params)` | Actualizar retraso de retiro |
| **Órdenes** | `Orders.ExecuteOrder(ctx, params)` | Ejecutar orden |
| | `Orders.ExecuteBatchOrders(ctx, orders)` | Ejecutar órdenes en lote |
| | `Orders.GetOpenOrders(ctx, params)` | Obtener órdenes abiertas |
| | `Orders.GetOrder(ctx, params)` | Obtener orden específica |
| | `Orders.CancelOrder(ctx, params)` | Cancelar orden |
| | `Orders.CancelAllOrders(ctx, symbol)` | Cancelar todas las órdenes |
| **Historial** | `History.GetOrderHistory(ctx, params)` | Obtener historial de órdenes |
| | `History.GetFillHistory(ctx, params)` | Obtener historial de fills |
| | `History.GetBorrowHistory(ctx, params)` | Obtener historial de préstamos |
| | `History.GetInterestHistory(ctx, params)` | Obtener historial de intereses |
| | `History.GetBorrowPositionHistory(ctx, params)` | Obtener historial de posiciones |
| | `History.GetFundingPayments(ctx, params)` | Obtener pagos de financiamiento |
| | `History.GetSettlementHistory(ctx, params)` | Obtener historial de liquidaciones |
| | `History.GetDustHistory(ctx, params)` | Obtener historial de conversión de polvo |
| | `History.GetPositionHistory(ctx, params)` | Obtener historial de posiciones |
| | `History.GetStrategyHistory(ctx, params)` | Obtener historial de estrategias |
| | `History.GetRFQHistory(ctx, params)` | Obtener historial de RFQ |
| | `History.GetQuoteHistory(ctx, params)` | Obtener historial de cotizaciones |
| | `History.GetRFQFillHistory(ctx, params)` | Obtener fills de RFQ |
| | `History.GetQuoteFillHistory(ctx, params)` | Obtener fills de cotizaciones |
| **Préstamo** | `BorrowLend.GetPositions(ctx)` | Obtener posiciones |
| | `BorrowLend.Execute(ctx, params)` | Préstamo o préstamo |
| | `BorrowLend.GetLiquidationPrice(ctx, params)` | Obtener precio de liquidación |
| **Posiciones** | `Positions.GetPositions(ctx, params)` | Obtener posiciones abiertas |
| **RFQ** | `RFQ.SubmitRFQ(ctx, params)` | Enviar RFQ |
| | `RFQ.SubmitQuote(ctx, params)` | Enviar cotización |
| | `RFQ.AcceptQuote(ctx, types.QuoteAcceptParams{QuoteID: quoteId})` | Aceptar cotización |
| | `RFQ.RefreshRFQ(ctx, rfqId)` | Actualizar RFQ |
| | `RFQ.CancelRFQ(ctx, types.RFQCancelParams{RfqID: rfqId})` | Cancelar RFQ |
| **Estrategia** | `Strategy.CreateStrategy(ctx, params)` | Crear estrategia |
| | `Strategy.GetStrategy(ctx, strategyId)` | Obtener estrategia |
| | `Strategy.GetOpenStrategies(ctx, symbol)` | Obtener estrategias abiertas |
| | `Strategy.CancelStrategy(ctx, strategyId)` | Cancelar estrategia |
| | `Strategy.CancelAllStrategies(ctx, symbol)` | Cancelar todas las estrategias |

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
    // Streams públicos (sin autenticación)
    wsClient := websocket.NewClient()

    // Streams privados (con autenticación)
    // wsClient := websocket.NewClient(
    //     websocket.WithCredentials("<API_KEY>", "<SECRET_KEY>"),
    // )

    if err := wsClient.Connect(context.Background()); err != nil {
        log.Fatal(err)
    }
    defer wsClient.Close()

    handler := websocket.NewHandler(wsClient)

    // Suscribirse a book ticker
    handler.OnBookTicker("SOL_USDC", func(ticker *types.WSBookTicker) {
        fmt.Printf("Mejor compra: %s, Mejor venta: %s\n", ticker.BidPrice, ticker.AskPrice)
    })

    // Suscribirse a trades
    handler.OnTrade("SOL_USDC", func(trade *types.WSTrade) {
        fmt.Printf("Trade: %s @ %s\n", trade.Quantity, trade.Price)
    })

    // Mantener en ejecución
    select {}
}
```

## Enumeraciones Disponibles

```go
import "github.com/solomeowl/backpack-exchange-sdk-go/backpack/enums"

// Relacionadas con órdenes
enums.OrderTypeLimit                  // Limit, Market
enums.SideBid                         // Bid, Ask
enums.TimeInForceGTC                  // GTC, IOC, FOK
enums.SelfTradePreventionRejectTaker  // RejectTaker, RejectMaker, RejectBoth
enums.OrderStatusNew                  // New, Filled, Cancelled, Expired, PartiallyFilled, TriggerPending, TriggerFailed

// Relacionadas con mercado
enums.MarketTypeSpot       // SPOT, PERP, IPERP, DATED, PREDICTION, RFQ
enums.KlineInterval1m      // 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d, 3d, 1w, 1month
enums.DepthLimit5          // "5", "10", "20", "50", "100", "500", "1000"
enums.OrderBookStateOpen   // Open, Closed, CancelOnly, LimitOnly, PostOnly

// Relacionadas con estado
enums.DepositStatusPending            // cancelled, confirmed, declined, expired, initiated, etc.
enums.WithdrawalStatusPending         // confirmed, ownershipVerificationRequired, pending, etc.

// Blockchain (30+ redes soportadas)
enums.BlockchainSolana
enums.BlockchainEthereum
enums.BlockchainBitcoin
enums.BlockchainArbitrum
enums.BlockchainBase
// ... y más
```

## Opciones de Configuración

```go
client, err := backpack.NewClient(
    backpack.WithCredentials(publicKey, secretKey),  // Credenciales de API
    backpack.WithBaseURL("https://api.backpack.exchange"),  // URL base personalizada
    backpack.WithTimeout(30 * time.Second),  // Timeout HTTP
    backpack.WithWindow(5000),  // Ventana de firma en ms
    backpack.WithDebug(true),  // Habilitar logs de depuración
    backpack.WithHTTPClient(customClient),  // Cliente HTTP personalizado
)
```

## Ejemplos

Ver el directorio [examples](./examples) para ejemplos completos de uso:

- `public/main.go` - Ejemplos de API pública
- `authenticated/main.go` - Ejemplos de API autenticada
- `websocket/main.go` - Ejemplos de streaming WebSocket
- `advanced/main.go` - Ejemplos de uso avanzado

## Documentación

Para documentación detallada de la API, visita [Backpack Exchange API Docs](https://docs.backpack.exchange/).

## Registro de Cambios

### v1.1.0
- Sincronización de todos los tipos y enumeraciones con la especificación OpenAPI
- Actualización de valores de enumeración para coincidir con API (SortDirection, OrderStatus, MarketType, etc.)
- Expansión de la enumeración Blockchain de 4 a 30+ redes
- Nuevas enumeraciones añadidas: PositionState, OrderBookState, PaymentType, RfqExecutionMode, etc.
- Corrección de definiciones de tipos: Order, Position, Strategy, RFQ, Capital
- Actualización de firmas de métodos de servicio

### v1.0.0
- Lanzamiento inicial con soporte completo de API
- Implementación de los 70 endpoints de API
- Cliente WebSocket con reconexión automática
- Autenticación ED25519
- Enumeraciones y tipos completos

## Soporte

Si este SDK te ha sido útil, por favor considera:

1. Usar mi enlace de referido para registrarte: [Registrarse en Backpack Exchange](https://backpack.exchange/refer/solomeowl)
2. Dar una estrella a este repositorio en GitHub

## Licencia

MIT License

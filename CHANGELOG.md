# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-01-20

### Added
- Initial release
- Complete REST API coverage (60+ endpoints)
  - System APIs (status, ping, time, wallets)
  - Assets APIs (assets, collateral)
  - Markets APIs (markets, tickers, depth, klines, mark prices, funding rates)
  - Trades APIs (recent trades, historical trades)
  - Account APIs (account info, settings, max quantities)
  - Capital APIs (balances, deposits, withdrawals, dust conversion)
  - Orders APIs (execute, cancel, batch orders)
  - Positions APIs (futures positions)
  - Borrow/Lend APIs (positions, execute, markets)
  - RFQ APIs (submit, quote, accept, cancel)
  - Strategy APIs (create, cancel, TWAP)
  - History APIs (fills, orders, funding, settlements)
  - Prediction Markets APIs
- WebSocket client with auto-reconnect
  - Public streams (trades, tickers, depth, klines, mark prices)
  - Private streams (order updates, position updates, RFQ updates)
- ED25519 authentication
- Type-safe handlers
- Comprehensive error handling
- Examples for all major use cases

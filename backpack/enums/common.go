// Package enums provides enum definitions for Backpack Exchange SDK.
package enums

// Side represents the order side.
type Side string

const (
	SideBid Side = "Bid"
	SideAsk Side = "Ask"
)

// Blockchain represents supported blockchain networks.
type Blockchain string

const (
	BlockchainSolana   Blockchain = "Solana"
	BlockchainEthereum Blockchain = "Ethereum"
	BlockchainPolygon  Blockchain = "Polygon"
	BlockchainBitcoin  Blockchain = "Bitcoin"
)

// SortDirection represents sort order.
type SortDirection string

const (
	SortAsc  SortDirection = "ASC"
	SortDesc SortDirection = "DESC"
)

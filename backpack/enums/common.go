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
	Blockchain0G          Blockchain = "0G"
	BlockchainAptos       Blockchain = "Aptos"
	BlockchainArbitrum    Blockchain = "Arbitrum"
	BlockchainAvalanche   Blockchain = "Avalanche"
	BlockchainBase        Blockchain = "Base"
	BlockchainBerachain   Blockchain = "Berachain"
	BlockchainBitcoin     Blockchain = "Bitcoin"
	BlockchainBitcoinCash Blockchain = "BitcoinCash"
	BlockchainBsc         Blockchain = "Bsc"
	BlockchainCardano     Blockchain = "Cardano"
	BlockchainDogecoin    Blockchain = "Dogecoin"
	BlockchainEclipse     Blockchain = "Eclipse"
	BlockchainEqualsMoney Blockchain = "EqualsMoney"
	BlockchainEthereum    Blockchain = "Ethereum"
	BlockchainFogo        Blockchain = "Fogo"
	BlockchainHyperEVM    Blockchain = "HyperEVM"
	BlockchainHyperliquid Blockchain = "Hyperliquid"
	BlockchainLinea       Blockchain = "Linea"
	BlockchainLitecoin    Blockchain = "Litecoin"
	BlockchainMonad       Blockchain = "Monad"
	BlockchainNear        Blockchain = "Near"
	BlockchainOptimism    Blockchain = "Optimism"
	BlockchainPlasma      Blockchain = "Plasma"
	BlockchainPolygon     Blockchain = "Polygon"
	BlockchainSei         Blockchain = "Sei"
	BlockchainSolana      Blockchain = "Solana"
	BlockchainStable      Blockchain = "Stable"
	BlockchainStory       Blockchain = "Story"
	BlockchainSui         Blockchain = "Sui"
	BlockchainTron        Blockchain = "Tron"
	BlockchainXRP         Blockchain = "XRP"
	BlockchainZcash       Blockchain = "Zcash"
)

// SortDirection represents sort order.
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "Asc"
	SortDirectionDesc SortDirection = "Desc"
)

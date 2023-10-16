package chain

import (
	"errors"
	"math/big"
	"slices"
)

type Chain int

const (
	Arbitrum       Chain = 42161    // Arbitrum One
	ArbitrumGoerli Chain = 421613   // Arbitrum Goerli Testnet
	ArbitrumNova   Chain = 42170    // Arbitrum Nova
	Avalanche      Chain = 43114    // Avalanche C-Chain
	AvalancheFuji  Chain = 43113    // Avalanche Fuji Testnet
	Baobab         Chain = 1001     // Klaytn Testnet
	Base           Chain = 8453     // Base
	BaseGoerli     Chain = 84531    // Base Goerli Testnet
	BSC            Chain = 56       // BNB Smart Chain Mainnet
	BSCTestNet     Chain = 97       // BNB Smart Chain Testnet
	Ethereum       Chain = 1        // Ethereum Mainnet
	Goerli         Chain = 5        // Goerli
	Klaytn         Chain = 8217     // Klaytn Mainnet Cypress
	Matic          Chain = 137      // Polygon Mainnet
	Mumbai         Chain = 80001    // Mumbai
	Optimism       Chain = 10       // OP Mainnet
	OptimismGoerli Chain = 420      // Optimism Goerli Testnet
	Sepolia        Chain = 11155111 // Sepolia Testnet
	Solana         Chain = -1       // Solana is not an Ethereum EVM compatible network.
	Soldev         Chain = -2       // Solana Devnet
	Zora           Chain = 7777777  // Zora
	ZoraTestNet    Chain = 999      // Zora Goerli Testnet
)

var (
	MainnetChains = []Chain{
		Arbitrum, ArbitrumNova, Avalanche, Base, BSC,
		Ethereum, Klaytn, Matic, Optimism, Solana, Zora,
	}

	TestnetChains = []Chain{
		ArbitrumGoerli, AvalancheFuji, Baobab, BaseGoerli, BSCTestNet,
		Goerli, OptimismGoerli, Mumbai, Sepolia, Soldev, ZoraTestNet,
	}
)

func (c Chain) ChainId() int {
	return int(c)
}

func (c Chain) ChainIdBigInt() *big.Int {
	return big.NewInt(int64(c.ChainId()))
}

func (c Chain) Currency() string {
	switch c {
	case Arbitrum:
		return "ETH"
	case ArbitrumGoerli:
		return "AGOR"
	case ArbitrumNova:
		return "ETH"
	case Avalanche:
		return "AVAX"
	case AvalancheFuji:
		return "AVAX"
	case Baobab:
		return "KLAY"
	case Base:
		return "ETH"
	case BaseGoerli:
		return "ETH"
	case BSC:
		return "BNB"
	case BSCTestNet:
		return "tBNB"
	case Ethereum:
		return "ETH"
	case Goerli:
		return "ETH"
	case Klaytn:
		return "KLAY"
	case Matic:
		return "MATIC"
	case Mumbai:
		return "MATIC"
	case Optimism:
		return "ETH"
	case OptimismGoerli:
		return "ETH"
	case Sepolia:
		return "ETH"
	case Solana:
		return "SOL"
	case Soldev:
		return "SOL"
	case Zora:
		return "ETH"
	case ZoraTestNet:
		return "ETH"
	default:
		return ""
	}
}

func (c Chain) Name() string {
	switch c {
	case Arbitrum:
		return "Arbitrum One"
	case ArbitrumGoerli:
		return "Arbitrum Goerli"
	case ArbitrumNova:
		return "Arbitrum Nova"
	case Avalanche:
		return "Avalanche C-Chain"
	case AvalancheFuji:
		return "Avalanche Fuji Testnet"
	case Baobab:
		return "Klaytn Testnet Baobab"
	case Base:
		return "Base"
	case BaseGoerli:
		return "Base Goerli Testnet"
	case BSC:
		return "BNB Smart Chain Mainnet"
	case BSCTestNet:
		return "BNB Smart Chain Testnet"
	case Ethereum:
		return "Ethereum Mainnet"
	case Goerli:
		return "Goerli"
	case Klaytn:
		return "Klaytn Mainnet Cypress"
	case Matic:
		return "Polygon Mainnet"
	case Mumbai:
		return "<Mumbai"
	case Optimism:
		return "OP Mainnet"
	case OptimismGoerli:
		return "Optimism Goerli Testnet"
	case Sepolia:
		return "Sepolia"
	case Solana:
		return "Solana"
	case Soldev:
		return "Solana Devnet"
	case Zora:
		return "Zora"
	case ZoraTestNet:
		return "Zora Testnet"
	default:
		return ""
	}
}

var valuesMapping = map[string]Chain{
	"arbitrum":        Arbitrum,
	"arbitrum_goerli": ArbitrumGoerli,
	"arbitrum_nova":   ArbitrumNova,
	"avalanche":       Avalanche,
	"avalanche_fuji":  AvalancheFuji,
	"baobab":          Baobab,
	"base":            Base,
	"base_goerli":     BaseGoerli,
	"bsc":             BSC,
	"bsctestnet":      BSCTestNet,
	"ethereum":        Ethereum,
	"goerli":          Goerli,
	"klaytn":          Klaytn,
	"matic":           Matic,
	"mumbai":          Mumbai,
	"optimism":        Optimism,
	"optimism_goerli": OptimismGoerli,
	"sepolia":         Sepolia,
	"solana":          Solana,
	"soldev":          Soldev,
	"zora":            Zora,
	"zora_testnet":    ZoraTestNet,
}

func (c Chain) Value() string {
	switch c {
	case Arbitrum:
		return "arbitrum"
	case ArbitrumGoerli:
		return "arbitrum_goerli"
	case ArbitrumNova:
		return "arbitrum_nova"
	case Avalanche:
		return "avalanche"
	case AvalancheFuji:
		return "avalanche_fuji"
	case Baobab:
		return "baobab"
	case Base:
		return "base"
	case BaseGoerli:
		return "base_goerli"
	case BSC:
		return "bsc"
	case BSCTestNet:
		return "bsctestnet"
	case Ethereum:
		return "ethereum"
	case Goerli:
		return "goerli"
	case Klaytn:
		return "klaytn"
	case Matic:
		return "matic"
	case Mumbai:
		return "mumbai"
	case Optimism:
		return "optimism"
	case OptimismGoerli:
		return "optimism_goerli"
	case Sepolia:
		return "sepolia"
	case Solana:
		return "solana"
	case Soldev:
		return "soldev"
	case Zora:
		return "zora"
	case ZoraTestNet:
		return "zora_testnet"
	default:
		return ""
	}
}

func (c Chain) IsTestNet() bool {
	return slices.Contains(TestnetChains, c)
}

func NewFromString(ch string) (Chain, error) {
	c, ok := valuesMapping[ch]
	if !ok {
		return -1, errors.New("unknown chain")
	}
	return c, nil
}

func RequireFromString(ch string) Chain {
	c, ok := valuesMapping[ch]
	if !ok {
		panic("unknown chain")
	}
	return c
}

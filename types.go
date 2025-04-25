import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block represents a simplified Ethereum block.
type Block struct {
	Header       BlockHeader
	Transactions []Transaction
}

type BlockHeader struct {
	ParentHash string
	Number     uint64
	Timestamp  uint64
	StateRoot  string // Merkle root of state (simplified)
}
// Transaction represents a basic Ethereum transaction.
type Transaction struct {
	From     string
	To       string
	Value    uint64
	Data     []byte // For smart contracts
	GasLimit uint64
	Nonce    uint64
}

// Account represents a user/contract account.
type Account struct {
	Balance uint64
	Nonce   uint64
	Code    []byte // For smart contracts
}
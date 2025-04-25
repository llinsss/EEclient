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
// WorldState tracks account balances (simplified).
type WorldState map[string]Account

// NewBlock creates a genesis block.
func NewGenesisBlock() Block {
	return Block{
		Header: BlockHeader{
			ParentHash: "0x0",
			Number:     0,
			Timestamp:  uint64(time.Now().Unix()),
			StateRoot:  "0x0",
		},
		Transactions: []Transaction{},
	}
}
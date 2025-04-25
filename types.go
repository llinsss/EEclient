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

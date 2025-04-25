package main

import (
	"encoding/json"
	"fmt"
)

type Blockchain struct {
	Blocks     []Block
	WorldState WorldState
}

func NewBlockchain() *Blockchain {
	genesis := NewGenesisBlock()
	return &Blockchain{
		Blocks:     []Block{genesis},
		WorldState: make(WorldState),
	}
}
// AddBlock processes transactions and adds a new block.
func (bc *Blockchain) AddBlock(txs []Transaction) error {
	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	// Process transactions (simplified)
	newState := bc.processTransactions(txs)

	// Create new block
	newBlock := Block{
		Header: BlockHeader{
			ParentHash: hashBlock(lastBlock),
			Number:     lastBlock.Header.Number + 1,
			Timestamp:  uint64(time.Now().Unix()),
			StateRoot:  "0x1", // In reality, compute Merkle root
		},
		Transactions: txs,
	}
	bc.Blocks = append(bc.Blocks, newBlock)
	bc.WorldState = newState
	return nil
}

// processTransactions updates the world state.
func (bc *Blockchain) processTransactions(txs []Transaction) WorldState {
	newState := copyWorldState(bc.WorldState)

	for _, tx := range txs {
		// Simplified: Just transfer value
		sender := newState[tx.From]
		receiver := newState[tx.To]

		sender.Balance -= tx.Value
		receiver.Balance += tx.Value

		newState[tx.From] = sender
		newState[tx.To] = receiver
	}
	return newState
}

// Helper functions
func hashBlock(block Block) string {
	bytes, _ := json.Marshal(block)
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}

func copyWorldState(ws WorldState) WorldState {
	newWs := make(WorldState)
	for k, v := range ws {
		newWs[k] = v
	}
	return newWs
}
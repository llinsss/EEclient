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

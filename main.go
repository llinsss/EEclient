package main

func main() {
	bc := NewBlockchain()

	// Create test accounts
	bc.WorldState["alice"] = Account{Balance: 1000}
	bc.WorldState["bob"] = Account{Balance: 500}

	// Create a transaction
	tx := Transaction{
		From:  "alice",
		To:    "bob",
		Value: 100,
		Nonce: 1,
	}

	// Add a block with the transaction
	bc.AddBlock([]Transaction{tx})

	// Print final state
	fmt.Printf("Alice balance: %d\n", bc.WorldState["alice"].Balance)
	fmt.Printf("Bob balance: %d\n", bc.WorldState["bob"].Balance)
	fmt.Printf("Blockchain length: %d\n", len(bc.Blocks))
}